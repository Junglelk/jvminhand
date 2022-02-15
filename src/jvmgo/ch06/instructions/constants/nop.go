package constants

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	常量指令，把常量推入操作数栈顶。
	常量可以来自于三个地方：
			1. 隐含在操作码里；
			2. 操作数；
			3. 运行时常量池。
	常量指令有 21 条。
*/

type NOP struct {
	base.NoOperandsInstruction
}

func (e *NOP) Execute(frame *rtda.Frame) {
	// 什么都不做（神奇海螺的声音）
}
