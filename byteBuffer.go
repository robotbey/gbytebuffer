package gbytebuffer

import (
	"bytes"
	"encoding/binary"
)

// VERSION is the version of the package.
const VERSION = "0.2.0"

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
func (b *ByteBuffer) ReadByte() byte {
	defer b.Advance(1)
	return b.data[b.index]
}

// SetByte sets the byte at the current index.
func (b *ByteBuffer) WriteByte(val byte) {
	defer b.Advance(1)
	b.data = append(b.data, val)
	b.UpdateLength()
}

// ReadInt8 returns the int8 at the current index.
func (b *ByteBuffer) ReadInt8() int8 {
	return int8(b.ReadByte())
}

// WriteInt8 sets the int8 at the current index.
func (b *ByteBuffer) WriteInt8(val int8) {
	b.WriteByte(byte(val))
}

// ReadUInt8 returns the uint8 at the current index.
func (b *ByteBuffer) ReadUInt8() uint8 {
	return uint8(b.ReadByte())
}

// WriteUInt8 sets the uint8 at the current index.
func (b *ByteBuffer) WriteUInt8(val uint8) {
	b.WriteByte(byte(val))
}

// ReadInt16 returns the int16 at the current index.
func (b *ByteBuffer) ReadInt16() int16 {
	return int16(b.ReadUInt16())
}

func (b *ByteBuffer) WriteInt16(val int16) {
	b.WriteInt16(int16(val))
}

func (b *ByteBuffer) ReadInt16BE() int16 {
	return int16(b.ReadUInt16BE())
}

func (b *ByteBuffer) WriteInt16BE(val int16) {
	b.WriteInt16BE(int16(val))
}

func (b *ByteBuffer) ReadUInt16() uint16 {
	defer b.Advance(2)
	var c uint16
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) WriteUInt16(val uint16) {
	defer b.Advance(2)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadUInt16BE() uint16 {
	defer b.Advance(2)
	var c uint16
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) WriteUInt16BE(val uint16) {
	defer b.Advance(2)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadUInt32() uint32 {
	defer b.Advance(4)
	var c uint32
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) WriteUInt32(val uint32) {
	defer b.Advance(4)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadUInt32BE() uint32 {
	defer b.Advance(4)
	var c uint32
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) WriteUInt32BE(val uint32) {
	defer b.Advance(4)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadInt32() int32 {
	return int32(b.ReadUInt32())
}

func (b *ByteBuffer) WriteInt32(val int32) {
	b.WriteUInt32(uint32(val))
}

func (b *ByteBuffer) ReadInt32BE() int32 {
	return int32(b.ReadUInt32BE())
}

func (b *ByteBuffer) WriteInt32BE(val int32) {
	b.WriteUInt32BE(uint32(val))
}

func (b *ByteBuffer) ReadUInt64() uint64 {
	defer b.Advance(8)
	var c uint64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) WriteUInt64(val uint64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadUInt64BE() uint64 {
	defer b.Advance(8)
	var c uint64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) WriteUInt64BE(val uint64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadInt64() int64 {
	return int64(b.ReadUInt64())
}

func (b *ByteBuffer) WriteInt64(val int64) {
	b.WriteUInt64(uint64(val))
}

func (b *ByteBuffer) ReadInt64BE() int64 {
	return int64(b.ReadUInt64BE())
}

func (b *ByteBuffer) WriteInt64BE(val int64) {
	b.WriteUInt64BE(uint64(val))
}

func (b *ByteBuffer) ReadFloat64() float64 {
	defer b.Advance(8)
	var c float64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.LittleEndian, &c)
	return c
}

func (b *ByteBuffer) WriteFloat64(val float64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.LittleEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadFloat64BE() float64 {
	defer b.Advance(8)
	var c float64
	buf := bytes.NewReader(b.Data())
	binary.Read(buf, binary.BigEndian, &c)
	return c
}

func (b *ByteBuffer) WriteFloat64BE(val float64) {
	defer b.Advance(8)
	var buffer bytes.Buffer
	binary.Write(&buffer, binary.BigEndian, &val)
	b.data = append(b.data, buffer.Bytes()...)
	b.UpdateLength()
}

func (b *ByteBuffer) ReadBytes(size int) []byte {
	defer b.Advance(size)
	return b.Data()[:size]
}

func (b *ByteBuffer) WriteBytes(data []byte) {
	defer b.Advance(len(data))
	b.data = append(b.data, data...)
	b.UpdateLength()
}
