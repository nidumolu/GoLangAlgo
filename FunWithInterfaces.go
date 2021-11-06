package main

import (
	"fmt"
)

type Friend interface {
	SayHello()
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

func Greet(f Friend) {
	f.SayHello()
}
func main() {
	var guy = new(Person)
	guy.Name = "Tom"
	Greet(guy)
	var dog = new(Dog)
	Greet(dog)
}
