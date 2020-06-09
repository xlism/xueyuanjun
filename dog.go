package main


type Dog struct {
	name string
	animal Animal
}

func (d Dog) FavorFood() string{
	return "gt"
}

func (d Dog) Call() string{
	return "ww"
}