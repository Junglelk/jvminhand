package extended

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

// GOTO_W 与GOTO指令的最大区别是索引从2字节变为了4字节
type GOTO_W struct {
	offset int
}

func (e *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	e.offset = int(reader.ReadInt32())
}

func (e *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, e.offset)
}
