package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func oop() {
	p := Point{3, 4}
	fmt.Println(p)
}

func (p *Point) Length() float64 {
	return math.Hypot((*p).X, (*p).Y)
}

func (p *Point) MoveXTo(newX float64) {
	(*p).X = newX
}

func (p Point) MoveYToValue(newY float64) {
	MoveYTo(&p, newY)
}

func MoveYTo(p *Point, newY float64) {
	(*p).Y = newY
}

func main() {
	oop()
	p := Point{3, 4}
	fmt.Println(p.Length())
	p.MoveXTo(6)
	fmt.Println(p.Length())
	p.MoveYToValue(7)
	fmt.Println(p.Length())

}
