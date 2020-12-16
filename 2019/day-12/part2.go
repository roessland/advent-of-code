package main

import "fmt"

func Sign(n int) int {
	if n > 0 {
		return 1
	}
	if n < 0 {
		return -1
	}
	return 0
}

type AxisState struct {
	P1, V1, P2, V2, P3, V3, P4, V4 int
}


func AxisUpdate(prevState AxisState) AxisState {
	state := prevState

	// Apply gravity
	state.V1 += Sign(prevState.P2 - prevState.P1)
	state.V1 += Sign(prevState.P3 - prevState.P1)
	state.V1 += Sign(prevState.P4 - prevState.P1)
	state.V2 += Sign(prevState.P1 - prevState.P2)
	state.V2 += Sign(prevState.P3 - prevState.P2)
	state.V2 += Sign(prevState.P4 - prevState.P2)
	state.V3 += Sign(prevState.P1 - prevState.P3)
	state.V3 += Sign(prevState.P2 - prevState.P3)
	state.V3 += Sign(prevState.P4 - prevState.P3)
	state.V4 += Sign(prevState.P1 - prevState.P4)
	state.V4 += Sign(prevState.P2 - prevState.P4)
	state.V4 += Sign(prevState.P3 - prevState.P4)


	// Apply velocity
	state.P1 += state.V1
	state.P2 += state.V2
	state.P3 += state.V3
	state.P4 += state.V4

	return state
}

func FindPeriod(axisState AxisState) int {
	lastVisited := map[AxisState]int{axisState: 0}
	for i := 1; ; i++ {
		axisState = AxisUpdate(axisState)
		if j, ok := lastVisited[axisState]; ok {
			return i-j
		}
		lastVisited[axisState] = i
	}
}

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	if a < b {
		return GCD(b, a)
	}
	return GCD(b, a%b)
}

func LCM(a, b int) int {
	return a*b/GCD(a,b)
}

func main() {
	x := FindPeriod(AxisState{P1: 14, P2: 12, P3: 1, P4: 16})
	y := FindPeriod(AxisState{P1: 4, P2: 10, P3: 7, P4: -5})
	z := FindPeriod(AxisState{P1: 5, P2: 8, P3: -10, P4: 3})

	// Planetary alignment algorithm
	fmt.Println(LCM(LCM(x, y), z))
}
