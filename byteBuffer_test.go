package gbytebuffer

import (
	"bytes"
	"reflect"
	"testing"
)

func TestByteBuffer_ByteBufferFrom(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want []byte
	}{
		{"", ByteBufferFrom([]byte{0}), []byte{0}},
		{"", ByteBufferFrom([]byte{42}), []byte{42}},
		{"", ByteBufferFrom([]byte{4, 8, 15, 16, 23, 42}), []byte{4, 8, 15, 16, 23, 42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.b.Length() != len(tt.want) {
				t.Errorf("Expected len %d, got %d", len(tt.want), tt.b.Length())
			}
			if !bytes.Equal(tt.b.Data(), tt.want) {
				t.Errorf("Expected data to be %v, got %v", tt.want, tt.b.Data())
			}
		})
	}
}

func TestByteBuffer_Index(t *testing.T) {
	tests := []struct {
		name string
		b    ByteBuffer
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Index(); got != tt.want {
				t.Errorf("ByteBuffer.Index() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_Length(t *testing.T) {
	tests := []struct {
		name string
		b    ByteBuffer
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Length(); got != tt.want {
				t.Errorf("ByteBuffer.Length() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_Data(t *testing.T) {
	tests := []struct {
		name string
		b    ByteBuffer
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Data(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteBuffer.Data() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_Advance(t *testing.T) {
	type args struct {
		len int
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Advance(tt.args.len); got != tt.want {
				t.Errorf("ByteBuffer.Advance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_Seek(t *testing.T) {
	type args struct {
		pos int
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.Seek(tt.args.pos); got != tt.want {
				t.Errorf("ByteBuffer.Seek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_UpdateLength(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.UpdateLength()
		})
	}
}

func TestByteBuffer_GetByte(t *testing.T) {
	tests := []struct {
		name string
		want byte
	}{
		{"", 0},
		{"", 1},
		{"", 42},
		{"", 255},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := ByteBufferFrom([]byte{tt.want})
			is := b.ReadByte()
			if is != tt.want {
				t.Errorf("ByteBuffer.BetByte() = %v, want %v", is, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetByte(t *testing.T) {
	tests := []struct {
		name string
		want byte
	}{
		{"", 0},
		{"", 1},
		{"", 42},
		{"", 255},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewByteBuffer()
			b.WriteByte(tt.want)
			if b.Buffer()[0] != tt.want {
				t.Errorf("ByteBuffer.SetByte() = %v, want %v", b.Buffer()[0], tt.want)
			}
		})
	}
}

func TestByteBuffer_ReadInt8(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadInt8(); got != tt.want {
				t.Errorf("ByteBuffer.ReadInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteInt8(t *testing.T) {
	type args struct {
		val int8
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteInt8(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadUInt8(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadUInt8(); got != tt.want {
				t.Errorf("ByteBuffer.ReadUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteUInt8(t *testing.T) {
	type args struct {
		val uint8
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteUInt8(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadInt16(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadInt16(); got != tt.want {
				t.Errorf("ByteBuffer.ReadInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteInt16(t *testing.T) {
	type args struct {
		val int16
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteInt16(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadInt16BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadInt16BE(); got != tt.want {
				t.Errorf("ByteBuffer.ReadInt16BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteInt16BE(t *testing.T) {
	type args struct {
		val int16
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteInt16BE(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadUInt16(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadUInt16(); got != tt.want {
				t.Errorf("ByteBuffer.ReadUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteUInt16(t *testing.T) {
	type args struct {
		val uint16
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteUInt16(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadUInt16BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadUInt16BE(); got != tt.want {
				t.Errorf("ByteBuffer.ReadUInt16BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteUInt16BE(t *testing.T) {
	type args struct {
		val uint16
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteUInt16BE(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadUInt32(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadUInt32(); got != tt.want {
				t.Errorf("ByteBuffer.ReadUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteUInt32(t *testing.T) {
	type args struct {
		val uint32
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteUInt32(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadUInt32BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadUInt32BE(); got != tt.want {
				t.Errorf("ByteBuffer.ReadUInt32BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteUInt32BE(t *testing.T) {
	type args struct {
		val uint32
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteUInt32BE(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadInt32(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadInt32(); got != tt.want {
				t.Errorf("ByteBuffer.ReadInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteInt32(t *testing.T) {
	type args struct {
		val int32
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteInt32(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadInt32BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadInt32BE(); got != tt.want {
				t.Errorf("ByteBuffer.ReadInt32BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteInt32BE(t *testing.T) {
	type args struct {
		val int32
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteInt32BE(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadUInt64(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadUInt64(); got != tt.want {
				t.Errorf("ByteBuffer.ReadUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteUInt64(t *testing.T) {
	type args struct {
		val uint64
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteUInt64(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadUInt64BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadUInt64BE(); got != tt.want {
				t.Errorf("ByteBuffer.ReadUInt64BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteUInt64BE(t *testing.T) {
	type args struct {
		val uint64
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteUInt64BE(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadInt64(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadInt64(); got != tt.want {
				t.Errorf("ByteBuffer.ReadInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteInt64(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteInt64(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadInt64BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadInt64BE(); got != tt.want {
				t.Errorf("ByteBuffer.ReadInt64BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteInt64BE(t *testing.T) {
	type args struct {
		val int64
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteInt64BE(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadFloat64(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadFloat64(); got != tt.want {
				t.Errorf("ByteBuffer.ReadFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteFloat64(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteFloat64(tt.args.val)
		})
	}
}

func TestByteBuffer_ReadFloat64BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadFloat64BE(); got != tt.want {
				t.Errorf("ByteBuffer.ReadFloat64BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_WriteFloat64BE(t *testing.T) {
	type args struct {
		val float64
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.b.WriteFloat64BE(tt.args.val)
		})
	}
}

func TestByteBuffer_GetBytes(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		b    *ByteBuffer
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.ReadBytes(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteBuffer.GetBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
