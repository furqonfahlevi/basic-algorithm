package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 10
const M = 10

var rate1, rate2, rate3, rate4, rate5 float64

// rate 1 = infection rate in a cell
// rate 2 = infection rate to neighbouring cell
// rate 3 = detection rate
// rate 4 = recovering people
// rate 5 = dying people

type cellData struct {
	population int // all living people (inluding undetected infected person)
	covid int // undetected infected
	hospitalized int // hospitalized
	// dead people = initial p - final p - final h
}

type Cell [N][M]cellData

func randPopulation(data *Cell) {
	// randomize the number of population of each cell
	rand.Seed(time.Now().UnixNano()) // will maximaze the randomness
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			data[i][j].population = rand.Intn(50) + 1 // 100 can be changed
		}
	}
}

func countPopulation(data Cell) int {
	var totPop int // total population

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			totPop = totPop + data[i][j].population
		}
	}

	return totPop
}

func countInfected(data Cell) int {
	var totInfect int // total population

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			totInfect = totInfect + data[i][j].covid
		}
	}

	return totInfect
}

func countHospitalized(data Cell) int {
	var totHos int // total hospitalized

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			totHos = totHos + data[i][j].hospitalized
		}
	}

	return totHos
}

func updateCell(data *Cell, totInfected int) {
	rate1 = float64(countInfected(*data) / totInfected)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			data[i][j].covid = int(float64(data[i][j].covid) * rate1) // updates infected people
		}
	}
}

func updateNeighbour(data *Cell, totInfected int) {
	rate2 = float64(countPopulation(*data) / totInfected)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var top bool = (i-1 >= 0)   // true if there is a cell on top
			var right bool = (j+1 < M)  // true if there is a cell to the right
			var left bool = (j-1 >= 0)  // true if there is a cell to the left
			var bottom bool = (i+1 < N) // true if there is a cell in botttom

			if top {
				// for top
				data[i-1][j].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i-1][j].population) + float64(data[i-1][j].covid))
				if data[i-1][j].covid > data[i-1][j].population {
					data[i-1][j].covid = data[i-1][j].population
				}
				if right {
					// for top-right
					data[i-1][j+1].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i-1][j+1].population) + float64(data[i-1][j+1].covid))
					if data[i-1][j+1].covid > data[i-1][j+1].population {
						data[i-1][j+1].covid = data[i-1][j+1].population
					}
				}
				if left {
					// for top-left
					data[i-1][j-1].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i-1][j-1].population) + float64(data[i-1][j-1].covid))
					if data[i-1][j-1].covid > data[i-1][j-1].population {
						data[i-1][j-1].covid = data[i-1][j-1].population
					}
				}
			}

			if bottom {
				// for bottom
				data[i+1][j].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i+1][j].population) + float64(data[i+1][j].covid))
				if data[i+1][j].covid > data[i+1][j].population {
					data[i+1][j].covid = data[i+1][j].population
				}
				if right {
					// for bottom-right
					data[i+1][j+1].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i+1][j+1].population) + float64(data[i+1][j+1].covid))
					if data[i+1][j+1].covid > data[i+1][j+1].population {
						data[i+1][j+1].covid = data[i+1][j+1].population
					}
				}
				if left {
					// for bottom-left
					data[i+1][j-1].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i+1][j-1].population) + float64(data[i+1][j-1].covid))
					if data[i+1][j-1].covid > data[i+1][j-1].population {
						data[i+1][j-1].covid = data[i+1][j-1].population
					}
				}
			}

			if right {
				// for right
				data[i][j+1].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i][j+1].population) + float64(data[i][j+1].covid))
				if data[i][j+1].covid > data[i][j+1].population {
					data[i][j+1].covid = data[i][j+1].population
				}
			}

			if left {
				// for left
				data[i][j-1].covid = int(float64(data[i][j].covid)*(rate2/100)*float64(data[i][j-1].population) + float64(data[i][j-1].covid))
				if data[i][j-1].covid > data[i][j-1].population {
					data[i][j-1].covid = data[i][j-1].population
				}
			}
		}
	}
}

func updateHospitalized(data *Cell, totInfected int) {
	rate3 = float64(countPopulation(*data) / totInfected)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var initH int = data[i][j].hospitalized
			var newH int = int(float64(data[i][j].covid)*(rate3/100) + float64(data[i][j].hospitalized))
			data[i][j].hospitalized = newH
			data[i][j].population = data[i][j].population - newH - initH
			if data[i][j].population < 0 {
				data[i][j].population = 0
			}
		}
	}
}

