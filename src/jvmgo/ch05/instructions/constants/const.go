package constants

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	这一系列指令将隐含在操作码里的常量值推入操作数栈顶
*/

type ACONST_NULL struct {
	base.NoOperandsInstruction
}
type DCONST_0 struct {
	base.NoOperandsInstruction
}
type DCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_0 struct {
	base.NoOperandsInstruction
}
type FCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_M1 struct {
	base.NoOperandsInstruction
}
type ICONST_0 struct {
	base.NoOperandsInstruction
}
type ICONST_1 struct {
	base.NoOperandsInstruction
}
type ICONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_3 struct {
	base.NoOperandsInstruction
}
type ICONST_4 struct {
	base.NoOperandsInstruction
}
type ICONST_5 struct {
	base.NoOperandsInstruction
}
type LCONST_0 struct {
	base.NoOperandsInstruction
}
type LCONST_1 struct {
	base.NoOperandsInstruction
}

// Execute ACONST_NULL 将null压入栈中
func (e ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// Execute DCONST_0 将 double 型 0 压入栈中
func (e DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

// Execute DCONST_1 将 double 型 1 压入栈中
func (e DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// Execute FCONST_0 将 float 型的 0 压入栈中
func (e FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

// Execute FCONST_1 将 float 型的 1 压入栈中
func (e FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

// Execute FCONST_2 将 float 型的 2 压入栈中
func (e FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

// Execute ICONST_M1 将int类型的 -1 压入栈中
func (e ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

// Execute ICONST_0 将int类型的 0 压入栈中
func (e ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

// Execute ICONST_1 将int类型的 1 压入栈中
func (e ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

// Execute ICONST_2 将int类型的 2 压入栈中
func (e ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

// Execute ICONST_3 将int类型的 3 压入栈中
func (e ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

// Execute ICONST_4 将int类型的 4 压入栈中
func (e ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

// Execute ICONST_5 将int类型的 5 压入栈中
func (e ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// Execute LCONST_0 将long类型的 0 压入栈中
func (e LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

// Execute LCONST_1 将long类型的 1 压入栈中
func (e LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}
