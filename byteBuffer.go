package gbytebuffer

import (
	"bytes"
	"encoding/binary"
)

// ByteBuffer is used to read and write data to a byte array.
type ByteBuffer struct {
	data   []byte
	index  int
	length int
}

// ByteBufferFrom creates a new ByteBuffer from a byte array.
func ByteBufferFrom(data []byte) *ByteBuffer {
	return &ByteBuffer{
		data:   data,
		index:  0,
		length: len(data),
	}
}

// NewByteBuffer creates a new ByteBuffer.
func NewByteBuffer() *ByteBuffer {
	return &ByteBuffer{
		data:   make([]byte, 0),
		index:  0,
		length: 0,
	}
}

// Index returns the current index.
func (b ByteBuffer) Index() int {
	return b.index
}

// Length returns the length of the buffer.
func (b ByteBuffer) Length() int {
	return b.length
}

// Buffer returns the underlying byte array.
func (b ByteBuffer) Buffer() []byte {
	return b.data
}

// Data returns a slice of the underlying byte array starting at the current index.
func (b ByteBuffer) Data() []byte {
	return b.data[b.index:]
}

// Advance moves the index forward by the specified amount.
func (b *ByteBuffer) Advance(len int) int {
	b.index += len
	return b.index
}

// Seek sets the current index to the specified value.
func (b *ByteBuffer) Seek(pos int) int {
	b.index = pos
	return b.index
}

// UpdateLength updates the length of the buffer.
func (b *ByteBuffer) UpdateLength() {
	b.length = len(b.data)
}

// GetByte returns the byte at the current index.
func (b *ByteBuffer) GetByte() byte {
	defer b.Advance(1)
	return b.data[b.index]
}

// SetByte sets the byte at the current index.
func (b *ByteBuffer) SetByte(val byte) {
	defer b.Advance(1)
	b.data = append(b.data, val)
	b.UpdateLength()
}

// GetInt8 returns the int8 at the current index.
func (b *ByteBuffer) GetInt8() int8 {
	return int8(b.GetByte())
}

// SetInt8 sets the int8 at the current index.
func (b *ByteBuffer) SetInt8(val int8) {
	b.SetByte(byte(val))
}

// GetUInt8 returns the uint8 at the current index.
func (b *ByteBuffer) GetUInt8() uint8 {
	return uint8(b.GetByte())
}

// SetUInt8 sets the uint8 at the current index.
func (b *ByteBuffer) SetUInt8(val uint8) {
	b.SetByte(byte(val))
}

// GetInt16 returns the int16 at the current index.
func (b *ByteBuffer) GetInt16() int16 {
	return int16(b.GetUInt16())
}

func (b *ByteBuffer) SetInt16(val int16) {
	b.SetInt16(int16(val))
}

func (b *ByteBuffer) GetInt16BE() int16 {
	return int16(b.GetUInt16BE())
}

func (b *ByteBuffer) SetInt16BE(val int16) {
	b.SetInt16BE(int16(val))
}

func (b *ByteBuffer) GetUInt16() uint16 {
	defer b.Advance(2)
	var c uint16
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) SetUInt16(val uint16) {
	defer b.Advance(2)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetUInt16BE() uint16 {
	defer b.Advance(2)
	var c uint16
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) SetUInt16BE(val uint16) {
	defer b.Advance(2)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetUInt32() uint32 {
	defer b.Advance(4)
	var c uint32
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) SetUInt32(val uint32) {
	defer b.Advance(4)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetUInt32BE() uint32 {
	defer b.Advance(4)
	var c uint32
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) SetUInt32BE(val uint32) {
	defer b.Advance(4)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetInt32() int32 {
	return int32(b.GetUInt32())
}

func (b *ByteBuffer) SetInt32(val int32) {
	b.SetUInt32(uint32(val))
}

func (b *ByteBuffer) GetInt32BE() int32 {
	return int32(b.GetUInt32BE())
}

func (b *ByteBuffer) SetInt32BE(val int32) {
	b.SetUInt32BE(uint32(val))
}

func (b *ByteBuffer) GetUInt64() uint64 {
	defer b.Advance(8)
	var c uint64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) SetUInt64(val uint64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetUInt64BE() uint64 {
	defer b.Advance(8)
	var c uint64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) SetUInt64BE(val uint64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetInt64() int64 {
	return int64(b.GetUInt64())
}

func (b *ByteBuffer) SetInt64(val int64) {
	b.SetUInt64(uint64(val))
}

func (b *ByteBuffer) GetInt64BE() int64 {
	return int64(b.GetUInt64BE())
}

func (b *ByteBuffer) SetInt64BE(val int64) {
	b.SetUInt64BE(uint64(val))
}

func (b *ByteBuffer) GetFloat64() float64 {
	defer b.Advance(8)
	var c float64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) SetFloat64(val float64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetFloat64BE() float64 {
	defer b.Advance(8)
	var c float64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) SetFloat64BE(val float64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) GetBytes(size int) []byte {
	defer b.Advance(size)
	return b.Data()[:size]
}
