package main

import "fmt"

type shape interface {

	// methods
	TotalArea() float64
	Volume() float64
}

type Cylinder struct {
	radius float64
	height float64
}

func (m *Cylinder) TotalArea() float64 {
	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m *Cylinder) Volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

// Main Method
func main() {

	// Accessing elements of
	// the tank interface
	//var t shape

	//var t = new(Cylinder)
	t := new(Cylinder)
	t.radius = 10.0
	t.height = 14.0
	fmt.Println("Area of tank :", t.TotalArea())
	fmt.Println("Volume of tank:", t.Volume())
}
