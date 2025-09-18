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
