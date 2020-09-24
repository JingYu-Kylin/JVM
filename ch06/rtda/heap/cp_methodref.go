package heap

import "JVM-GO/ch06/classfile"

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
	// 如果还没有解析过符号引用，调用resolveMethodRef（）方法进行解析，否则直接返回方法指针
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// jvms8 5.4.3.3
func (self *MethodRef) resolveMethodRef() {
	//class := self.Class()
	// todo
}
