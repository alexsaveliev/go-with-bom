package main

type Foo struct {
}

func (self Foo) Bar() {
	println("привет");
}

func main() {
	foo := Foo{}
	foo.Bar()
}