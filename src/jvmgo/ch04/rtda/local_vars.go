package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

/*
	由于操作局部变量表和操作数栈的指令都是隐含类型信息的,所以需要给LocalVars类型定义一些方法，用来存取不同类型的变量。
*/
// SetInt 存取 int
func (e LocalVars) SetInt(index uint, val int32) {
	e[index].num = val
}

func (e LocalVars) GetInt(index uint) int32 {
	return e[index].num
}

// SetFloat 先将 float 转为 int，然后再按照 int 处理
func (e LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	e[index].num = int32(bits)
}

func (e LocalVars) getFloat(index uint) float32 {
	bits := uint32(e[index].num)
	return math.Float32frombits(bits)
}

// SetLong long类型按照两个 int 处理。
func (e LocalVars) SetLong(index uint, val int64) {
	// 将 64 位的 int 转为 32 位的，直接取低 32 位
	e[index].num = int32(val)
	// 将该数左移 32 位后，原高 32 位移到低 32 位，高 32 位全为 0，再转，就是正常的 32 位
	e[index+1].num = int32(val >> 32)
}

func (e LocalVars) GetLong(index uint) int64 {
	low := uint32(e[index].num)
	high := uint32(e[index+1].num)
	// 喵啊
	return int64(high)<<32 | int64(low)
}

// SetDouble double 可以转成 long 类型再进行存取
func (e LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	e.SetLong(index, int64(bits))
}

func (e LocalVars) GetDouble(index uint) float64 {
	bits := uint64(e.GetLong(index))
	return math.Float64frombits(bits)
}

// SetRef 引用类型的直接存取
func (e LocalVars) SetRef(index uint, ref *Object) {
	e[index].ref = ref
}

func (e LocalVars) GetRef(index uint) *Object {
	return e[index].ref
}
