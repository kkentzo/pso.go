package pso

import (
	"fmt"
)

type Result struct {
	Position
	Step int
}

type Swarm struct {
	Gbest     *Position
	settings  *SwarmSettings
	particles []*Particle
}

func NewSwarm(settings *SwarmSettings) *Swarm {
	var swarm Swarm
	// store settings
	swarm.settings = settings
	// initialize gbest
	swarm.Gbest = NewPosition(settings)
	swarm.Gbest.Fitness = 1e20
	// initialize particles
	swarm.particles = make([]*Particle, settings.Function.dim)
	for i := 0; i < swarm.settings.Function.dim; i++ {
		swarm.particles[i] = NewParticle(settings)
		swarm.updateBest(swarm.particles[i])
	}
	return &swarm
}

func (swarm *Swarm) updateBest(particle *Particle) {
	if particle.pbest.IsBetterThan(swarm.Gbest) {
		swarm.Gbest = particle.pbest.Copy()
	}
}

func (swarm *Swarm) Run() Result {
	// the algorithm goes here
	var step int
	for step = 0; step < swarm.settings.Steps; step++ {
		for _, particle := range swarm.particles {
			particle.Update(swarm.Gbest)
			swarm.updateBest(particle)
			if swarm.Gbest.Fitness < swarm.settings.Function.Goal {
				return Result{
					Position: *swarm.Gbest,
					Step:     step,
				}
			}
		}
		if step%swarm.settings.PrintEvery == 0 {
			fmt.Printf("Step %d :: min err=%.5e\n", step, swarm.Gbest.Fitness)
		}
	}
	return Result{
		Position: *swarm.Gbest,
		Step:     step,
	}

}
