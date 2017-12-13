package main

import "bufio"
import "fmt"
import "os"
import "strconv"
import "strings"

func main() {
	severity := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			break
		}
		fields := strings.Split(scanner.Text(), ": ")
		depth, _ := strconv.Atoi(fields[0])
		size, _ := strconv.Atoi(fields[1])
		if depth%(2*size-2) == 0 {
			severity += depth * size
		}
	}
	fmt.Println(severity)
}
