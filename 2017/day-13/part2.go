package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func main() {
	depths := []int{}
	sizes := []int{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		fields := strings.Split(scanner.Text(), ": ")
		depth, _ := strconv.Atoi(fields[0])
		size, _ := strconv.Atoi(fields[1])
		depths = append(depths, depth)
		sizes = append(sizes, size)
	}

	for delay := 0; ; delay++ {
		caught := false
		for i := 0; i < len(depths); i++ {
			if (depths[i]+delay)%(2*sizes[i]-2) == 0 {
				caught = true
				break
			}
		}
		if !caught {
			fmt.Println(delay)
			break
		}
	}
}
