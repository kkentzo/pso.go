package pso

import "math"

func Sphere(vec []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(vec); i++ {
		sum += math.Pow(vec[i], 2.0)
	}
	return sum
}

func Rosenbrock(vec []float64) float64 {
	var sum float64 = 0
	for i := 0; i < len(vec)-1; i++ {
		sum += 100.0*
			math.Pow((vec[i+1]-math.Pow(vec[i], 2.0)), 2.0) +
			math.Pow((1-vec[i]), 2.0)
	}
	return sum
}

func Griewank(vec []float64) float64 {
	var sum float64 = 0
	var prod float64 = 1

	for i := 0; i < len(vec); i++ {
		sum += math.Pow(vec[i], 2.0)
		prod *= math.Cos(vec[i] / math.Sqrt(float64(i+1)))
	}
	return sum/4000.0 - prod + 1.0
}
