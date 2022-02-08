package comparisons

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	比较指令可以分为两类：
		1. 将比较结果推入操作数栈顶；
		2. 根据比较结果跳转
	比较指令是实现if-else、for、while等语句的基石
*/

type LCMP struct {
	base.NoOperandsInstruction
}

func (e *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
