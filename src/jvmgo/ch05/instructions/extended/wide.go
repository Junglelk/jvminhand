package extended

import (
	"jvmgo/jvmgo/ch05/instructions/base"
	loads2 "jvmgo/jvmgo/ch05/instructions/loads"
	"jvmgo/jvmgo/ch05/instructions/math"
	stores2 "jvmgo/jvmgo/ch05/instructions/stores"
	"jvmgo/jvmgo/ch05/rtda"
)

/*
	加载类指令、存储类指令、ret指令和iinc指令需要按索引访问局部变量表，索引以uint8的形式存在字节码中。
	一般来说，方法的局部变量表不会超过 256 ，所以一个字节就够用了，如果有的方法的局部变量表超过这个限制时，
	就需要使用wide指令来扩展该指令。
*/
// WIDE wide指令改变其他指令的行为，modifiedInstruction字段存放被改变的指令。
type WIDE struct {
	modifiedInstruction base.Instruction
}

func (e *WIDE) FetchOperands(reader *base.BytecodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	// iload
	case 0x15:
		inst := &loads2.ILOAD{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// lload
	case 0x16:
		inst := &loads2.LLOAD{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// fload
	case 0x17:
		inst := &loads2.FLOAD{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// dload
	case 0x18:
		inst := &loads2.DLOAD{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// aload
	case 0x19:
		inst := &loads2.ALOAD{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// istore
	case 0x36:
		inst := &stores2.ISTORE{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// lstore
	case 0x37:
		inst := &stores2.LSTORE{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// fstore
	case 0x38:
		inst := &stores2.FSTORE{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// dstore
	case 0x39:
		inst := &stores2.DSTORE{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// astore
	case 0x3a:
		inst := &stores2.ASTORE{}
		inst.Index = uint(reader.ReadUint16())
		e.modifiedInstruction = inst
	// iinc
	case 0x84:
		inst := &math.IINC{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
		e.modifiedInstruction = inst
	case 0xa9: // ret
		panic("Unsupported opcode 0xa9!")
	}
}

// Execute wide仅作指令扩展，所以直接调用扩展指令的Execute方法即可
func (e *WIDE) Execute(frame *rtda.Frame) {
	e.modifiedInstruction.Execute(frame)
}
