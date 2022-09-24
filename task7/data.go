package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

type AttributeValue byte

const (
	YES_ATTRIBUTE_VALUE     AttributeValue = 'y'
	NO_ATTRIBUTE_VALUE      AttributeValue = 'n'
	UNKNOWN_ATTRIBUTE_VALUE AttributeValue = '?'
)

func NewAttributeValue(stringValue string) AttributeValue {
	switch stringValue {
	case "y":
		return YES_ATTRIBUTE_VALUE
	case "n":
		return NO_ATTRIBUTE_VALUE
	default:
		return UNKNOWN_ATTRIBUTE_VALUE
	}
}

type Class string

const (
	DEMOCRAT_CLASS   Class = "democrat"
	REPUBLICAN_CLASS Class = "republican"
)

const AttributesCount = 16

type Entry struct {
	Class
	Attributes [AttributesCount]AttributeValue
}

func (e Entry) String() string {
	var (
		b = strings.Builder{}
	)
	b.WriteString("[")
	b.WriteString(string(e.Class))
	b.WriteString(", ")
	for i := 0; i < AttributesCount; i++ {
		b.WriteString(string(e.Attributes[i]))
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
		Data[i].Class = Class(dataCSV[i][0])
		for j := 1; j <= AttributesCount; j++ {
			Data[i].Attributes[j-1] = NewAttributeValue(dataCSV[i][j])
		}
	}
}
