Particle Swarm Optimization (PSO) in Go
===

An implementation of the Particle Swarm Optimization (PSO) algorithm
[1,2] in Go. PSO is used for problems involving global stochastic
optimization of a continuous function (called the objective
function). PSO can also be used for discrete optimization problems,
but this behaviour is not implemented in the current version of this
library.

The library implements a simple PSO which keeps track of a single
global best position for the swarm (i.e. no fancy neighborhood
strategies etc.) and maintains a constant inertia weight.

It is actually a rewrite of [PSO in C](https://github.com/kkentzo/pso)
for the purpose of learning Go :-)

## Usage

Install using:

```
go get github.com/kkentzo/pso.go
```

The library can be used as shown in the following code snippet:

```go
package main

import (
    "github.com/kkentzo/pso.go"
)

func main() {
    settings := pso.DefaultSettings()
    settings.ObjectiveFunction = pso.Rosenbrock
    swarm := pso.Initialize(settings)
    swarm.Run()
}
```

## TODO

- Implement linearly decreasing inertia weight
- Implement neighborhood information strategies
- Investigate particle parallelism using goroutines
- Implement specific settings for objective functions
- Add tests!