func updateRecovered(data *Cell) {
	if countHospitalized(*data) > 0 {
		rate4 = float64(countPopulation(*data) / countHospitalized(*data))
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var recovered int = int(float64(data[i][j].hospitalized) * (rate4 / 100))
			data[i][j].hospitalized = data[i][j].hospitalized - recovered
			data[i][j].population = data[i][j].population + recovered
			if data[i][j].hospitalized < 0 {
				data[i][j].hospitalized = 0
			}
		}
	}
}

func updateDied(data *Cell) {
	if countHospitalized(*data) > 0 {
		rate5 = float64(countPopulation(*data) / countHospitalized(*data))
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var died int = int(float64(data[i][j].hospitalized) * (rate5 / 100))
			data[i][j].hospitalized = data[i][j].hospitalized - died
			if data[i][j].hospitalized < 0 {
				data[i][j].hospitalized = 0
			}
		}
	}
}

func update(data *Cell, totInfected int) {
	updateCell(data, totInfected)         // updates infected people
	updateNeighbour(data, totInfected)    // updates neighboring infected people
	updateHospitalized(data, totInfected) // updates hospitalized people
	updateRecovered(data)                 // updates recovered people
	updateDied(data)                      // updates dying people
}

func main() {
	var complex Cell
	var t int // time limit
	var time int = 0
	var i, j int                // inital cell that has will be inputed by user
	var p, c int                // initial living and undetected infected people in cell(i,j)
	var hiInfect int            // highest increase of infection
	var startOfReduc int        // start of infected reduction
	var totInfected int = 0     // total infected since time = 0 to time = t
	var totHospitalized int = 0 // total infected since time = 0 to time = t

	randPopulation(&complex) // randomize the population
	fmt.Printf("Matrix Size: %v x %v\n", N, M)
	fmt.Printf("Input initial cell location (i,j): ")
	fmt.Scanln(&i, &j)
	for i > N || j > M {
		// to make sure i and j is within N and M
		fmt.Printf("Out of Bounds \n")
		fmt.Printf("Input initial cell location (i,j): ")
		fmt.Scanln(&i, &j)
	}

	fmt.Printf("Input living people (p) and undetected infected people (c): ")
	fmt.Scanln(&p, &c)
	for c > p {
		fmt.Printf("c is larger than p \n")
		fmt.Printf("Input living people (p) and undetected infected people (c): ")
		fmt.Scanln(&p, &c)
	}

	complex[i-1][j-1].population = p
	complex[i-1][j-1].covid = c
	var initialP int = countPopulation(complex) // count the initial population

	fmt.Printf("Input time limit: ")
	fmt.Scanln(&t)

	var infected, befInfected int = 0, 0
	var hospitalized, befHospitalized int = 0, 0
	var highestInfect int = 0
	for time < t {
		infected = countInfected(complex)
		hospitalized = countHospitalized(complex)
		if (infected - befInfected) > 0 {
			totInfected = totInfected + (infected - befInfected)
		}
		if (hospitalized - befHospitalized) > 0 {
			totHospitalized = totHospitalized + (hospitalized - befHospitalized)
		}
		if infected-befInfected > highestInfect {
			highestInfect = infected - befInfected
			hiInfect = time
		}
		if befInfected > infected && startOfReduc == 0 {
			startOfReduc = time
		}
		befInfected = infected
		befHospitalized = hospitalized
		update(&complex, totInfected)
		time++
	}
	var finalP = countPopulation(complex)
	var peopleDied int = initialP - finalP
	if countHospitalized(complex) > 1 {
		peopleDied = peopleDied - countHospitalized(complex)
		if peopleDied < 0 {
			peopleDied = peopleDied * -1
		}
	}

	if totInfected < 0 {
		fmt.Println("No on is longer infected")
	}

	fmt.Println()
	fmt.Printf("Original total size of population = %v \n", initialP)
	fmt.Printf("The total people hospitalized during simulation = %v \n", totHospitalized)
	fmt.Printf("The total people died during simulation = %v \n", peopleDied)
	fmt.Printf("The highest increase of infection during simulation is at t = %v with %v people \n", hiInfect, highestInfect)
	fmt.Printf("The start of reduction of the number of infected people is at t = %v \n", startOfReduc)
}
