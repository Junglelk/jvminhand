package rtda

// 实现运行时数据区

// Thread 线程，pc 即程序计数器，stack 为Java虚拟机栈指针。
// 和堆一样，Java虚拟机栈可以是连续的空间，也可以不连续；可以固定大小，也可以在运行时动态扩展，
// 如果拓展超过了虚拟机栈的限制，会抛 StackOverflowError 异常，
// 如果内存已经耗尽导致Java虚拟机栈不能拓展，会抛出 OutOfMemoryError 异常。
type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	// TODO 此处可以修改命令行参数来修改大小
	return &Thread{stack: newStack(1 * 1024)}
}

// PC getter方法
func (e *Thread) PC() int {
	return e.pc
}

// SetPC setter方法
func (e *Thread) SetPC(pc int) {
	e.pc = pc
}

func (e *Thread) PushFrame(frame *Frame) {
	e.stack.push(frame)
}

func (e *Thread) PopFrame() *Frame {
	return e.stack.pop()
}
func (e *Thread) CurrentFrame() *Frame {
	return e.stack.top()
}

func (e *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return NewFrame(e, maxLocals, maxStack)
}
