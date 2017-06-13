package pso

type Vector []float64

type Position struct {
	Location Vector
	Fitness  float64
	settings *SwarmSettings
}

func NewPosition(settings *SwarmSettings) *Position {
	loc := make([]float64, settings.Function.dim)
	return &Position{
		Location: loc,
		Fitness:  settings.Function.Evaluate(loc),
		settings: settings,
	}
}

func RandomPosition(settings *SwarmSettings) *Position {
	pos := NewPosition(settings)
	x_lo := settings.Function.x_lo
	x_hi := settings.Function.x_hi
	for i := 0; i < len(pos.Location); i++ {
		pos.Location[i] = x_lo + (x_hi-x_lo)*settings.rng.Float64()
	}
	pos.Fitness = settings.Function.Evaluate(pos.Location)
	return pos
}

func (position *Position) Update(delta Vector) {
	for i := 0; i < len(position.Location); i++ {
		position.Location[i] += delta[i]
	}
	position.Fitness = position.settings.Function.Evaluate(position.Location)
}

func (position *Position) Copy() *Position {
	newPosition := NewPosition(position.settings)
	copy(newPosition.Location, position.Location)
	newPosition.Fitness = position.Fitness
	return newPosition
}

func (position *Position) IsBetterThan(other *Position) bool {
	return position.Fitness < other.Fitness
}
