package math

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type IXOR struct {
	base.NoOperandsInstruction
}

type LXOR struct {
	base.NoOperandsInstruction
}

func (e *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

func (e *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
