package comparisons

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (e *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	v1, v2 := _cmpPop(frame)
	if v1 == v2 {
		base.Branch(frame, e.Offset)
	}
}
func (e *IF_ICMPNE) Execute(frame *rtda.Frame) {
	v1, v2 := _cmpPop(frame)
	if v1 != v2 {
		base.Branch(frame, e.Offset)
	}
}
func (e *IF_ICMPLT) Execute(frame *rtda.Frame) {
	v1, v2 := _cmpPop(frame)
	if v1 < v2 {
		base.Branch(frame, e.Offset)
	}
}
func (e *IF_ICMPLE) Execute(frame *rtda.Frame) {
	v1, v2 := _cmpPop(frame)
	if v1 <= v2 {
		base.Branch(frame, e.Offset)
	}
}
func (e *IF_ICMPGT) Execute(frame *rtda.Frame) {
	v1, v2 := _cmpPop(frame)
	if v1 > v2 {
		base.Branch(frame, e.Offset)
	}
}
func (e *IF_ICMPGE) Execute(frame *rtda.Frame) {
	v1, v2 := _cmpPop(frame)
	if v1 >= v2 {
		base.Branch(frame, e.Offset)
	}
}

func _cmpPop(frame *rtda.Frame) (v1, v2 int32) {
	stack := frame.OperandStack()
	v2 = stack.PopInt()
	v1 = stack.PopInt()
	return
}
