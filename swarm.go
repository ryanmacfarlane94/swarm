package main

// Center of mass for the swarm
type swarm struct {
	centerX  float64
	centerY  float64
	rotation float64
	boids    []boid
}

func newSwarm(x float64, y float64, rotation float64) swarm {
	var boids []boid

	boids = append(boids,
		newBoid(-20, 10, 45),
		newBoid(-20, -20, 45),
		newBoid(20, 220, 45),
	)

	swarm := swarm{
		centerX:  x,
		centerY:  y,
		rotation: rotation,
		boids:    boids,
	}

	return swarm
}

func moveSwarm(swarm *swarm) {
	swarm.centerX += 2
	swarm.centerY += 2
}
