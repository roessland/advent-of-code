package main

import "fmt"
import "regexp"
import "io/ioutil"
import "strconv"

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	s := string(buf)
	re := regexp.MustCompile(`-{0,1}[0-9]+`)
	sum := 0
	for _, numStr := range re.FindAllString(s, -1) {
		num, _ := strconv.Atoi(numStr)
		sum += num
	}
	fmt.Printf("%v\n", sum)
}
