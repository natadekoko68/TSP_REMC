package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	nCities   = 50
	nIter     = 1000000
	nM        = 200
	deltaBeta = 0.5
)

var cities [][2]float64
var T []float64

func init() {
	rand.Seed(time.Now().UnixNano())

	T = make([]float64, nM+1)
	for m := 0; m <= nM; m++ {
		T[m] = 1 / float64((m+1)*nM)
	}

	cities = make([][2]float64, nCities)
	for i := 0; i < nCities; i++ {
		cities[i] = [2]float64{rand.Float64(), rand.Float64()}
	}
}

func calcDist(i, j int) float64 {
	dx := cities[i][0] - cities[j][0]
	dy := cities[i][1] - cities[j][1]
	return math.Sqrt(dx*dx + dy*dy)
}

func calcDistFromPath(path []int) float64 {
	dist := 0.0
	for i := 0; i < len(path)-1; i++ {
		dist += calcDist(path[i], path[i+1])
	}
	dist += calcDist(path[len(path)-1], path[0])
	return dist
}

func main() {
	tempPath := make([][]int, nM)
	for i := 0; i < nM; i++ {
		tempPath[i] = make([]int, nCities)
		for j := 0; j < nCities; j++ {
			tempPath[i][j] = j
		}
	}

	for iter := 0; iter < nIter; iter++ {
		dists := make([]float64, nM)
		for m := 0; m < nM; m++ {
			distInit := calcDistFromPath(tempPath[m])

			k := rand.Intn(nCities - 1)
			l := rand.Intn(nCities-k-1) + k + 1
			tempPath[m][k], tempPath[m][l] = tempPath[m][l], tempPath[m][k]

			distFin := calcDistFromPath(tempPath[m])

			deltaR := distFin - distInit
			if deltaR < 0 || rand.Float64() < math.Exp(-deltaR/T[m]) {
				// accepted
			} else {
				tempPath[m][k], tempPath[m][l] = tempPath[m][l], tempPath[m][k] // revert
			}
			dists[m] = calcDistFromPath(tempPath[m])
		}

		for m := 0; m < nM-1; m++ {
			deltaS := (1/T[m+1] - 1/T[m]) * (dists[m] - dists[m+1])
			if deltaS < 0 || rand.Float64() < math.Exp(-deltaS) {
				tempPath[m], tempPath[m+1] = tempPath[m+1], tempPath[m]
			}
		}
	}

	// Write the final path to the file
	file, err := os.Create("/Users/kotaro/Desktop/MCMC_2d.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, cityIndex := range tempPath[nM-1] {
		_, err := fmt.Fprintf(file, "%f %f\n", cities[cityIndex][0], cities[cityIndex][1])
		if err != nil {
			panic(err)
		}
	}
}
