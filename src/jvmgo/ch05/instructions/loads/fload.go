package loads

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

type FLOAD struct {
	base.Index8Instruction
}

type FLOAD_0 struct {
	base.NoOperandsInstruction
}
type FLOAD_1 struct {
	base.NoOperandsInstruction
}
type FLOAD_2 struct {
	base.NoOperandsInstruction
}
type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

func (e *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, uint(e.Index))
}

func (e *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (e *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (e FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (e FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
