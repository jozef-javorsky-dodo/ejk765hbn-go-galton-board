package main

import (
	"fmt"

	"math"

	"math/rand"

	"strconv"
)

type Statistics struct {
	Mean float64

	StdDev float64

	Variance float64
}

func (s Statistics) String() string {

	return fmt.Sprintf("\n•••> mean: %.2f\n•••> std-dev: %.2f\n•••> variance: %.2f\n", s.Mean, s.StdDev, s.Variance)

}

func getInput(prompt string, defaultValue int) int {

	fmt.Print(prompt)

	var input string

	_, err := fmt.Scanln(&input)

	if err != nil {

		return defaultValue

	}

	val, err := strconv.Atoi(input)

	if err != nil {
		return defaultValue
	}

	return val

}

func simulate(rows, balls int) []int {

	distribution := make([]int, rows+1)

	for i := 0; i < balls; i++ {

		bin := 0

		for j := 0; j < rows; j++ {

			if rand.Intn(2) == 1 {

				bin++

			}

		}

		distribution[bin]++

	}

	return distribution

}

func calculateStatistics(distribution []int) Statistics {

	var sum, sumSq, totalBalls float64

	var mean, stdDev, variance float64

	for i, freq := range distribution {

		totalBalls += float64(freq)

		sum += float64(i) * float64(freq)

	}

	if totalBalls > 0 {

		mean = sum / totalBalls

		for i, freq := range distribution {

			sumSq += float64(freq) * math.Pow(float64(i)-mean, 2)

		}

		variance = sumSq / totalBalls

		stdDev = math.Sqrt(variance)

	}

	return Statistics{Mean: mean, StdDev: stdDev, Variance: variance}

}

func visualize(distribution []int, stats Statistics) {

	maxCount := 0

	for _, count := range distribution {

		if count > maxCount {

			maxCount = count

		}

	}

	const barWidth = 40

	for i, count := range distribution {

		barLength := 0

		if maxCount > 0 {

			barLength = int((float64(count) / float64(maxCount)) * barWidth)

		}

		isOutlier := float64(i) < stats.Mean-(1.5*stats.StdDev) || float64(i) > stats.Mean+(1.5*stats.StdDev)

		bar := ""

		for j := 0; j < barLength; j++ {

			if isOutlier {

				bar += "~"

			} else {

				bar += "#"

			}

		}

		fmt.Printf("{ [...(%d).]: %s (%d) }\n", i, bar, count)

	}

}

func main() {

	rows := getInput("\nEnter the number of rows (default 9): ", 9)

	balls := getInput("Enter the number of balls (default 256): ", 256)

	distribution := simulate(rows, balls)

	stats := calculateStatistics(distribution)

	fmt.Printf("Galton Board Simulation with %d balls and %d rows.\n\n\n", balls, rows)

	visualize(distribution, stats)

	fmt.Printf("\n\nStatistics: %v\n", stats)
}
