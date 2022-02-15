package control

import (
	"jvmgo/jvmgo/ch06/instructions/base"
	"jvmgo/jvmgo/ch06/rtda"
)

/*
	java 中的switch-case语句有两种实现方式：
		1. 如果case可以被编码成一个索引表，则实现成tableswitch指令；
		2. 否则实现成lookupswitch指令。
	比如这个就会被编译成tableswitch：

	int chooseNear(int i ){
		switch (i) {
			case 0 : return 0;
			case 1 : return 1;
			case 2 : return 2;
			default: return -1;
		}
	}

	这个则会被编码成lookupswitch
	int chooseNear(int i ){
		switch (i) {
			case -100 : return -1;
			case 0 : return 0;
			case 100 : return 1;
			default: return -1;
		}
	}
*/

type TABLE_SWITCH struct {
	defaultOffset int32
	low           int32
	high          int32
	jumpOffsets   []int32
}

// FetchOperands tableswitch指令操作码之后有 0~3 字节的padding，以保证defaultOffset在字节码中的地址为 4 的倍数。
func (e *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	e.defaultOffset = reader.ReadInt32()
	e.low = reader.ReadInt32()
	e.high = reader.ReadInt32()
	jumpOffsetCount := e.high - e.low + 1
	e.jumpOffsets = reader.ReadUint32s(jumpOffsetCount)
}

// Execute 先从操作数栈中弹出一个int变量，然后看它是否在low和high给定范围内，如果在，则从jumpOffset表中查出偏移量进行跳转，否则按照defaultOffset跳转。
func (e *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	var offset int
	if index >= e.low && index <= e.high {
		offset = int(e.jumpOffsets[index-e.low])
	} else {
		offset = int(e.defaultOffset)
	}
	base.Branch(frame, offset)
}
