package stores

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	存储指令将变量从操作数栈中弹出，然后存入局部变量表
*/

type LSTORE struct {
	base.Index8Instruction
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}
type LSTORE_1 struct {
	base.NoOperandsInstruction
}
type LSTORE_2 struct {
	base.NoOperandsInstruction
}
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

func (e *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, e.Index)
}

func (e *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (e LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (e LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (e LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}
