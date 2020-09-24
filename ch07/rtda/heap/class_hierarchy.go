package heap


// jvms8 6.5.instanceof
// jvms8 6.5.checkcast
func (self *Class) isAssignableFrom(other *Class) bool {
	// 在三种情况下，S类型的引用值可以赋值给T类型：S和T是同一类型；T是类且S是T的子类；或者T是接口且S实现了T接口。
	// 这是简化版的判断逻辑，因为还没有实现数组
	s, t := other, self

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

// self extends c
// 判断S是否是T的子类，实际上也就是判断T是否是S的（直接或间接）超类。
func (self *Class) IsSubClassOf(other *Class) bool {
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// self implements iface
// 判断S是否实现了T接口，就看S或S的（直接或间接）超类是否实现了某个接口T'，T'要么是T，要么是T的子接口。
func (self *Class) IsImplements(iface *Class) bool {
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
// isSubInterfaceOf（）方法和isSubClassOf（）方法类似，但是用到了递归，
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

// c extends self
func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}
