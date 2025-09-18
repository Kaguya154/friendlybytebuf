package friendlybytebuf

func (f *FriendlyByteBuf) WriteByteArray(data []byte) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	_, err := f.WriteBytes(data)
	return err
}

func (f *FriendlyByteBuf) ReadByteArray() ([]byte, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]byte, length)
	_, err = f.buf.Read(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteIntArray(data []int) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteVarInt(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadIntArray() ([]int, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]int, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadVarInt()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteStringArray(data []string) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, s := range data {
		if _, err := f.WriteString(s); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadStringArray() ([]string, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]string, length)
	for i := 0; i < length; i++ {
		s, err := f.ReadString()
		if err != nil {
			return nil, err
		}
		data[i] = s
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteFloat32Array(data []float32) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteFloat32(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadFloat32Array() ([]float32, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]float32, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadFloat32()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}
func (f *FriendlyByteBuf) WriteFloat64Array(data []float64) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteFloat64(v); err != nil {
			return err
		}
	}
	return nil
}
func (f *FriendlyByteBuf) ReadFloat64Array() ([]float64, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]float64, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadFloat64()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteBoolArray(data []bool) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if err := f.WriteBool(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadBoolArray() ([]bool, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]bool, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadBool()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteUintArray(data []uint) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteUint(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadUintArray() ([]uint, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]uint, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadUint()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteInt64Array(data []int64) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteInt64(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadInt64Array() ([]int64, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]int64, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadInt64()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteUint64Array(data []uint64) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteUint64(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadUint64Array() ([]uint64, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]uint64, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadUint64()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteUint32Array(data []uint32) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteUint32(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadUint32Array() ([]uint32, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]uint32, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadUint32()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteUint16Array(data []uint16) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteUint16(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadUint16Array() ([]uint16, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]uint16, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadUint16()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteUint8Array(data []uint8) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if err := f.WriteUint8(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadUint8Array() ([]uint8, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]uint8, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadUint8()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteInt8Array(data []int8) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if err := f.WriteInt8(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadInt8Array() ([]int8, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]int8, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadInt8()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteInt16Array(data []int16) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteInt16(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadInt16Array() ([]int16, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]int16, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadInt16()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}

func (f *FriendlyByteBuf) WriteInt32Array(data []int32) error {
	if _, err := f.WriteVarInt(len(data)); err != nil {
		return err
	}
	for _, v := range data {
		if _, err := f.WriteInt32(v); err != nil {
			return err
		}
	}
	return nil
}

func (f *FriendlyByteBuf) ReadInt32Array() ([]int32, error) {
	length, err := f.ReadVarInt()
	if err != nil {
		return nil, err
	}
	data := make([]int32, length)
	for i := 0; i < length; i++ {
		v, err := f.ReadInt32()
		if err != nil {
			return nil, err
		}
		data[i] = v
	}
	return data, nil
}
