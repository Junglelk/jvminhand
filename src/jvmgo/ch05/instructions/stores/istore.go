package stores

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	存储指令将变量从操作数栈中弹出，然后存入局部变量表
*/

type ISTORE struct {
	base.Index8Instruction
}

type ISTORE_0 struct {
	base.NoOperandsInstruction
}
type ISTORE_1 struct {
	base.NoOperandsInstruction
}
type ISTORE_2 struct {
	base.NoOperandsInstruction
}
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func _istore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

func (e *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, e.Index)
}

func (e *ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (e ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (e ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (e ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}
