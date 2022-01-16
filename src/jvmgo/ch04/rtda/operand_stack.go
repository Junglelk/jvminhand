package rtda

import "math"

/*
	操作数栈大小是编译器已经确定的，所以可以用 []Slot 来实现。
	size 字段记录栈顶位置。
*/
type OperandStack struct {
	size  uint
	slots []Slot
}

func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{slots: make([]Slot, maxStack)}
	}
	return nil
}

/*
	和局部变量表一样，需要定义一些方法，从操作数栈中弹出或者压入各种类型的变量
*/

func (e *OperandStack) PushInt(val int32) {
	e.slots[e.size].num = val
	e.size++
}

func (e *OperandStack) PopInt() int32 {
	e.size--
	return e.slots[e.size].num
}

func (e *OperandStack) PushFloat(val float32) {
	e.slots[e.size].num = int32(math.Float32bits(val))
	e.size++
}

func (e *OperandStack) PopFloat() float32 {
	e.size--
	return math.Float32frombits(uint32(e.slots[e.size].num))
}

func (e *OperandStack) PushLong(val int64) {
	e.slots[e.size].num = int32(val)
	e.slots[e.size+1].num = int32(val >> 32)
	e.size += 2
}

func (e *OperandStack) PopLong() int64 {
	e.size -= 2
	low := uint32(e.slots[e.size].num)
	high := uint32(e.slots[e.size+1].num)
	return int64(high)<<32 | int64(low)
}

func (e *OperandStack) PushDouble(val float64) {
	e.PushLong(int64(math.Float64bits(val)))
}

func (e *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(e.PopLong()))
}

func (e *OperandStack) PushRef(val *Object) {
	e.slots[e.size].ref = val
	e.size++
}

func (e *OperandStack) PopRef() *Object {
	e.size--
	ref := e.slots[e.size].ref
	e.slots[e.size].ref = nil
	return ref
}
