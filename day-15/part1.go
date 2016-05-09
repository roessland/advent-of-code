package main

import "fmt"

type Ingredient []int
type Recipe []int

func Score(ings []Ingredient, rec Recipe) int {
	if rec[0]+rec[1]+rec[2]+rec[3] != 100 {
		fmt.Printf("%v\n", rec)
		panic("Invalid recipe")
	}
	props := make([]int, 5)
	for i, _ := range rec {
		props[0] += rec[i] * ings[i][0]
		props[1] += rec[i] * ings[i][1]
		props[2] += rec[i] * ings[i][2]
		props[3] += rec[i] * ings[i][3]
		props[4] += rec[i] * ings[i][4]
	}
	if props[0] < 1 || props[1] < 1 || props[2] < 1 || props[3] < 1 {
		return 0
	}
	if props[4] != 500 {
		return 0
	}
	return props[0] * props[1] * props[2] * props[3]
}

func main() {
	ings := []Ingredient{
		Ingredient{2, 0, -2, 0, 3},
		Ingredient{0, 5, -3, 0, 3},
		Ingredient{0, 0, 5, -1, 8},
		Ingredient{0, -1, 0, 5, 8},
	}

	maxScore := 0

	for c1 := 1; c1 <= 100; c1++ {
		for c2 := 1; c1+c2 <= 100; c2++ {
			for c4 := 1; 3*c4 < 2*c1+3*c2; c4++ {
				for c3 := 1; c3 < 5*c4; c3++ {
					if c1+c2+c3+c4 != 100 {
						continue
					}
					rec := Recipe{c1, c2, c3, c4}
					s := Score(ings, rec)
					if s > maxScore {
						maxScore = s
					}
				}
			}
		}
	}
	fmt.Printf("yeah %v\n", maxScore)
}
