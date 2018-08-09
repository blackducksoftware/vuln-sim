package main

import (
	"time"

	"github.com/jayunit100/vuln-sim/pkg/model2"
) // <- ui shortcut, optional

func main() {
	c := &model2.ClusterSim{
		ChurnProbability: .10,
		MaxPodsPerApp:    10,
		NumUsers:         100,
		TimePeriod:       1 * time.Minute,
	}
	model2.Simulate(c)
}
