package heap

import "JVM-GO/ch07/classfile"

/**
存放字段和方法共有的信息
 */
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class // 存放Class结构体指针 这样可以通过字段或方法访问到它所属的类
}

/**
从class文件中复制数据
 */
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

// getters
func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// jvms 5.4.4
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	// 如果字段是public，则任何类都可以访问
	if self.IsPublic() {
		return true
	}
	// 如果字段是protected，则只有子类和同一个包下的类可以访问
	c := self.class
	if self.IsProtected() {
		return d == c || d.IsSubClassOf(c) ||
			c.GetPackageName() == d.GetPackageName()
	}
	// 如果字段有默认访问权限（非public，非protected，也非privated），则只有同一个包下的类可以访问
	if !self.IsPrivate() {
		return c.GetPackageName() == d.GetPackageName()
	}
	// 否则，字段是private，只有声明这个字段的类才能访问。
	return d == c
}