package internal

import (
	"math"
	"slices"
)

type route struct {
	route    string
	distance int
}

/*
FindShortRoute implements a greedy algorithm to find a short route that visits
all cities starting from a specified city. The function does not guarantee the
shortest possible route (not optimal like the Traveling Salesman Problem) but
uses a greedy approach to choose the nearest unvisited city at each step.

Parameters:

	cities map[string]map[string]int: A map where keys represent city names and
	values are maps of neighboring cities with distances.
	startCity string: The name of the city from which the route starts.

Returns:

	route: A struct containing the route order as a string and the total
	distance as an integer.
*/
func FindShortRoute(cities map[string]map[string]int, startCity string) route {
	visitedCities := []string{startCity}
	nearbyCities := cities[startCity]
	totalDistance := 0
	routeOrder := startCity

	for len(visitedCities) < len(cities) {
		nextCity := ""
		shortestDistance := math.MaxInt

		for city, distance := range nearbyCities {
			if distance < shortestDistance && !slices.Contains(visitedCities, city) {
				nextCity = city
				shortestDistance = distance
			}
		}

		routeOrder += " -> " + nextCity
		nearbyCities = cities[nextCity]
		totalDistance += shortestDistance
		visitedCities = append(visitedCities, nextCity)
	}

	return route{routeOrder, totalDistance}
}
