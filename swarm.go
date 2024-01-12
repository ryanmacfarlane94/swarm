package main

// Center of mass for the swarm
type swarm struct {
	center   point
	rotation float64
	boids    []boid
}

func newSwarm(x float64, y float64, rotation float64) swarm {
	var boids []boid

	boids = append(boids,
		newBoid(40, 40, 0),
		newBoid(40, -40, 0),
		newBoid(-40, 40, 0),
		newBoid(-40, -40, 0),
	)

	swarm := swarm{
		center:   newPoint(x, y),
		rotation: rotation,
		boids:    boids,
	}

	return swarm
}

func moveSwarm(swarm *swarm) {
	changeSwarmTrajectory(swarm)
}

// Change trajectory towards a point
func changeSwarmTrajectory(swarm *swarm) {
	swarm.rotation -= 0.005
	swarm.center = nextPoint(swarm.center, swarm.rotation, 3)
}
