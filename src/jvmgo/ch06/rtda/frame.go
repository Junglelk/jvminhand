package rtda

// Frame lower 字段用来实现链表数据结构，localVars 字段保存局部变量指针，operandStack 字段保存操作数指针
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPc       int
}

/*
	执行方法所需的局部变量表大小和操作数栈的深度是由编译器预先计算好的，存储在 class 文件 method_info 结构的 Code 属性中

	Java虚拟机栈的链表结构为：
	+-------+		  +-------+			+-------+		  +-------+
	| Stack	|	+---->| Frame |	  +---->| Frame	|	+---->| Frame |
	+-------+	|	  +-------+	  |		+-------+	|	  +-------+
	| _top	|---+	  | lower |---+		| lower	|---+	  | lower  |
	+-------+		  +-------+	   		+-------+		  +-------+
*/
// 参数类型一致时，可以只写一个
func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// LocalVars getters
func (e *Frame) LocalVars() LocalVars {
	return e.localVars
}
func (e *Frame) OperandStack() *OperandStack {
	return e.operandStack
}

func (e *Frame) Thread() *Thread {
	return e.thread
}

func (e *Frame) SetThread(thread *Thread) {
	e.thread = thread
}

func (e *Frame) NextPC() int {
	return e.nextPc
}

func (e *Frame) SetNextPC(nextPc int) {
	e.nextPc = nextPc
}
