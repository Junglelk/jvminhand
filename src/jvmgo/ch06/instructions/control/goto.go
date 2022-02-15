package control

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	控制指令有 11 条，jsr和ret指令在Java6之后就不再用来实现finally子句，
	return系列指令有6条，用于从方法调用中返回。本章不实现。
	剩下三条为：goto、tableswitch、lookupswitch
*/

type GOTO struct {
	base.BranchInstruction
}

func (e *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, e.Offset)
}
