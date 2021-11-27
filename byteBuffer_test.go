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
			is := b.GetByte()
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
			b.SetByte(tt.want)
			if b.Buffer()[0] != tt.want {
				t.Errorf("ByteBuffer.SetByte() = %v, want %v", b.Buffer()[0], tt.want)
			}
		})
	}
}

func TestByteBuffer_GetInt8(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetInt8(); got != tt.want {
				t.Errorf("ByteBuffer.GetInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetInt8(t *testing.T) {
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
			tt.b.SetInt8(tt.args.val)
		})
	}
}

func TestByteBuffer_GetUInt8(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetUInt8(); got != tt.want {
				t.Errorf("ByteBuffer.GetUInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetUInt8(t *testing.T) {
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
			tt.b.SetUInt8(tt.args.val)
		})
	}
}

func TestByteBuffer_GetInt16(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetInt16(); got != tt.want {
				t.Errorf("ByteBuffer.GetInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetInt16(t *testing.T) {
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
			tt.b.SetInt16(tt.args.val)
		})
	}
}

func TestByteBuffer_GetInt16BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetInt16BE(); got != tt.want {
				t.Errorf("ByteBuffer.GetInt16BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetInt16BE(t *testing.T) {
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
			tt.b.SetInt16BE(tt.args.val)
		})
	}
}

func TestByteBuffer_GetUInt16(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetUInt16(); got != tt.want {
				t.Errorf("ByteBuffer.GetUInt16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetUInt16(t *testing.T) {
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
			tt.b.SetUInt16(tt.args.val)
		})
	}
}

func TestByteBuffer_GetUInt16BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint16
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetUInt16BE(); got != tt.want {
				t.Errorf("ByteBuffer.GetUInt16BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetUInt16BE(t *testing.T) {
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
			tt.b.SetUInt16BE(tt.args.val)
		})
	}
}

func TestByteBuffer_GetUInt32(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetUInt32(); got != tt.want {
				t.Errorf("ByteBuffer.GetUInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetUInt32(t *testing.T) {
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
			tt.b.SetUInt32(tt.args.val)
		})
	}
}

func TestByteBuffer_GetUInt32BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetUInt32BE(); got != tt.want {
				t.Errorf("ByteBuffer.GetUInt32BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetUInt32BE(t *testing.T) {
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
			tt.b.SetUInt32BE(tt.args.val)
		})
	}
}

func TestByteBuffer_GetInt32(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetInt32(); got != tt.want {
				t.Errorf("ByteBuffer.GetInt32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetInt32(t *testing.T) {
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
			tt.b.SetInt32(tt.args.val)
		})
	}
}

func TestByteBuffer_GetInt32BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetInt32BE(); got != tt.want {
				t.Errorf("ByteBuffer.GetInt32BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetInt32BE(t *testing.T) {
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
			tt.b.SetInt32BE(tt.args.val)
		})
	}
}

func TestByteBuffer_GetUInt64(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetUInt64(); got != tt.want {
				t.Errorf("ByteBuffer.GetUInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetUInt64(t *testing.T) {
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
			tt.b.SetUInt64(tt.args.val)
		})
	}
}

func TestByteBuffer_GetUInt64BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetUInt64BE(); got != tt.want {
				t.Errorf("ByteBuffer.GetUInt64BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetUInt64BE(t *testing.T) {
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
			tt.b.SetUInt64BE(tt.args.val)
		})
	}
}

func TestByteBuffer_GetInt64(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetInt64(); got != tt.want {
				t.Errorf("ByteBuffer.GetInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetInt64(t *testing.T) {
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
			tt.b.SetInt64(tt.args.val)
		})
	}
}

func TestByteBuffer_GetInt64BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetInt64BE(); got != tt.want {
				t.Errorf("ByteBuffer.GetInt64BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetInt64BE(t *testing.T) {
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
			tt.b.SetInt64BE(tt.args.val)
		})
	}
}

func TestByteBuffer_GetFloat64(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetFloat64(); got != tt.want {
				t.Errorf("ByteBuffer.GetFloat64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetFloat64(t *testing.T) {
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
			tt.b.SetFloat64(tt.args.val)
		})
	}
}

func TestByteBuffer_GetFloat64BE(t *testing.T) {
	tests := []struct {
		name string
		b    *ByteBuffer
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.GetFloat64BE(); got != tt.want {
				t.Errorf("ByteBuffer.GetFloat64BE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteBuffer_SetFloat64BE(t *testing.T) {
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
			tt.b.SetFloat64BE(tt.args.val)
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
			if got := tt.b.GetBytes(tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ByteBuffer.GetBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
