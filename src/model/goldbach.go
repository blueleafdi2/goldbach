package model

type Goldbach struct {
	Id       int64
	Seq      int64
	Num      int64
	A        int64
	B        int64
	AAndB    int64
	AOrB     int64
	AXorB    int64
	AAndNotB int64
	NumByte  int64
	AByte    int64
	BByte    int64
	ATri     string
	BTri     string
	ASix     string
	BSix     string
}
