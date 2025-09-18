package friendlybytebuf

import "bytes"

type FriendlyByteBuf struct {
	buf bytes.Buffer
}

// ErrVarIntOverflow error
type ErrVarIntOverflow struct{}

func (e ErrVarIntOverflow) Error() string {
	return "varint overflow"
}
