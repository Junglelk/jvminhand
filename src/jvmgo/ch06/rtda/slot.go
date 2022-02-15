package rtda

import "jvmgo/jvmgo/ch06/rtda/heap"

/*
	局部变量表是按索引访问的，所以可以想象成一个数组。根据Java虚拟机规范，
	这个数组的每个元素至少可以容纳一个 int 或引用值，两个连续的元素可以容纳一个 long 或 double 值。
	Golang 的几种实现方式都不能较好地实现该“数组”，所以自行定义一个结构体，使其同时容纳一个 int 值
	和引用值。
*/
// Slot num 字段存放整数，ref 字段存放引用。
type Slot struct {
	num int32
	ref *heap.Object
}
