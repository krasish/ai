package main

import "fmt"

const dataPath = "/Users/i537925/SAPDevelop/temp/ai/task7/data/house-votes-84.data"

var Data []Entry

func init() {
	readDataFromFile()
}

func NaiveBayesClassifier(trainingSet, classificationSet []Entry) {
	const yesIndex, noIndex = 1, 0
	var (
		tsLen, csLen = len(trainingSet), len(classificationSet)

		demoCount, repCount                     int
		demoAttributeCounts, repAttributeCounts = new([2][AttributesCount]int), new([2][AttributesCount]int)

		demoProb, repProb                     float64
		demoAttributeProbs, repAttributeProbs = new([2][AttributesCount]float64), new([2][AttributesCount]float64)

		classifications        = make([]Class, len(classificationSet))
		correctClassifications = 0
	)

	//Learning phase

	for i := 0; i < tsLen; i++ {
		if trainingSet[i].Class == DEMOCRAT_CLASS {
			demoCount++
			for j := 0; j < AttributesCount; j++ {
				if trainingSet[i].Attributes[j] == YES_ATTRIBUTE_VALUE {
					demoAttributeCounts[yesIndex][j]++
				} else if trainingSet[i].Attributes[j] == NO_ATTRIBUTE_VALUE {
					demoAttributeCounts[noIndex][j]++
				}
			}
		} else {
			repCount++
			for j := 0; j < AttributesCount; j++ {
				if trainingSet[i].Attributes[j] == YES_ATTRIBUTE_VALUE {
					repAttributeCounts[yesIndex][j]++
				} else if trainingSet[i].Attributes[j] == NO_ATTRIBUTE_VALUE {
					repAttributeCounts[noIndex][j]++
				}
			}
		}
	}

	demoProb, repProb = float64(demoCount)/float64(tsLen), float64(repCount)/float64(tsLen)
	for i := 0; i < AttributesCount; i++ {
		demoAttributeProbs[yesIndex][i] = float64(demoAttributeCounts[yesIndex][i]) / float64(demoCount)
		demoAttributeProbs[noIndex][i] = float64(demoAttributeCounts[noIndex][i]) / float64(demoCount)

		repAttributeProbs[yesIndex][i] = float64(repAttributeCounts[yesIndex][i]) / float64(repCount)
		repAttributeProbs[noIndex][i] = float64(repAttributeCounts[noIndex][i]) / float64(repCount)
	}

	//Classification phase

	for i := 0; i < csLen; i++ {
		currentDemoProb, currentRepProb := demoProb, repProb
		for j := 0; j < AttributesCount; j++ {
			if classificationSet[i].Attributes[j] == YES_ATTRIBUTE_VALUE {
				currentDemoProb *= demoAttributeProbs[yesIndex][j]
				currentRepProb *= repAttributeProbs[yesIndex][j]
			} else if classificationSet[i].Attributes[j] == NO_ATTRIBUTE_VALUE {
				currentDemoProb *= demoAttributeProbs[noIndex][j]
				currentRepProb *= repAttributeProbs[noIndex][j]
			}
		}
		if currentDemoProb > currentRepProb {
			classifications[i] = DEMOCRAT_CLASS
		} else {
			classifications[i] = REPUBLICAN_CLASS
		}
	}

	for i := 0; i < csLen; i++ {
		if classifications[i] == classificationSet[i].Class {
			correctClassifications++
		}
	}
	fmt.Printf("Correct classifications were %d out of %d which is an accuracy of %f\n", correctClassifications, csLen, float64(correctClassifications)/float64(csLen))
}

func main() {
	fmt.Println()
	for i := 0; i < 10; i++ {
		fmt.Printf("=== RUN %d ===\n", i)
		var (
			segmentSize                    = len(Data) / 10
			trainingSet, classificationSet []Entry
		)
		if i == 0 {
			classificationSet, trainingSet = Data[0:segmentSize], Data[segmentSize:]
		} else if i == 9 {
			classificationSet, trainingSet = Data[len(Data)-segmentSize:], Data[0:len(Data)-segmentSize]
		} else {
			classificationSet, trainingSet = Data[i*segmentSize:i*segmentSize+segmentSize], append(Data[0:i*segmentSize], Data[i*segmentSize+segmentSize:]...)
		}
		NaiveBayesClassifier(trainingSet, classificationSet)
		fmt.Printf("=== END OF RUN %d ===\n\n", i)
	}
}
