package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Class string

const (
	SETOSA_CLASS     Class = "Iris-setosa"
	VERSICOLOR_CLASS Class = "Iris-versicolor"
	VIRGINICA_CLASS  Class = "Iris-virginica"
)

const (
	AttributesCount = 4
	dataPath        = "/Users/i537925/SAPDevelop/temp/ai/task8/data/iris.data"
)

type Entry struct {
	Class
	Attributes [AttributesCount]float64
}

func (e Entry) String() string {
	var (
		b = strings.Builder{}
	)
	b.WriteString("[")
	b.WriteString(string(e.Class))
	b.WriteString(", ")
	for i := 0; i < AttributesCount; i++ {
		b.WriteString(fmt.Sprintf("%f", e.Attributes[i]))
		if i < AttributesCount-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString("]\n")
	return b.String()
}

func readDataFromFile() {
	dataF, err := os.Open(dataPath)
	if err != nil {
		log.Fatalf("while opening dataCSV file: %v", err)
	}

	dataReader := csv.NewReader(dataF)

	dataCSV, err := dataReader.ReadAll()
	if err != nil {
		log.Fatalf("while reading dataCSV file: %v", err)
	}

	Data = make([]Entry, len(dataCSV))

	for i := 0; i < len(dataCSV); i++ {
		for j := 0; j < AttributesCount; j++ {
			currAttrF, err := strconv.ParseFloat(dataCSV[i][j], 64)
			if err != nil {
				log.Fatalf("while reading data: %v", err)
			}
			Data[i].Attributes[j] = currAttrF
		}
		Data[i].Class = Class(dataCSV[i][AttributesCount])
	}
}
