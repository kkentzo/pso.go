package main

import (
	"flag"
	"fmt"

	"github.com/kkentzo/pso.go"
)

func main() {
	function := flag.String("function", "sphere",
		"function to optimize (options: sphere|rosenbrock|griewank)")
	printEvery := flag.Int("print", 1000, "step frequency of output")
	flag.Parse()
	fmt.Printf("Optimizing function %s\n", *function)
	settings := pso.DefaultSettings()
	settings.PrintEvery = *printEvery
	switch *function {
	case "sphere":
		settings.Function = pso.Sphere
	case "rosenbrock":
		settings.Function = pso.Rosenbrock
	case "griewank":
		settings.Function = pso.Griewank
	}
	swarm := pso.NewSwarm(settings)
	result := swarm.Run()
	if result.Position.Fitness < settings.Function.Goal {
		fmt.Printf("Yay! Goal was reached @ step %d (fitness=%.2e) :-)",
			result.Step, result.Position.Fitness)
	} else {
		fmt.Printf("Goal was not reached after %d steps (fitness=%.2e) :-)",
			result.Step, result.Position.Fitness)

	}
}
