package main

import "fmt"
import "math"

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5
	fmt.Printf("The circle's area is %.2fcm²\n", c.area())
	fmt.Println("New values:")
	fmt.Println(c.x, c.y, c.r)
	fmt.Printf("The new circle's area is %.2fcm²\n", c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Printf("The rectangle's area is %.2fcm²\n", r.area())
}
