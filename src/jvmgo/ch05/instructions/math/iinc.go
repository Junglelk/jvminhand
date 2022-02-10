package math

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	iinc指令给局部变量表中的int变量添加常量值，局部量表索引和常量值都由指令的操作数提供。
*/

type IINC struct {
	Index uint
	Const int32
}

// FetchOperands 方法从字节码中获取操作数
func (e *IINC) FetchOperands(reader *base.BytecodeReader) {
	e.Index = uint(reader.ReadUint8())
	e.Const = int32(reader.ReadInt8())
}

// Execute 从局部变量表中获取变量，给他加上常量值，再将结果写回局部变量表
func (e *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(e.Index)
	val += e.Const
	localVars.SetInt(e.Index, val)
}
