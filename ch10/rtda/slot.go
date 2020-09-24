package rtda

import "JVM-GO/ch10/rtda/heap"

type Slot struct {
	num int32
	ref *heap.Object
}
