package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

func simulateWire(in1, in2, op, out string, chs map[string]chan uint16, ready, done chan bool) {
	<-ready

	var value1, value2, value uint16

	// Get in1
	val, err := strconv.Atoi(in1)
	if err != nil {
		value1 = <-chs[in1]
	} else {
		value1 = uint16(val)
	}

	// Get in2
	if len(in2) > 0 {
		val, err := strconv.Atoi(in2)
		if err != nil {
			value2 = <-chs[in2]
		} else {
			value2 = uint16(val)
		}
	}

	if op == "" {
		value = value1
	} else if op == "NOT" {
		value = ^value1
	} else if op == "OR" {
		value = value1 | value2
	} else if op == "AND" {
		value = value1 & value2
	} else if op == "LSHIFT" {
		value = value1 << value2
	} else if op == "RSHIFT" {
		value = value1 >> value2
	}

	if out == "b" {
		value = uint16(16076)
	}

	if out == "a" {
		fmt.Printf("a signal: %v\n", value)
		done <- true
	}

	// Keep sending result
	for {
		select {
		case chs[out] <- value:
			continue
		case <-ready:
			return
		}
	}
}

func main() {
	buf, _ := ioutil.ReadFile("input.txt")
	chs := make(map[string]chan uint16)
	_ = chs
	ready := make(chan bool)
	done := make(chan bool)
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		var in1, in2, op, out string
		parts := strings.Split(line, " -> ")
		out = parts[1]
		part0 := strings.Split(parts[0], " ")
		if len(part0) == 1 {
			in1 = part0[0]
		} else if len(part0) == 2 {
			in1 = part0[1]
			op = part0[0]
		} else if len(part0) == 3 {
			in1 = part0[0]
			op = part0[1]
			in2 = part0[2]
		}
		chs[out] = make(chan uint16, 2)
		fmt.Printf("print to prevent deadlock??!\n")
		go simulateWire(in1, in2, op, out, chs, ready, done)
	}
	for _, _ = range chs {
		ready <- true
	}
	fmt.Printf("ready sent to all\n")
	<-done
	for _, _ = range chs {
		ready <- true
	}
	fmt.Printf("DONE")
}
