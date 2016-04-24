package pso

import (
	"time"
	"math/rand"
)

const pso_max_size int = 100
const pso_inertia float64 = 0.7298 // default value of w (see clerc02)

type Settings struct {
    dim int // problem dimensionality
    x_lo float64 // lower range limit
    x_hi float64 // higher range limit
    goal float64 // optimization goal (error threshold)

    size int // swarm size (number of particles)
    print_every int // ... N steps (set to 0 for no output)
    steps int // maximum number of iterations
    step int // current PSO step
    c1 float64 // cognitive coefficient
    c2 float64 // social coefficient
    w_max float64 // max inertia weight value
    w_min float64 // min inertia weight value
	w float64 // current inertia weight value

    clamp_pos bool // whether to keep particle position within defined bounds (TRUE)
	// or apply periodic boundary conditions (FALSE)

	rng *rand.Rand // the random number generator

	ObjectiveFunction func(vec []float64)float64 // the objective function
}

func DefaultSettings() *Settings {
	settings := new(Settings)

	settings.dim = 30
	settings.x_lo =  -20
	settings.x_hi =  20
	settings.goal = 1e-5;
    settings.size = CalculateSwarmSize(settings.dim, pso_max_size);
    settings.print_every = 1000;
    settings.steps = 100000;
    settings.c1 = 1.496;
    settings.c2 = 1.496;
    settings.w_max = pso_inertia;
    settings.w_min = 0.3;
	settings.w = settings.w_max;

    settings.clamp_pos = true;

	// initialize the RNG
	settings.rng = rand.New(rand.NewSource(time.Now().UnixNano()))

	settings.ObjectiveFunction = Sphere

	return settings
}
