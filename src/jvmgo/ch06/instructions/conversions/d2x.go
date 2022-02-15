package conversions

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	double 的一些转换
*/

type D2F struct {
	base.NoOperandsInstruction
}

type D2I struct {
	base.NoOperandsInstruction
}

type D2L struct {
	base.NoOperandsInstruction
}

func (e *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

func (e *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

func (e *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
