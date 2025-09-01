package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(other Point) float64 {
	return math.Sqrt(math.Pow((p.x-other.x), 2) + math.Pow((p.y-other.y), 2))
}

func main() {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(2, 2)

	fmt.Println(p1, p2, p1.Distance(p2))

}
