package main

import "fmt"
import "io/ioutil"
import "log"

func ReadInput() []int {
	buf, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	nums := make([]int, len(buf)-1)
	for i, c := range buf {
		if c < '0' || c > '9' {
			break
		}
		nums[i] = int(c - '0')
	}
	return nums
}

func SolveCaptcha1(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[(i+1)%len(nums)] {
			sum += nums[i]
		}
	}
	return sum
}

func SolveCaptcha2(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[(i+len(nums)/2)%len(nums)] {
			sum += nums[i]
		}
	}
	return sum
}

func main() {
	nums := ReadInput()
	fmt.Println(SolveCaptcha1(nums))
	fmt.Println(SolveCaptcha2(nums))
}
