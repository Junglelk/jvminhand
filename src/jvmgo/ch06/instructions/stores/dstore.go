package stores

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	存储指令将变量从操作数栈中弹出，然后存入局部变量表
*/

type DSTORE struct {
	base.Index8Instruction
}

type DSTORE_0 struct {
	base.NoOperandsInstruction
}
type DSTORE_1 struct {
	base.NoOperandsInstruction
}
type DSTORE_2 struct {
	base.NoOperandsInstruction
}
type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func _dstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

func (e *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, e.Index)
}

func (e *DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

func (e DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

func (e DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

func (e DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
