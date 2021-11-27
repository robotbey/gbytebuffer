# gbytebuffer

A byte buffer library to read/write simple variable types from/into byte slices. Conversions happen in little endian format by default but there is an big endian version of each function. Supported types are:

- int8/byte
- int16/uint16
- int32/uint32
- int64/uint64
- float64

## Usage

```go
    package main

    import (
        "github.com/robotbey/gbytebuffer"
        "fmt"
    )

    func main(){

        // Write simple variable into byte slice
        buffer1 := NewByteBuffer()
        i8var := int8(42)
        buffer.WriteInt8(i8var)
        fmt.Println(buffer.Buffer())
        // output: [42]

        // Read from buffer into a variable
        buffer2 := ByteBufferFrom([]byte{0x69, 0xAB, 0xCD, 0xDE, 0xAD, 0xC0, 0xDE})   // 0xDEADC0DE = 3735929054
        i8var := buffer2.ReadInt8()
        i16var := uffer2.ReadInt16()
        i32var := uffer2.ReadInt32()

        fmt.Println(i8var)
        fmt.Println(i16var)
        fmt.Println(i32var)
        // output: 105
        //         43981
        //         3735929054

    }

```
