package main

import "fmt"
import "math"

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

type Circle struct {
	x, y, r float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.r, 2)
}

func (c Circle) perimeter() float64 {
	return math.Pi * 2 * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func (r Rectangle) perimeter() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return (2 * l) + (2 * w)
}

func main() {
	var c Shape = Circle{x: 10, y: 5, r: 5}
	var r Shape = Rectangle{0, 0, 10, 10}
	fmt.Println(totalArea(c, r))
	fmt.Printf("The perimeter of the circle is %.2fcm\n", c.perimeter())
	fmt.Printf("The perimeter of the rectangle is %.2fcm\n", r.perimeter())
}
