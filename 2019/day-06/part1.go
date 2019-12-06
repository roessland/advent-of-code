package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Object struct {
	Name string
	Orbits *Object
	Orbiters map[string]*Object
	Depth int
}

func GetOrCreateObject(objects map[string]*Object, name string) *Object {
	_, ok := objects[name]
	if !ok {
		objects[name] = &Object{Name: name, Orbits: nil, Orbiters: make(map[string]*Object)}
	}
	return objects[name]
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	objects := make(map[string]*Object)
	objects["COM"] = &Object{Name: "COM", Orbits: nil, Orbiters: make(map[string]*Object)}
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ")")
		orbiteeName := parts[0]
		orbiterName := parts[1]
		orbitee := GetOrCreateObject(objects, orbiteeName)
		orbiter := GetOrCreateObject(objects, orbiterName)
		orbitee.Orbiters[orbiterName] = orbiter
		orbiter.Orbits = orbitee
	}

	var updateDepth func(*Object)
	updateDepth = func(obj *Object) {
		for _, orbiter := range obj.Orbiters {
			orbiter.Depth = obj.Depth + 1
			updateDepth(orbiter)
		}
	}

	updateDepth(objects["COM"])

	var sumDepth func(*Object) int
	sumDepth = func(obj *Object) int {
		if obj == nil {
			return 0
		}
		c := obj.Depth
		for _, orbiter := range obj.Orbiters {
			c += sumDepth(orbiter)
		}
		return c
	}


	fmt.Println(sumDepth(objects["COM"]))
}
