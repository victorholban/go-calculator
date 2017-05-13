package calculator

import (
	"encoding/binary"
	"bytes"
)


func Add(nr1, nr2 int32) int32 {
	return nr1 + nr2
}

func Difference(a, b int32) int32 {
	return a - b
}

func Int32ToBytes (a int32) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, a)
	return buf.Bytes()
}

func GetBitFromPosition(bytes []byte, position uint) bool {
	selectedByte := bytes[position/8]
	mask := byte(1 << (position%8))
	return (selectedByte & mask) != 0
}