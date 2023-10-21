package main

import (
	"fmt"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/fstanis/screenresolution"
	"golang.org/x/image/colornames"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	resolution := screenresolution.GetPrimary()
	fmt.Println(resolution)

	cfg := pixelgl.WindowConfig{
		Title:  "Swarm",
		Bounds: pixel.R(0, 0, float64(resolution.Width), float64(resolution.Height)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)
	swarm := newSwarm(0, 0, 45)

	for !win.Closed() {
		imd.Clear()

		moveSwarm(&swarm)
		imd.Push(pixel.V(swarm.centerX, swarm.centerY))
		imd.Circle(5, 0)

		for index := range swarm.boids {
			moveBoid(&swarm.boids[index])
			drawBoid(imd, swarm.boids[index])
		}

		win.Clear(colornames.Black)
		imd.Draw(win)
		win.Update()
	}
}
