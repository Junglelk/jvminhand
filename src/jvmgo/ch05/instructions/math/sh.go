package math

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

type ISHL struct {
	base.NoOperandsInstruction
}

type ISHR struct {
	base.NoOperandsInstruction
}

type IUSHR struct {
	base.NoOperandsInstruction
}

type LSHL struct {
	base.NoOperandsInstruction
}

type LSHR struct {
	base.NoOperandsInstruction
}

type LUSHR struct {
	base.NoOperandsInstruction
}

// Execute 左移指令，先从操作数栈中弹出两个int变量，v1 是要进行位移的变量，v2 是要位移多少比特
// 由于int只有 32 位，所以取 v2 的前 5 个比特就足以表示位移的位数
// go 语言位移操作符右侧必须是无符号整数
func (e *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// 0x1f 二进制为 0001 1111
	s := uint32(v2) & 0x1f
	// v1 左移s
	result := v1 << s
	stack.PushInt(result)
}

func (e *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// 0x3f 二进制为 0011 1111 因为需要拓展符号位
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushInt(result)
}

// Execute 无符号右移
func (e *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (e *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 << s
	stack.PushLong(result)
}

func (e *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := v1 >> s
	stack.PushLong(result)
}

// Execute 无符号右移
func (e *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}
