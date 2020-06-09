package main

type Number1 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number2 interface {
	Equal(i int) bool
	LessThan(i int) bool
	MoreThan(i int) bool
}

type Number int

func (n Number) Equal(i int) bool{
	return int(n) == i
}

func (n Number) LessThan(i int) bool{
	return int(n) < i
}

func (n Number) MoreThan(i int) bool{
	return int(n) > i
}
