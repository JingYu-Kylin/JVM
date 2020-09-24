package heap

// symbolic reference
type SymRef struct {
	cp        *ConstantPool // 符号引用所在的运行时常量池指针 这样就可以通过符号引用访问到运行时常量池，进一步又可以访问到类数据
	className string // 类的完全限定名
	class     *Class // 解析后的类结构体指针  这样类符号引用只需要解析一次就可以了，后续可以直接使用缓存值

	// 对于类符号引用，只要有类名，就可以解析符号引用。
	// 对于字段，首先要解析类符号引用得到类数据，然后用字段名和描述符查找字段数据。
	// 方法符号引用的解析过程和字段符号引用类似。
}

func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		// 否则调用resolveClassRef（）方法进行解析
		self.resolveClassRef()
	}
	// 如果类符号引用已经解析，ResolvedClass（）方法直接返回类指针
	return self.class
}

// jvms8 5.4.3.1
func (self *SymRef) resolveClassRef() {
	// 如果类D通过符号引用N引用类C的话 要解析N，

	//先用D的类加载器加载C
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	//然后检查D是否有权限访问C
	// 如果没有，则抛出IllegalAccessError异常
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}

	self.class = c
}
