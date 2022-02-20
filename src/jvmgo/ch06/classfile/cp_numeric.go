package classfile

import (
	"math"
)

// CONSTANT_Integer_info 使用 4 字节存储整数常量，结构体定义如下：
// CONSTANT_Integer_info{
// 	u1 tag;
// 	u4 bytes;
// }

// ConstantIntegerInfo CONSTANT_Integer_info正好可以容纳一个Java的 int 型的常量，
// 但实际上，比 int 更小的boolean、byte、short和char类型的常量也放在 CONSTANT_Integer_info 中
type ConstantIntegerInfo struct {
	val int32
}

// readInfo 先读取一个 uint32 的数据，再转为int32类型
func (e *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	e.val = int32(bytes)
}

func (e *ConstantIntegerInfo) Value() int32 {
	return e.val
}

// ConstantFloatInfo 32位浮点数常量
type ConstantFloatInfo struct {
	val float32
}

func (e *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	e.val = math.Float32frombits(bytes)
}

func (e *ConstantFloatInfo) Value() float32 {
	return e.val
}

type ConstantLongInfo struct {
	val int64
}

// readInfo 先读取一个 uint64 的数据，再转为int64类型
func (e *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	e.val = int64(bytes)
}

func (e *ConstantLongInfo) Value() int64 {
	return e.val
}

// ConstantDoubleInfo 64位浮点数常量
type ConstantDoubleInfo struct {
	val float64
}

func (e *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	e.val = math.Float64frombits(bytes)
}

func (e *ConstantDoubleInfo) Value() float64 {
	return e.val
}
