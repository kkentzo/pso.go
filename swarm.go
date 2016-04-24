package pso

import (
	"fmt"
)

type Swarm struct {
	settings *Settings
	gbest[] float64
	fitness float64
	particles []*Particle
}

func Initialize(settings *Settings) Swarm {
	var swarm Swarm
	// store settings
	swarm.settings = settings
	// initialize gbest array
	swarm.gbest = make([]float64, settings.dim)
	swarm.fitness = 1e20
	// initialize particles
	swarm.particles = make([]*Particle, settings.dim)
	for i:=0; i<swarm.settings.dim; i++ {
		swarm.particles[i] = NewParticle(settings)
		swarm.updateBest(swarm.particles[i])
	}
	return swarm
}

func (swarm *Swarm) updateBest(particle *Particle) {
	if particle.best < swarm.fitness {
		swarm.fitness = particle.best
		for i:=0; i<swarm.settings.dim; i++ {
			swarm.gbest[i] = particle.pbest[i]
		}
	}
}

func (swarm *Swarm) Run() {
	// the algorithm goes here
	for step:=0; step<swarm.settings.steps; step++ {
		for _, particle := range(swarm.particles) {
			particle.Update(swarm.gbest)
			swarm.updateBest(particle)
			//fmt.Printf("f = %.5f\n", particle.fitness)
			if swarm.fitness < swarm.settings.goal {
				fmt.Printf("Goal was reached @ step %d :-)", step)
				return
			}
		}
		if step % swarm.settings.print_every == 0 {
			fmt.Printf("Step %d :: min err=%.5e\n", step, swarm.fitness)
		}
	}
}
