package comparisons

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type IFEQ struct {
	base.BranchInstruction
}
type IFNE struct {
	base.BranchInstruction
}
type IFLT struct {
	base.BranchInstruction
}
type IFLE struct {
	base.BranchInstruction
}
type IFGT struct {
	base.BranchInstruction
}
type IFGE struct {
	base.BranchInstruction
}

func (e *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, e.Offset)
	}
}

func (e *IFNE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val != 0 {
		base.Branch(frame, e.Offset)
	}
}

func (e *IFLT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val < 0 {
		base.Branch(frame, e.Offset)
	}
}

func (e *IFLE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val <= 0 {
		base.Branch(frame, e.Offset)
	}
}

func (e *IFGT) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val > 0 {
		base.Branch(frame, e.Offset)
	}
}

func (e *IFGE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val >= 0 {
		base.Branch(frame, e.Offset)
	}
}
