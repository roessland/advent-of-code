package main

import "fmt"
import "io/ioutil"
import "encoding/json"

func sumJsonData(data interface{}) int {
	sum := 0
	switch dat := data.(type) {
	case float64:
		sum += int(data.(float64))
	case map[string]interface{}:
		// Ignore maps with a value "red"
		for _, subdata := range dat {
			if stringVal, ok := subdata.(string); ok {
				if stringVal == "red" {
					return 0
				}
			}
		}
		for _, subdata := range dat {
			sum += sumJsonData(subdata)
		}
	case []interface{}:
		for _, subdata := range dat {
			sum += sumJsonData(subdata)
		}
	case string:
	default:
		fmt.Printf("Type not understood %v, %T\n", data, data)
	}
	return sum
}

func main() {
	jsonBlob, _ := ioutil.ReadFile("input.txt")

	var jsonData interface{}

	if err := json.Unmarshal(jsonBlob, &jsonData); err != nil {
		fmt.Errorf("%v\n", err)
	}
	fmt.Printf("WOAH %v\n", sumJsonData(jsonData))
}
