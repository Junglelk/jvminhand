package constants

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	加载指令从局部变量表中获取变量，然后推入操作数栈。
	加载指令共 33 条，按照操作类型的变量共有 6 类：
											1. aload 系列指令操作引用类型变量；
											2. dload 系列指令操作double类型的变量；
											3. fload 系列指令操作float类型的变量；
											4. iload 系列指令操作int类型的变量；
											5. lload 系列的指令操作long类型的变量；
											6. xaload 系列的指令操作数组。
	本节定义 25 条，数组和xaload后续完成
*/

type ILOAD struct {
	base.Index8Instruction
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}
type ILOAD_1 struct {
	base.NoOperandsInstruction
}
type ILOAD_2 struct {
	base.NoOperandsInstruction
}
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

func (e *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(e.Index))
}

func (e *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (e *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (e ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (e ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
