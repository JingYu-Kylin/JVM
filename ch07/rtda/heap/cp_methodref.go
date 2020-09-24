package heap

import "JVM-GO/ch07/classfile"

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

func (self *MethodRef) resolveMethodRef() {
	// 如果类D想通过方法符号引用访问类C的某个方法，先要解析符号引用得到类C。
	d := self.cp.class
	c := self.ResolvedClass()
	// 如果C是接口，则抛出IncompatibleClassChangeError异常，否则根据方法名和描述符查找方法。
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(c, self.name, self.descriptor)
	// 如果找不到对应的方法，则抛出NoSuchMethodError异常，否则检查类D是否有权限访问该方法。
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 如果没有，则抛出IllegalAccessError异常。
	if !method.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

func lookupMethod(class *Class, name, descriptor string) *Method {
	// 先从C的继承层次中找
	method := LookupMethodInClass(class, name, descriptor)
	// 如果找不到，就去C的接口中找。
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}