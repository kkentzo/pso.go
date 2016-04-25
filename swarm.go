package pso

import (
	"fmt"
)

type Swarm struct {
	settings  *SwarmSettings
	gbest     []float64
	fitness   float64
	particles []*Particle
}

func Initialize(settings *SwarmSettings) Swarm {
	var swarm Swarm
	// store settings
	swarm.settings = settings
	// initialize gbest array
	swarm.gbest = make([]float64, settings.Function.dim)
	swarm.fitness = 1e20
	// initialize particles
	swarm.particles = make([]*Particle, settings.Function.dim)
	for i := 0; i < swarm.settings.Function.dim; i++ {
		swarm.particles[i] = NewParticle(settings)
		swarm.updateBest(swarm.particles[i])
	}
	return swarm
}

func (swarm *Swarm) updateBest(particle *Particle) {
	if particle.best < swarm.fitness {
		swarm.fitness = particle.best
		for i := 0; i < swarm.settings.Function.dim; i++ {
			swarm.gbest[i] = particle.pbest[i]
		}
	}
}

func (swarm *Swarm) Run() {
	// the algorithm goes here
	for step := 0; step < swarm.settings.Steps; step++ {
		for _, particle := range swarm.particles {
			particle.Update(swarm.gbest)
			swarm.updateBest(particle)
			if swarm.fitness < swarm.settings.Function.goal {
				fmt.Printf("Goal was reached @ step %d (fitness=%.2e) :-)",
					step, swarm.fitness)
				return
			}
		}
		if step%swarm.settings.PrintEvery == 0 {
			fmt.Printf("Step %d :: min err=%.5e\n", step, swarm.fitness)
		}
	}
	fmt.Printf("Finished; Goal was not reached (fitness=%.2e) :-(", swarm.fitness)

}
