package pointer

import (
    "math"
)


type Point struct {
    x float64
    y float64
}


func NewPoint(x, y float64) *Point {
    return &Point{x: x, y: y}
}


func (p *Point) getX() float64 {
    return p.x
}

func (p *Point) getY() float64 {
    return p.y
}


func Distance(p1, p2 *Point) float64 {
	x1 := p1.getX()
	y1 := p1.getY()
	x2 := p2.getX()
	y2 := p2.getY()
	ans := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
    return ans
}
