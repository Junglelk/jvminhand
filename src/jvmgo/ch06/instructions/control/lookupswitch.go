package control

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

// FetchOperands matchOffsets 的索引是case的值，索引值是跳转偏移量。
func (e *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	e.defaultOffset = reader.ReadInt32()
	e.npairs = reader.ReadInt32()
	e.matchOffsets = reader.ReadUint32s(e.npairs * 2)
}

func (e *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	// int 和 uint 在Go语言中其实是不确定的，代表了计算机的字长，可以为32位，也可以为64位。
	for i := int32(0); i < e.npairs*2; i += 2 {
		if e.matchOffsets[i] == key {
			offset := e.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(e.defaultOffset))
}
