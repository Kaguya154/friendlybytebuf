package friendlybytebuf

import "math"

func (f *FriendlyByteBuf) WriteFloat32(value float32) (int, error) {
	bits := math.Float32bits(value)
	return f.WriteUint32(bits)
}
func (f *FriendlyByteBuf) ReadFloat32() (float32, error) {
	bits, err := f.ReadUint32()
	if err != nil {
		return 0, err
	}
	value := math.Float32frombits(bits)
	return value, nil
}
func (f *FriendlyByteBuf) WriteFloat64(value float64) (int, error) {
	bits := math.Float64bits(value)
	return f.WriteUint64(bits)
}
func (f *FriendlyByteBuf) ReadFloat64() (float64, error) {
	bits, err := f.ReadUint64()
	if err != nil {
		return 0, err
	}
	value := math.Float64frombits(bits)
	return value, nil
}
