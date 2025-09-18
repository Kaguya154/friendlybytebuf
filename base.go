package friendlybytebuf

import "bytes"

func Create() *FriendlyByteBuf {
	return &FriendlyByteBuf{
		buf: bytes.Buffer{},
	}
}

func FromBytes(data []byte) *FriendlyByteBuf {
	return &FriendlyByteBuf{
		buf: *bytes.NewBuffer(data),
	}
}

func FromBuffer(buf *bytes.Buffer) *FriendlyByteBuf {
	return &FriendlyByteBuf{
		buf: *buf,
	}
}

func (f *FriendlyByteBuf) Buffer() *bytes.Buffer {
	return &f.buf
}

func (f *FriendlyByteBuf) Bytes() []byte {
	return f.buf.Bytes()
}

func (f *FriendlyByteBuf) WriteBytes(data []byte) (int, error) {
	return f.buf.Write(data)
}

func (f *FriendlyByteBuf) WriteString(s string) (int, error) {
	return f.buf.WriteString(s + "\x00")
}

func (f *FriendlyByteBuf) ReadString() (string, error) {
	s, err := f.buf.ReadString(0)
	if err != nil {
		return "", err
	}
	return s[:len(s)-1], nil
}

func (f *FriendlyByteBuf) WriteByte(b byte) error {
	return f.buf.WriteByte(b)
}

func (f *FriendlyByteBuf) ReadByte() (byte, error) {
	return f.buf.ReadByte()
}

func (f *FriendlyByteBuf) WriteUint8(u uint8) error {
	return f.buf.WriteByte(byte(u))
}

func (f *FriendlyByteBuf) ReadUint8() (uint8, error) {
	b, err := f.buf.ReadByte()
	if err != nil {
		return 0, err
	}
	return uint8(b), nil
}

func (f *FriendlyByteBuf) WriteUint16(u uint16) (int, error) {
	return f.buf.Write([]byte{byte(u >> 8), byte(u & 0xff)})
}

func (f *FriendlyByteBuf) ReadUint16() (uint16, error) {
	b := make([]byte, 2)
	_, err := f.buf.Read(b)
	if err != nil {
		return 0, err
	}
	return uint16(b[0])<<8 | uint16(b[1]), nil
}

func (f *FriendlyByteBuf) WriteUint32(u uint32) (int, error) {
	return f.buf.Write([]byte{
		byte(u >> 24),
		byte((u >> 16) & 0xff),
		byte((u >> 8) & 0xff),
		byte(u & 0xff),
	})
}
func (f *FriendlyByteBuf) ReadUint32() (uint32, error) {
	b := make([]byte, 4)
	_, err := f.buf.Read(b)
	if err != nil {
		return 0, err
	}
	return uint32(b[0])<<24 | uint32(b[1])<<16 | uint32(b[2])<<8 | uint32(b[3]), nil
}
func (f *FriendlyByteBuf) WriteUint64(u uint64) (int, error) {
	return f.buf.Write([]byte{
		byte(u >> 56),
		byte((u >> 48) & 0xff),
		byte((u >> 40) & 0xff),
		byte((u >> 32) & 0xff),
		byte((u >> 24) & 0xff),
		byte((u >> 16) & 0xff),
		byte((u >> 8) & 0xff),
		byte(u & 0xff),
	})
}

func (f *FriendlyByteBuf) ReadUint64() (uint64, error) {
	b := make([]byte, 8)
	_, err := f.buf.Read(b)
	if err != nil {
		return 0, err
	}
	return uint64(b[0])<<56 | uint64(b[1])<<48 | uint64(b[2])<<40 | uint64(b[3])<<32 |
		uint64(b[4])<<24 | uint64(b[5])<<16 | uint64(b[6])<<8 | uint64(b[7]), nil
}

func (f *FriendlyByteBuf) WriteInt8(i int8) error {
	return f.WriteUint8(uint8(i))
}
func (f *FriendlyByteBuf) ReadInt8() (int8, error) {
	u, err := f.ReadUint8()
	return int8(u), err
}

func (f *FriendlyByteBuf) WriteInt16(i int16) (int, error) {
	return f.WriteUint16(uint16(i))
}
func (f *FriendlyByteBuf) ReadInt16() (int16, error) {
	u, err := f.ReadUint16()
	return int16(u), err
}
func (f *FriendlyByteBuf) WriteInt32(i int32) (int, error) {
	return f.WriteUint32(uint32(i))
}
func (f *FriendlyByteBuf) ReadInt32() (int32, error) {
	u, err := f.ReadUint32()
	return int32(u), err
}
func (f *FriendlyByteBuf) WriteInt64(i int64) (int, error) {
	return f.WriteUint64(uint64(i))
}
func (f *FriendlyByteBuf) ReadInt64() (int64, error) {
	u, err := f.ReadUint64()
	return int64(u), err
}

func (f *FriendlyByteBuf) WriteInt(i int) (int, error) {
	return f.WriteInt64(int64(i))
}
func (f *FriendlyByteBuf) ReadInt() (int, error) {
	i, err := f.ReadInt64()
	return int(i), err
}

func (f *FriendlyByteBuf) WriteUint(u uint) (int, error) {
	return f.WriteUint64(uint64(u))
}
func (f *FriendlyByteBuf) ReadUint() (uint, error) {
	u, err := f.ReadUint64()
	return uint(u), err
}

func (f *FriendlyByteBuf) WriteBool(b bool) error {
	if b {
		return f.WriteUint8(1)
	}
	return f.WriteUint8(0)
}
func (f *FriendlyByteBuf) ReadBool() (bool, error) {
	b, err := f.ReadUint8()
	if err != nil {
		return false, err
	}
	return b != 0, nil
}

func (f *FriendlyByteBuf) WriteVarInt(i int) (int, error) {
	var buf [5]byte
	var n int
	ux := uint(i) << 1
	if i < 0 {
		ux = ^ux
	}
	for ux >= 0x80 {
		buf[n] = byte(ux) | 0x80
		ux >>= 7
		n++
	}
	buf[n] = byte(ux)
	n++
	return f.buf.Write(buf[:n])
}

func (f *FriendlyByteBuf) ReadVarInt() (int, error) {
	var ux uint
	var s uint
	for i := 0; ; i++ {
		b, err := f.buf.ReadByte()
		if err != nil {
			return 0, err
		}
		if b < 0x80 {
			if i > 4 || i == 4 && b > 15 {
				return 0, ErrVarIntOverflow{}
			}
			ux |= uint(b) << s
			break
		}
		ux |= uint(b&0x7f) << s
		s += 7
	}
	x := int(ux >> 1)
	if ux&1 != 0 {
		x = ^x
	}
	return x, nil
}
