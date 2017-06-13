package pso

type Particle struct {
	position *Position
	pbest    *Position
	velocity Vector
	settings *SwarmSettings
}

func NewParticle(settings *SwarmSettings) *Particle {
	particle := new(Particle)
	particle.settings = settings
	particle.position = RandomPosition(settings)
	particle.pbest = particle.position.Copy()

	particle.velocity = make([]float64, settings.Function.dim)
	x_lo := settings.Function.x_lo
	x_hi := settings.Function.x_hi
	for i := 0; i < settings.Function.dim; i++ {
		a := x_lo + (x_hi-x_lo)*settings.rng.Float64()
		b := x_lo + (x_hi-x_lo)*settings.rng.Float64()
		particle.velocity[i] = (a - b) / 2.0
	}
	return particle
}

func (particle *Particle) Update(gbest *Position) {
	settings := particle.settings
	for i := 0; i < settings.Function.dim; i++ {
		// calculate stochastic coefficients
		rho1 := settings.C1 * settings.rng.Float64()
		rho2 := settings.C2 * settings.rng.Float64()
		// update velocity
		particle.velocity[i] =
			settings.inertia*particle.velocity[i] +
				rho1*(particle.pbest.Location[i]-particle.position.Location[i]) +
				rho2*(gbest.Location[i]-particle.position.Location[i])
	}
	// update position
	particle.position.Update(particle.velocity)

	// update personal best position?
	if particle.position.IsBetterThan(particle.pbest) {
		particle.pbest = particle.position.Copy()
	}
}
