package point

import "math"

type point struct {
	x, y float64
}

func Point(x, y float64) point {
	return point{x, y}
}

func (p1 point) Distance(p2 point) float64 {
	sX := p2.x - p1.x
	sY := p2.y - p1.y
	return math.Sqrt(sX*sX + sY*sY)
}
