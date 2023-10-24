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

// Get the slope of the line between two points
func getSlope(p1 point, p2 point) float64 {
	var rise = p1.y - p2.y
	var run = p1.x - p2.x

	if run == 0 {
		return 0
	}
	println("Slope:", rise/run)
	return rise / run
}

// Everything seems to be done in radians
func angleOfSlope(boidCenter point, swarmCenter point) float64 {
	var slope = getSlope(boidCenter, swarmCenter)
	var angle = math.Atan(slope)
	if boidCenter.x-swarmCenter.x > 0 {
		return angle + (180 * math.Pi / 180)
	}
	return angle
}
