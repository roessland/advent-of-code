package main

import "fmt"


func Transform(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value %= 20201227
	}
	return value
}

func InvTransform(subjectNumber, targetValue int) int {
	value := 1
	for i := 0; ; i++ {
		if value == targetValue {
			return i
		}
		value = (value*subjectNumber)%20201227
	}
}

func main() {
	cardPub := 10705932
	doorPub := 12301431
	cardPriv := InvTransform(7, cardPub)
	fmt.Println("Part 1:", Transform(doorPub, cardPriv))
}