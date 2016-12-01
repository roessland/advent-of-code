package main

import "fmt"

type Reindeer struct {
	flyingSpeed      int
	flyingTime       int
	flyingTimeLeft   int
	restingTime      int
	restingTimeLeft  int
	distanceTraveled int
	points           int
	state            string
}

func (r *Reindeer) Update() {
	if r.flyingTimeLeft > 0 {
		r.flyingTimeLeft--
		r.distanceTraveled += r.flyingSpeed
		if r.flyingTimeLeft == 0 {
			r.restingTimeLeft = r.restingTime
			r.state = "resting"
		}
	} else if r.restingTimeLeft > 0 {
		r.restingTimeLeft--
		if r.restingTimeLeft == 0 {
			r.flyingTimeLeft = r.flyingTime
			r.state = "flying"
		}
	}
}

func NewReindeer(flyingSpeed, flyingTime, restingTime int) Reindeer {
	return Reindeer{
		flyingSpeed:    flyingSpeed,
		flyingTime:     flyingTime,
		flyingTimeLeft: flyingTime,
		restingTime:    restingTime,
		state:          "flying",
	}
}

func LeadingReindeers(reindeers []Reindeer) []*Reindeer {
	maxReindeers := []*Reindeer{}
	maxDistance := -1
	for i, _ := range reindeers {
		if reindeers[i].distanceTraveled > maxDistance {
			maxDistance = reindeers[i].distanceTraveled
			maxReindeers = []*Reindeer{&reindeers[i]}
		} else if reindeers[i].distanceTraveled == maxDistance {
			maxReindeers = append(maxReindeers, &reindeers[i])
		}
	}
	return maxReindeers
}

func main() {
	reindeers := []Reindeer{
		NewReindeer(19, 7, 124),
		NewReindeer(3, 15, 28),
		NewReindeer(19, 9, 164),
		NewReindeer(19, 9, 158),
		NewReindeer(13, 7, 82),
		NewReindeer(25, 6, 145),
		NewReindeer(14, 3, 38),
		NewReindeer(3, 16, 37),
		NewReindeer(25, 6, 143),
	}

	for i := 0; i < 2503; i++ {
		for i, _ := range reindeers {
			reindeers[i].Update()
		}
		for _, reindeer := range LeadingReindeers(reindeers) {
			reindeer.points++
		}
	}
	maxPoints := -1
	maxDistance := -1
	for _, reindeer := range reindeers {
		if reindeer.points > maxPoints {
			maxPoints = reindeer.points
		}
		if reindeer.distanceTraveled > maxDistance {
			maxDistance = reindeer.distanceTraveled
		}
	}
	fmt.Printf("Winner has %v points.\n", maxPoints)
	fmt.Printf("Leader has travelled %v km.\n", maxDistance)

}
