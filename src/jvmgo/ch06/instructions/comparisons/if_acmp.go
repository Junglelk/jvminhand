package comparisons

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (e *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v2 == v1 {
		base.Branch(frame, e.Offset)
	}
}

func (e *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopRef()
	v1 := stack.PopRef()
	if v2 != v1 {
		base.Branch(frame, e.Offset)
	}
}
