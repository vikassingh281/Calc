package NormalCalc

type Nos struct {
	No1 int
	No2 int
}

func (n Nos) Add() int {
	return n.No2 + n.No2
}
