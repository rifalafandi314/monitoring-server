package services

import "math"

var cpuHistory []float64
const window = 30
const k = 2.0

func AdaptiveThreshold(value float64) (string, float64) {
	cpuHistory = append(cpuHistory, value)
	if len(cpuHistory) > window {
		cpuHistory = cpuHistory[1:]
	}

	if len(cpuHistory) < 10 {
		return "LEARNING", 0
	}

	sum := 0.0
	for _, v := range cpuHistory {
		sum += v
	}
	mean := sum / float64(len(cpuHistory))

	variance := 0.0
	for _, v := range cpuHistory {
		variance += (v - mean) * (v - mean)
	}
	std := math.Sqrt(variance / float64(len(cpuHistory)))

	threshold := mean + k*std

	status := "NORMAL"
	if value > threshold {
		status = "CRITICAL"
	} else if value > mean+std {
		status = "WARNING"
	}

	return status, threshold
}