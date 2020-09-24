package heap

import (
	"JVM-GO/ch07/classfile"
	"strings"
)

// name, superClassName and interfaceNames are all binary names(jvms8-4.2.1)
type Class struct {
	accessFlags       uint16 // 类的访问标志，总共16比特
	// 这些类名都是完全限定名，具有java/lang/Object的形式
	name              string // 类名
	superClassName    string // 超类名
	interfaceNames    []string // 接口名

	constantPool      *ConstantPool //存放运行时常量池指针
	fields            []*Field // 字段表
	methods           []*Method // 方法表
	loader            *ClassLoader
	superClass        *Class
	interfaces        []*Class
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
	initStarted       bool // 类的初始化其实分为几个阶段，但由于我们的类加载器还不够完善，所以先使用一个简单的布尔状态就足够了。initStarted字段 表示类的<clinit>方法是否已经开始执行。
}

/**
用来把ClassFile结构体转换成Class结构体
 */
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

/**
判断某个访问标志是否被设置
 */
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (self *Class) Name() string {
	return self.name
}
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) Fields() []*Field {
	return self.fields
}
func (self *Class) Methods() []*Method {
	return self.methods
}
func (self *Class) SuperClass() *Class {
	return self.superClass
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}
// InitStarted（）是Getter方法，返回initStarted字段值
func (self *Class) InitStarted() bool {
	return self.initStarted
}
// StartInit（）方法把initStarted字段设置成true
func (self *Class) StartInit() {
	self.initStarted = true
}

// Java虚拟机规范5.4.4节给出了类的访问控制规则，把这个规则翻译成Class结构体的isAccessibleTo（）方法
func (self *Class) isAccessibleTo(other *Class) bool {
	// 如果类D想访问类C，需要满足两个条件之一：C是public，或者C和D在同一个运行时包内
	return self.IsPublic() ||
		self.GetPackageName() == other.GetPackageName()
}

func (self *Class) GetPackageName() string {
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	return ""
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}
func (self *Class) GetClinitMethod() *Method {
	return self.getStaticMethod("<clinit>", "()V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

func (self *Class) NewObject() *Object {
	return newObject(self)
}
