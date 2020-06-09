package main

type Animal struct {
	name string
}

func (a Animal) Call() string {
	return "call"
}

func (a Animal) FavorFood() string {
	return "food"
}

func (a Animal) GetName() string {
	return a.name
}


