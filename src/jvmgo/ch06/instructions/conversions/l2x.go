package conversions

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type L2F struct {
	base.NoOperandsInstruction
}

type L2D struct {
	base.NoOperandsInstruction
}

type L2I struct {
	base.NoOperandsInstruction
}

func (e *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	d := float64(l)
	stack.PushDouble(d)
}

func (e *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	f := float32(l)
	stack.PushFloat(f)
}

func (e *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l := stack.PopLong()
	i := int32(l)
	stack.PushInt(i)
}
