package heap

import (
	"JVM-GO/ch06/classfile"
	"JVM-GO/ch06/classpath"
	"fmt"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/
type ClassLoader struct {
	cp       *classpath.Classpath  // ClassLoader依赖Classpath来搜索和读取class文件，cp字段保存Classpath指针
	classMap map[string]*Class // 记录已经加载的类数据，key是类的完全限定名
}

/**
创建ClassLoader实例
 */
func NewClassLoader(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

/**
把类数据加载到方法区
 */
func (self *ClassLoader) LoadClass(name string) *Class {
	// 先查找classMap，看类是否已经被加载
	if class, ok := self.classMap[name]; ok {
		// 如果是，直接返回类数据
		return class
	}
	// 否则调用loadNonArrayClass（）方法加载类
	return self.loadNonArrayClass(name)
}

func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 首先找到class文件并把数据读取到内存
	data, entry := self.readClass(name)
	// 然后解析class文件，生成虚拟机可以使用的类数据，并放入方法区
	class := self.defineClass(data)
	// 最后进行链接
	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class
}

/**
调用了Classpath的ReadClass（）方法，并进行了错误处理。
需要解释一下它的返回值。为了打印类加载信息，把最终加载class文件的类路径项也返回给了调用者。
 */
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.cp.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// jvms 5.3.5
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 首先调用parseClass（）函数把class文件数据转换成Class结构体。
	class := parseClass(data)
	class.loader = self
	// Class结构体的superClass和interfaces字段存放超类名和直接接口表，这些类名其实都是符号引用
	// 调用resolveSuperClass（）和resolveInterfaces（）函数解析这些类符号引用
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// jvms 5.4.3.1
func resolveSuperClass(class *Class) {
	// 除java.lang.Object以外，所有的类都有且仅有一个超类。
	// 因此，除非是Object类，否则需要递归调用LoadClass（）方法加载它的超类。
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	// 验证阶段 为了确保安全性，Java虚拟机规范要求在执行类的任何代码之前，对类进行严格的验证
	verify(class)
	// 准备阶段主要是给类变量分配空间并给予初始值
	prepare(class)
}

func verify(class *Class) {
	// todo
}

// jvms 5.4.2
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class) // 计算实例字段的个数 同时给它们编号
	calcStaticFieldSlotIds(class) // 计算静态字段的个数 同时给它们编号
	allocAndInitStaticVars(class) // 给类变量分配空间，然后给它们赋予初始值
}

func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		// 因为Go语言会保证新创建的Slot结构体有默认值（num字段是0，ref字段是nil），而浮点数0编码之后和整数0相同，所以不用做任何操作就可以保证静态变量有默认初始值（数字类型是0，引用类型是null）。
		// 如果静态变量属于基本类型或String类型，有final修饰符，且它的值在编译期已知，则该值存储在class文件常量池中。
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field) // 从常量池中加载常量值，然后给静态变量赋值
		}
	}
}

func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			panic("todo")
		}
	}
}
