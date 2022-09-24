package main

import (
	"errors"
	"log"
	"strings"
	"utils"
)

type City struct {
	Name string
	utils.FloatCoord
}

func NewCity(name string, x, y float64) City {
	return City{
		Name: name,
		FloatCoord: utils.FloatCoord{
			X: x,
			Y: y,
		},
	}
}

type Individual struct {
	route    []City
	distance float64
	Fitness  float64
}

func NewIndividual(route []City) *Individual {
	var (
		i   = &Individual{route: route}
		err error
	)

	i.distance, err = i.calculateDistance()
	if err != nil {
		log.Fatalf("while creating individual: %v\n", err)
	}
	i.Fitness = 1 / i.distance
	return i
}

func (in *Individual) String() string {
	var (
		b        = strings.Builder{}
		routeLen = len(in.route)
	)
	b.WriteString("[")
	for i := 0; i < routeLen; i++ {
		b.WriteString(in.route[i].Name)
		if i < routeLen-1 {
			b.WriteString(" , ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func (in *Individual) calculateDistance() (float64, error) {
	routeLen := len(in.route)
	if routeLen <= 1 {
		return 0, errors.New("cannot calc distance of a route with only one city")
	}
	dist := float64(0)
	for j := 0; j < routeLen; j++ {
		if j < routeLen-1 {
			dist += in.route[j].Distance(in.route[j+1].FloatCoord)
		} else {
			//TODO: Enable if we want to actually return to the first city
			//dist += in.route[j].Distance(in.route[0].FloatCoord)
		}
	}
	return dist, nil
}
