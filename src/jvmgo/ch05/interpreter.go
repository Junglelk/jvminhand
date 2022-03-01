package main

import (
	"fmt"
	"jvmgo/jvmgo/ch05/classfile"
	"jvmgo/jvmgo/ch05/instructions"
	"jvmgo/jvmgo/ch05/instructions/base"
	"jvmgo/jvmgo/ch05/rtda"
)

// 参数为方法字段表。
func interpret(methodInfo *classfile.MemberInfo) {
	// 获取code属性
	codeAttr := methodInfo.CodeAttribute()

	// 获取执行方法需要的局部变量表和操作数栈空间以及方法中的字节码
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	// 先创建一个Thread实例，然后创建一个帧，并把它推入Java虚拟机栈顶，最后执行方法
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, bytecode)
}

// 目前尚未实现Java的return指令，所以解释器运行一定会出错，所以解释器逻辑要转到catchErr()函数
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)
		// decode
		reader.Reset(bytecode, pc)
		// 获取操作码
		opcode := reader.ReadUint8()
		// 获取指令
		inst := instructions.NewInstruction(opcode)
		// 获取操作数
		inst.FetchOperands(reader)
		// 设置下一条指令
		frame.SetNextPC(reader.PC())
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		// 指令执行
		inst.Execute(frame)
	}
}
