package math

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type DDIV struct {
	base.NoOperandsInstruction
}
type FDIV struct {
	base.NoOperandsInstruction
}

type IDIV struct {
	base.NoOperandsInstruction
}
type LDIV struct {
	base.NoOperandsInstruction
}

func (e *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushInt(result)
}

func (e *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushLong(result)
}

func (e *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushDouble(result)
}

func (e *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushFloat(result)
}
