package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

// push方法把栈帧压入栈顶
func (e *Stack) push(frame *Frame) {
	if e.size >= e.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if e._top != nil {
		frame.lower = e._top
	}
	e._top = frame
	e.size++
}

// pop方法把栈顶元素弹出
func (e *Stack) pop() *Frame {
	if e._top == nil {
		panic("jvm stack is empty")
	}
	// 看起来像是个双向链表
	top := e._top
	e._top = top.lower
	top.lower = nil
	e.size--
	return top
}

func (e *Stack) top() *Frame {
	if e._top == nil {
		panic("jvm stack is empty")
	}
	return e._top
}
