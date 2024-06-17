package main

import "fmt"

type rect struct {
	width, height int
}

// Methods can be defined for either pointer or value receiver types.
// this area method has a receiver type of *rect.
func (r *rect) area() int {
	return r.width * r.height
}

// Here’s an example of a value receiver.
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("area: ", rp.perim())

}
