package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (e Slots) SetInt(index uint, val int32) {
	e[index].num = val
}
func (e Slots) GetInt(index uint) int32 {
	return e[index].num
}

func (e Slots) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	e[index].num = int32(bits)
}
func (e Slots) GetFloat(index uint) float32 {
	bits := uint32(e[index].num)
	return math.Float32frombits(bits)
}

func (e Slots) SetLong(index uint, val int64) {
	e[index].num = int32(val)
	e[index+1].num = int32(val >> 32)
}
func (e Slots) GetLong(index uint) int64 {
	low := uint32(e[index].num)
	high := uint32(e[index+1].num)
	return int64(high)<<32 | int64(low)
}

func (e Slots) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	e.SetLong(index, int64(bits))
}
func (e Slots) GetDouble(index uint) float64 {
	bits := uint64(e.GetLong(index))
	return math.Float64frombits(bits)
}

func (e Slots) SetRef(index uint, ref *Object) {
	e[index].ref = ref
}
func (e Slots) GetRef(index uint) *Object {
	return e[index].ref
}
