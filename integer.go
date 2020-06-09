package main

type IntNumber interface {
	Equal(j Integer) bool
	LessThan(j Integer) bool
	MoreThan(j Integer) bool
	Increase()
	Decrease()
}

type IntNumber2 interface {
	Equal(j Integer) bool
	LessThan(j Integer) bool
	MoreThan(j Integer) bool
}

type Integer int

func (i Integer) Equal(j Integer) bool {
	return i == j
}

func (i Integer) LessThan(j Integer) bool {
	return i < j
}

func (i Integer) MoreThan(j Integer) bool {
	return i > j
}

func (i *Integer) Increase(){
	*i += 1
}

func (i *Integer) Decrease(){
	*i -= 1
}

