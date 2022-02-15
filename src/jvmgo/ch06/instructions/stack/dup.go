package stack

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

// DUP 复制栈顶的单个变量
type DUP struct {
	base.NoOperandsInstruction
}

// DUP_X1 复制栈顶元素，到一个元素前
type DUP_X1 struct {
	base.NoOperandsInstruction
}

// DUP_X2 复制栈顶元素，到两个元素前
type DUP_X2 struct {
	base.NoOperandsInstruction
}
type DUP2 struct {
	base.NoOperandsInstruction
}
type DUP2_X1 struct {
	base.NoOperandsInstruction
}
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

/*
	bottom------>top
	....[d][c][b][a]
				   \_
				     |
					 v
	....[d][c][b][a][a]
*/
// Execute 实际操作
func (e *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

/*
	bottom------>top
	....[c][b][a]
			 _/
			|
			v
	....[c][a][b][a]
*/
// Execute 操作
func (e *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

/*
	bottom------>top
	....[c][b][a]
		  ____/
		 |
		 v
	....[a][c][b][a]
*/
// Execute 操作
func (e *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (e *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (e *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}

func (e *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
}
