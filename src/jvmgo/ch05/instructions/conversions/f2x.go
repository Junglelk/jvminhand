package conversions

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	double 的一些转换
*/

type F2D struct {
	base.NoOperandsInstruction
}

type F2I struct {
	base.NoOperandsInstruction
}

type F2L struct {
	base.NoOperandsInstruction
}

func (e *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f := stack.PopFloat()
	i := int32(f)
	stack.PushInt(i)
}

func (e *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	f := float64(d)
	stack.PushDouble(f)
}

func (e *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopFloat()
	l := int64(d)
	stack.PushLong(l)
}
