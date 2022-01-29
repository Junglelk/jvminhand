package math

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

type DSUB struct {
	base.NoOperandsInstruction
}
type FSUB struct {
	base.NoOperandsInstruction
}

type ISUB struct {
	base.NoOperandsInstruction
}
type LSUB struct {
	base.NoOperandsInstruction
}

func (e *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// java 虚拟机规范规定第一个值为v2
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	stack.PushDouble(result)
}

func (e *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// java 虚拟机规范规定第一个值为v2
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	stack.PushFloat(result)
}

func (e *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// java 虚拟机规范规定第一个值为v2
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	stack.PushInt(result)
}

func (e *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// java 虚拟机规范规定第一个值为v2
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	stack.PushLong(result)
}
