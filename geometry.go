package main

import "math"

type point struct {
	x float64
	y float64
}

func newPoint(x float64, y float64) point {
	point := point{
		x: x,
		y: y,
	}

	return point
}

func rotatePoint(point *point, angle float64) {
	var x = point.x
	var y = point.y
	point.x = (x * math.Cos(angle)) - (y * math.Sin(angle))
	point.y = (y * math.Cos(angle)) + (x * math.Sin(angle))
}
