package comparisons

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	由于浮点类型的计算非精确计算，所以可能产生大于、等于、小于、无法比较四种情况
	所以使用两个指令来对无法比较的情况进行定义
*/

type FCMPG struct {
	base.NoOperandsInstruction
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

func (e *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (e *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
