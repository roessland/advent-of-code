package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"
import "math/rand"

func ReadDistances(filename string) map[string]map[string]int {
	dist := make(map[string]map[string]int)
	buf, _ := ioutil.ReadFile(filename)
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		words := strings.Split(line, " ")
		from := words[0]
		to := words[2]
		distance, _ := strconv.Atoi(words[4])

		if _, ok := dist[from]; !ok {
			dist[from] = make(map[string]int)
		}
		if _, ok := dist[to]; !ok {
			dist[to] = make(map[string]int)
		}
		dist[from][to] = distance
		dist[to][from] = distance
	}
	return dist
}

func GetCities(dist map[string]map[string]int) []string {
	cities := make([]string, 0)
	for city, _ := range dist {
		cities = append(cities, city)
	}
	return cities
}

func ShuffleStringSlice(s []string) {
	for i, _ := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
}

func TotalDistance(itinerary []string, dist map[string]map[string]int) int {
	total := 0
	for i := 0; i < len(itinerary)-1; i++ {
		total += dist[itinerary[i]][itinerary[i+1]]
	}
	return total
}

func main() {
	dist := ReadDistances("input.txt")
	itinerary := GetCities(dist)

	minDist := 999999999
	maxDist := 0
	for {
		d := TotalDistance(itinerary, dist)
		if d < minDist {
			minDist = d
			fmt.Printf("minDist: %v\n", minDist)
		}
		if d > maxDist {
			maxDist = d
			fmt.Printf("maxDist: %v\n", maxDist)
		}
		// bogo sort powah
		ShuffleStringSlice(itinerary)
	}
}
