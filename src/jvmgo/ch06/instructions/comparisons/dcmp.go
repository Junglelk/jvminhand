package comparisons

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type DCMPG struct {
	base.NoOperandsInstruction
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (e *DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func (e *DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}
