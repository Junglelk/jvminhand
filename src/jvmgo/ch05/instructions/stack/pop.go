package stack

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	栈指令直接对操作数栈进行操作，共 9 条。
	pop和pop2指令将栈顶变量弹出，dup系列指令复制栈顶变量，swap指令交换栈顶的两个变量。
*/

// POP 指令将栈顶变量弹出，只能弹出占用了一个位置的变量
type POP struct {
	base.NoOperandsInstruction
}

// POP2 指令将long、double等占用两个位置的变量弹出
type POP2 struct {
	base.NoOperandsInstruction
}

func (e *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

func (e *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
