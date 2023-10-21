package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type boid struct {
	centerX  float64
	centerY  float64
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
		centerX:  x,
		centerY:  y,
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
			pixel.V(boid.points[index].x+boid.centerX, boid.points[index].y+boid.centerY),
			pixel.V(boid.points[(index+1)%len(boid.points)].x+boid.centerX, boid.points[(index+1)%len(boid.points)].y+boid.centerY),
		)
		imd.Line(2)
	}
	for index := range boid.points {
		rotatePoint(&boid.points[index], -boid.rotation)
	}
}

func moveBoid(boid *boid) {
	boid.centerX += 2
	boid.centerY += 2
}
