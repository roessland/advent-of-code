package main

import "fmt"
import "log"
import "encoding/csv"
import "os"

func IsValid(words []string) bool {
	s := make(map[string]bool)
	for _, w := range words {
		s[w] = true
	}
	return len(s) == len(words)
}

func main() {
	reader := csv.NewReader(os.Stdin)
	reader.Comma = ' '
	reader.FieldsPerRecord = -1
	phrases, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	valid := 0
	for _, phrase := range phrases {
		if IsValid(phrase) {
			valid++
		}
	}
	fmt.Println(valid)
}
