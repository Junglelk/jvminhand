package base

import "jvmgo/jvmgo/ch05/rtda"

/*
	Java解释器的大致逻辑如下：
		do {
			atomically calculate pc and fetch opcode at pc;
			if (operands) fetch operands;
			execute the action for the opcode;
		} while (there is more to do)
	每次循环包含三个部分：计算pc、指令解码、指令执行。不难看出这个逻辑可以用for循环和switch-case来实现，
	但几乎无可读性，也不符合正常的编码逻辑。
	所以采用另一种方式：将指令抽象成接口，代码和逻辑写在具体的指令实现中。
	for {
		pc := calculatePC()
		opcode := bytecode[pc]
		inst := createInst(opcode)
		inst.fetchOperands(bytecode)
		inst.execute()
	}
*/
// 指令接口
type Instruction interface {
	// FetchOperands 从字节码中提取操作数
	FetchOperands(reader *BytecodeReader)
	// Execute 执行指令逻辑
	Execute(frame *rtda.Frame)
}

// NoOperandsInstruction 代表无操作数的指令，所以不定义任何字段
type NoOperandsInstruction struct{}

func (e NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// nothing
}

// BranchInstruction 表示跳转指令，offset表示偏移量
type BranchInstruction struct {
	Offset int
}

func (e *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	e.Offset = int(reader.ReadInt16())
}

// Index8Instruction 存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。
// 将其抽象成Index8Instruction结构体，用Index字段表示局部变量索引。FetchOperands()方法从字节码中读取一个int8的整数，转成uint后赋给Index字段
type Index8Instruction struct {
	Index uint
}

func (e *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	e.Index = uint(reader.ReaderUint8())
}

// Index16Instruction 有些指令需要访问运行时常量池，常量池索引由两字节操作数给出
type Index16Instruction struct {
	Index uint
}

func (e *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	e.Index = uint(reader.ReaderUint16())
}
