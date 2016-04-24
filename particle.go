package pso

type Particle struct {
	position []float64
	velocity []float64
	pbest []float64
	fitness float64
	best float64
	settings *Settings
}

func NewParticle(settings *Settings) *Particle {
	particle := new(Particle)
	particle.settings = settings
	particle.position = make([]float64, settings.Dim)
	particle.pbest = make([]float64, settings.Dim)
	particle.velocity = make([]float64, settings.Dim)

	for i:=0; i<settings.Dim; i++ {
		//fmt.Printf("%.2f %.2f\n", settings.rng.Float64(), settings.rng.Float64())
		a := settings.x_lo + (settings.x_hi - settings.x_lo) * settings.rng.Float64()
		b := settings.x_lo + (settings.x_hi - settings.x_lo) * settings.rng.Float64()

		particle.position[i] = a
		particle.pbest[i] = a
		particle.velocity[i] = (a - b) / 2.0
	}

	particle.fitness = settings.ObjectiveFunction(particle.position)
	particle.best = particle.fitness

	return particle
}

func (particle *Particle) Update(gbest []float64) {
	settings := particle.settings
	for i:=0; i<settings.Dim; i++ {
		// calculate stochastic coefficients
		rho1 := settings.c1 * settings.rng.Float64()
		rho2 := settings.c2 * settings.rng.Float64()
		// update velocity
		particle.velocity[i] =
			settings.w * particle.velocity[i] +
		    rho1 * (particle.pbest[i] - particle.position[i]) +
		    rho2 * (gbest[i] - particle.position[i])
		// update position
		particle.position[i] += particle.velocity[i];
	}

	// update particle fitness
	particle.fitness = settings.ObjectiveFunction(particle.position)
	// update personal best position?
	if particle.fitness < particle.best {
		particle.best = particle.fitness
		for i:=0; i<settings.Dim; i++ {
			particle.pbest[i] = particle.position[i]
		}
	}
}
