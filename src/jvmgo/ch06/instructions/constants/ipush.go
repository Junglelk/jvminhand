package constants

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	bipush 指令从操作数中获取一个byte型整数，扩展成int型，然后推入栈顶；
	sipush 指令从操作数中获取一个short型整数，扩展成int型，然后推入栈顶。
*/

type BIPUSH struct {
	val int8
}

type SIPUSH struct {
	val int16
}

func (e *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	e.val = reader.ReadInt8()
}

func (e *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(e.val)
	frame.OperandStack().PushInt(i)
}

func (e *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	e.val = reader.ReadInt16()
}

func (e *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(e.val)
	frame.OperandStack().PushInt(i)
}
