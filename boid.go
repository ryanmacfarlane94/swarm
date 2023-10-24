package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type boid struct {
	center   point
	rotation float64
	points   []point
}

func newBoid(x float64, y float64, rotation float64) boid {
	var points []point

	points = append(
		points,
		newPoint(8, 0),
		newPoint(-8, 4),
		newPoint(-4, 0),
		newPoint(-8, -4),
	)

	boid := boid{
		center:   newPoint(x, y),
		rotation: rotation,
		points:   points,
	}

	return boid
}

func drawBoid(imd *imdraw.IMDraw, boid boid) {
	imd.Color = colornames.Blueviolet

	for index := range boid.points {
		rotatePoint(&boid.points[index], boid.rotation)
	}

	for index := range boid.points {
		imd.EndShape = imdraw.SharpEndShape

		imd.Push(
			pixel.V(boid.points[index].x+boid.center.x, boid.points[index].y+boid.center.y),
			pixel.V(boid.points[(index+1)%len(boid.points)].x+boid.center.x, boid.points[(index+1)%len(boid.points)].y+boid.center.y),
		)
		imd.Line(2)
	}
	for index := range boid.points {
		rotatePoint(&boid.points[index], -boid.rotation)
	}
}

func moveBoid(boid *boid, swarm swarm) {
	changeTrajectory(boid, swarm)
}

// Change trajectory towards a point
func changeTrajectory(boid *boid, swarm swarm) {
	var angle = angleOfSlope(boid.center, swarm.center)
	boid.rotation = angle

	var slope = getSlope(boid.center, swarm.center)
	var rise = slope
	var run = 1.0

	boid.center.x += rise
	boid.center.y += run
}
