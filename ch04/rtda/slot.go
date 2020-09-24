package rtda

/**
 * 定义一个结构体，让它可以同时容纳一个int值和一个引用值
 */
type Slot struct {
	num int32 // 整数
	ref *Object // 引用
}

