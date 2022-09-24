package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

const K = 5

var Data []Entry

func init() {
	readDataFromFile()
	normalizeData()
}

func normalizeData() {
	max1, min1 := float64(-1), float64(10000)
	max2, min2 := float64(-1), float64(10000)
	max3, min3 := float64(-1), float64(10000)
	max4, min4 := float64(-1), float64(10000)

	for i := 0; i < len(Data); i++ {
		min1 = math.Min(min1, Data[i].Attributes[0])
		max1 = math.Max(max1, Data[i].Attributes[0])
		min2 = math.Min(min2, Data[i].Attributes[1])
		max2 = math.Max(max2, Data[i].Attributes[1])
		min3 = math.Min(min3, Data[i].Attributes[2])
		max3 = math.Max(max3, Data[i].Attributes[2])
		min4 = math.Min(min4, Data[i].Attributes[3])
		max4 = math.Max(max4, Data[i].Attributes[3])
	}
	for i := 0; i < len(Data); i++ {
		Data[i].Attributes[0] = (Data[i].Attributes[0] - min1) / (max1 - min1)
		Data[i].Attributes[1] = (Data[i].Attributes[1] - min2) / (max2 - min2)
		Data[i].Attributes[2] = (Data[i].Attributes[2] - min3) / (max3 - min3)
		Data[i].Attributes[3] = (Data[i].Attributes[3] - min4) / (max4 - min4)
	}
}

func getDistanceBetweenAttributes(x1, y1, w1, z1, x2, y2, w2, z2 float64) float64 {
	return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2) + math.Pow(w1-w2, 2) + math.Pow(z1-z2, 2))
}

func getNearestNeighbours(e Entry, originalSet []Entry) []Entry {
	type ED struct {
		Entry
		distance float64
	}
	distanceArr := make([]ED, 0, len(originalSet))
	for i := 0; i < len(originalSet); i++ {
		curr := originalSet[i]
		distanceArr = append(distanceArr, ED{
			Entry: curr,
			distance: getDistanceBetweenAttributes(e.Attributes[0], e.Attributes[1], e.Attributes[2], e.Attributes[3],
				curr.Attributes[0], curr.Attributes[1], curr.Attributes[2], curr.Attributes[3]),
		})
	}
	sort.Slice(distanceArr, func(i, j int) bool {
		return distanceArr[i].distance < distanceArr[j].distance
	})

	res := make([]Entry, 0, K)
	for i := 0; i < K; i++ {
		res = append(res, distanceArr[i].Entry)
	}
	return res
}

func KNN(originalSet, classificationSet []Entry) {
	classifications := make([]Class, 0, len(classificationSet))
	for i := 0; i < len(classificationSet); i++ {
		nearestNeighbours := getNearestNeighbours(classificationSet[i], originalSet)
		setosaCount, versicolorCount, virginicaCount := 0, 0, 0
		for _, neighbour := range nearestNeighbours {
			switch neighbour.Class {
			case SETOSA_CLASS:
				setosaCount++
			case VERSICOLOR_CLASS:
				versicolorCount++
			case VIRGINICA_CLASS:
				virginicaCount++
			}
		}
		if setosaCount >= versicolorCount && setosaCount >= virginicaCount {
			classifications = append(classifications, SETOSA_CLASS)
		} else if versicolorCount >= setosaCount && versicolorCount >= virginicaCount {
			classifications = append(classifications, VERSICOLOR_CLASS)
		} else {
			classifications = append(classifications, VIRGINICA_CLASS)
		}
	}

	correctClassifications := 0
	for i := 0; i < len(classificationSet); i++ {
		if classifications[i] == classificationSet[i].Class {
			correctClassifications++
		}
	}
	fmt.Printf("Correct classifications were %d out of %d which is an accuracy of %f\n", correctClassifications, len(classificationSet), float64(correctClassifications)/float64(len(classificationSet)))
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(Data), func(i, j int) {
		Data[i], Data[j] = Data[j], Data[i]
	})
	KNN(Data[25:], Data[:25])
}
