package pso

type Particle struct {
	position []float64
	velocity []float64
	pbest    []float64
	fitness  float64
	best     float64
	settings *SwarmSettings
}

func NewParticle(settings *SwarmSettings) *Particle {
	particle := new(Particle)
	particle.settings = settings
	particle.position = make([]float64, settings.Function.dim)
	particle.pbest = make([]float64, settings.Function.dim)
	particle.velocity = make([]float64, settings.Function.dim)

	x_lo := settings.Function.x_lo
	x_hi := settings.Function.x_hi

	for i := 0; i < settings.Function.dim; i++ {
		a := x_lo + (x_hi-x_lo)*settings.rng.Float64()
		b := x_lo + (x_hi-x_lo)*settings.rng.Float64()

		particle.position[i] = a
		particle.pbest[i] = a
		particle.velocity[i] = (a - b) / 2.0
	}

	particle.fitness = settings.Function.Evaluate(particle.position)
	particle.best = particle.fitness

	return particle
}

func (particle *Particle) Update(gbest []float64) {
	settings := particle.settings
	for i := 0; i < settings.Function.dim; i++ {
		// calculate stochastic coefficients
		rho1 := settings.C1 * settings.rng.Float64()
		rho2 := settings.C2 * settings.rng.Float64()
		// update velocity
		particle.velocity[i] =
			settings.inertia*particle.velocity[i] +
				rho1*(particle.pbest[i]-particle.position[i]) +
				rho2*(gbest[i]-particle.position[i])
		// update position
		particle.position[i] += particle.velocity[i]
	}

	// update particle fitness
	particle.fitness = settings.Function.Evaluate(particle.position)
	// update personal best position?
	if particle.fitness < particle.best {
		particle.best = particle.fitness
		for i := 0; i < settings.Function.dim; i++ {
			particle.pbest[i] = particle.position[i]
		}
	}
}
