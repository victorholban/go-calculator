package calculator

//import "encoding/binary"


func Add(nr1, nr2 int32) int32 {
	//a := make([]byte, 4)
	//b := make([]byte, 4)
	//
	//binary.BigEndian.PutUint32(a, nr1)
	//binary.BigEndian.PutUint32(b, nr2)

	return nr1 + nr2
}

func Difference(a, b int32) int32 {
	return a - b
}
