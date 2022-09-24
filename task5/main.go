package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"

	exp_rand "golang.org/x/exp/rand"

	"gonum.org/v1/gonum/stat/distuv"
)

var (
	populationSize = 150
	eliteSize      = 35
	generations    = 2000
	mutationRate   = 0.20

	cityNamesFilePath       = `/Users/i537925/SAPDevelop/temp/ai/task5/UK_TSP/uk12_name.csv`
	cityCoordinatesFilePath = `/Users/i537925/SAPDevelop/temp/ai/task5/UK_TSP/uk12_xy.csv`

	cities     = make([]City, 0, 12)
	population []*Individual
)

func readCitiesFromFile() {
	namesF, err := os.Open(cityNamesFilePath)
	if err != nil {
		log.Fatalf("while reading names file: %v", err)
	}
	coordsF, err := os.Open(cityCoordinatesFilePath)
	if err != nil {
		log.Fatalf("while reading names file: %v", err)
	}
	namesReader := csv.NewReader(namesF)
	coordsReader := csv.NewReader(coordsF)

	names, err := namesReader.ReadAll()
	if err != nil {
		log.Fatalf("while reading names file: %v", err)
	}

	coords, err := coordsReader.ReadAll()
	if err != nil {
		log.Fatalf("while reading coords file: %v", err)
	}

	for i := 0; i < len(names); i++ {

		floatX, err := strconv.ParseFloat(coords[i][0], 64)
		if err != nil {
			log.Fatalf("could not read floatX: %v", err)
		}
		floatY, err := strconv.ParseFloat(coords[i][1], 64)
		if err != nil {
			log.Fatalf("could not read floatY: %v", err)
		}
		cities = append(cities, NewCity(names[i][0], floatX, floatY))
	}
}

func init() {
	readCitiesFromFile()
}

func GeneticAlgorithm() {
	generateInitialPopulation(populationSize)

	for i := 0; i < generations; i++ {
		population = nextGeneration()
	}
	sortPopulation()
}

func generateInitialPopulation(size int) {
	citiesLen := len(cities)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		route := make([]City, citiesLen)
		copy(route, cities)
		rand.Shuffle(citiesLen, func(i, j int) { route[i], route[j] = route[j], route[i] })
		population = append(population, NewIndividual(route))
	}
}

func nextGeneration() []*Individual {
	sortPopulation()
	fmt.Println("The current best is: ", population[0])
	result := getMatingPool()
	result = breedGeneration(result)
	result = mutateGeneration(result)
	return result
}

func sortPopulation() {
	sort.SliceStable(population, func(i, j int) bool {
		return population[i].Fitness > population[j].Fitness
	})
}

func getMatingPool() []*Individual {
	var (
		weights = make([]float64, 0, populationSize)
		ins     = make([]*Individual, 0, populationSize)
		i       = 0
	)
	for i := 0; i < populationSize; i++ {
		weights = append(weights, population[i].Fitness)
	}
	source := exp_rand.NewSource(uint64(time.Now().UnixNano()))
	weightsDist := distuv.NewCategorical(weights, source)

	for ; i < eliteSize; i++ {
		ins = append(ins, population[i])
	}
	for ; i < populationSize; i++ {
		ins = append(ins, population[int(weightsDist.Rand())])
	}
	return ins
}

func breedGeneration(matingPool []*Individual) []*Individual {
	var (
		children = make([]*Individual, 0, populationSize)
		i        = 0
	)
	for ; i < eliteSize; i++ {
		children = append(children, matingPool[i])
	}

	rand.Seed(time.Now().UnixNano())
	for ; i < populationSize; i++ {
		children = append(children, breed(matingPool[rand.Intn(populationSize)], matingPool[rand.Intn(populationSize)]))
	}

	return children
}

func breed(p1 *Individual, p2 *Individual) *Individual {
	var (
		geneLength = len(p1.route)
		childRoute = make([]City, geneLength)
	)

	firstGene, secondGene := rand.Intn(geneLength), rand.Intn(geneLength)
	startGene, endGene := int(math.Min(float64(firstGene), float64(secondGene))), int(math.Max(float64(firstGene), float64(secondGene)))

	if startGene == endGene {
		return NewIndividual(p2.route)
	}

	copy(childRoute[startGene:endGene], p1.route[startGene:endGene])
	alreadyPresent := make(map[string]struct{}, geneLength)
	for i := startGene; i < endGene; i++ {
		alreadyPresent[childRoute[i].Name] = struct{}{}
	}
	for i := endGene; i%geneLength != startGene; i++ {
		for j := endGene; ; j++ {
			current := p2.route[j%geneLength]
			if _, ok := alreadyPresent[current.Name]; !ok {
				childRoute[i%geneLength] = current
				alreadyPresent[current.Name] = struct{}{}
				break
			}
		}
	}
	return NewIndividual(childRoute)
}

func mutateGeneration(currentGeneration []*Individual) []*Individual {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(currentGeneration); i++ {
		if rand.Float64() < mutationRate {
			var (
				routeLen   = len(currentGeneration[i].route)
				firstGene  = rand.Intn(routeLen)
				secondGene = rand.Intn(routeLen)
			)
			if firstGene == secondGene {
				continue
			}
			currentGeneration[i].route[firstGene], currentGeneration[i].route[secondGene] = currentGeneration[i].route[secondGene], currentGeneration[i].route[firstGene]
		}
	}
	return currentGeneration
}

func main() {
	GeneticAlgorithm()
	//fmt.Println()
	//fmt.Println("TOP 10: ")
	//fmt.Println()
	//for i := 0; i < 10; i++ {
	//	fmt.Printf("%v - %f \n", population[i], population[i].Fitness)
	//}
	fmt.Println()
	fmt.Println("BEST: ")
	best := NewIndividual([]City{cities[0], cities[5], cities[9], cities[4], cities[2], cities[7], cities[11], cities[3], cities[6], cities[10], cities[1], cities[8]})
	fmt.Println()
	fmt.Printf("%v - %f ; %f \n", best, best.Fitness, best.distance)

	//fmt.Println()
	//fmt.Println("MY BEST: ")
	//fmt.Println()
	//fmt.Printf("%v - %f ; %f \n", population[0], population[0].Fitness, population[0].distance)

	//source := exp_rand.NewSource(uint64(time.Now().UnixNano()))
	//catDist := distuv.NewCategorical([]float64{0.00350, 0.00300, 0.00200, 0.00150, 0.08}, source)
	//for i := 0; i < 200; i++ {
	//	fmt.Println(catDist.Rand())
	//}
}
