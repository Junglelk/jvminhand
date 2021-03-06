package stores

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	存储指令将变量从操作数栈中弹出，然后存入局部变量表
*/

type FSTORE struct {
	base.Index8Instruction
}

type FSTORE_0 struct {
	base.NoOperandsInstruction
}
type FSTORE_1 struct {
	base.NoOperandsInstruction
}
type FSTORE_2 struct {
	base.NoOperandsInstruction
}
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func _fstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

func (e *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, e.Index)
}

func (e *FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

func (e FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

func (e FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

func (e FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
