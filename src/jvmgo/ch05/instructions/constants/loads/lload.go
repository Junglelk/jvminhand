package loads

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

type LLOAD struct {
	base.Index8Instruction
}

type LLOAD_0 struct {
	base.NoOperandsInstruction
}
type LLOAD_1 struct {
	base.NoOperandsInstruction
}
type LLOAD_2 struct {
	base.NoOperandsInstruction
}
type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

func (e *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, uint(e.Index))
}

func (e *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (e *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (e LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (e LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
