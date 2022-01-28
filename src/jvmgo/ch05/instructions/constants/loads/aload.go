package loads

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

type ALOAD struct {
	base.Index8Instruction
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}
type ALOAD_1 struct {
	base.NoOperandsInstruction
}
type ALOAD_2 struct {
	base.NoOperandsInstruction
}
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func _aload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

func (e *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, uint(e.Index))
}

func (e *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (e *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (e ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (e ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}
