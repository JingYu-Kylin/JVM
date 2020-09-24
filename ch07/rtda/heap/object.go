package heap

/**
引用类型
 */
type Object struct {
	class  *Class // 对象的Class指针
	fields Slots // 实例变量
}

func newObject(class *Class) *Object {
	return &Object{
		class:  class,
		fields: newSlots(class.instanceSlotCount),
	}
}

// getters
func (self *Object) Class() *Class {
	return self.class
}
func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
	// 这个方法稍微有些复杂，为了避免class.go文件变得过长，把它写在另一个文件中
}
