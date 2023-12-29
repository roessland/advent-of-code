package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
)

func debugSend(src, dst string, pulse Pulse) {
	return
	if pulse == Lo {
		fmt.Printf("%s -low-> %s\n", src, dst)
	} else {
		fmt.Printf("%s -high-> %s\n", src, dst)
	}
}

func debugReceive(src, dst string, pulse Pulse) {
	return
	if pulse == Lo {
		fmt.Printf("%s <-low- %s\n", src, dst)
	} else {
		fmt.Printf("%s <-high- %s\n", src, dst)
	}
}

func main() {
	input := ReadInput()

	modules := map[string]Module{}
	for moduleName, inputModule := range input.Modules {
		bm := inputModule
		bm.Modules = modules
		switch inputModule.Type {
		case "broadcaster":
			modules[moduleName] = &BroadcastModule{BaseModule: bm}
		case "button":
			modules[moduleName] = &ButtonModule{BaseModule: bm}
		case "conjunction":
			modules[moduleName] = &ConjunctionModule{BaseModule: bm}
		case "flipflop":
			modules[moduleName] = &FlipFlopModule{BaseModule: bm}
		case "output":
			modules[moduleName] = &OutputModule{BaseModule: bm}

		default:
			panic(fmt.Sprint("unknown module type", inputModule.Type))
		}
		modules[moduleName].Init()
	}

	numLo, numHi := 0, 0
	for i := 0; i < 1000; i++ {
		pq = append(pq, QueueItem{
			Process: func() (int, int) { return modules["button"].Process(Lo, "main") },
			PulseID: GetPulseID(),
		})
		for len(pq) > 0 {
			qi := pq[0]
			pq = pq[1:]
			l, h := qi.Process()
			numLo += l
			numHi += h
			FixPQ()
		}
	}

	fmt.Println("numHi:", numHi)
	fmt.Println("numLo:", numLo)
	fmt.Println("Part 1:", numLo*numHi)
}

type State struct {
	Modules map[string]BaseModule
}

type Pulse byte

const (
	Lo Pulse = 1
	Hi Pulse = 2
)

type Module interface {
	Init()
	Process(pulse Pulse, from string) (numLo, numHi int)
	Dsts() []string
}

type BaseModule struct {
	Modules map[string]Module
	Name    string
	Inputs  map[string]struct{}
	Type    string
	dsts    []string
}

func (m *BaseModule) Dsts() []string {
	return m.dsts
}

type ButtonModule struct {
	BaseModule
}

func (m *ButtonModule) Init() {
}

func (m *ButtonModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	return Emit(m.BaseModule, pulse)
}

var _ Module = &ButtonModule{}

type BroadcastModule struct {
	BaseModule
}

// Init implements Module.
func (*BroadcastModule) Init() {
}

var _ Module = &BroadcastModule{}

func (m *BroadcastModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	debugReceive(m.Name, "flipflop", pulse)
	return Emit(m.BaseModule, pulse)
}

type QueueItem struct {
	Process func() (int, int)
	PulseID int
}

var nextPulseID int

func GetPulseID() int {
	nextPulseID++
	return nextPulseID
}

var pq []QueueItem

func FixPQ() {
	slices.SortStableFunc(pq, func(a, b QueueItem) int {
		return cmp.Compare(a.PulseID, b.PulseID)
	})
}

func Emit(m BaseModule, pulse Pulse) (numLo, numHi int) {
	pulseID := GetPulseID()
	for _, dstName := range m.dsts {
		dstName := dstName
		switch pulse {
		case Lo:
			numLo++
		case Hi:
			numHi++
		default:
			panic("bad pulse")
		}

		if len(m.Modules) == 0 {
			panic(fmt.Sprint(m, "has no modules"))
		}
		dstModule := m.Modules[dstName]
		if dstModule == nil {
			panic(fmt.Sprint(m.Name, "sending to dstmodule", dstModule, "was nil"))
		}

		debugSend(m.Name, dstName, pulse)
		fn := func() (int, int) { return m.Modules[dstName].Process(pulse, m.Name) }
		pq = append(pq, QueueItem{PulseID: pulseID, Process: fn})
	}
	return numLo, numHi
}

type FlipFlopModule struct {
	BaseModule
	On bool
}

func (m *FlipFlopModule) Init() {
	m.On = false
}

func (m *FlipFlopModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	switch pulse {
	case Hi:
		return 0, 0
	case Lo:
		wasOff := !m.On
		m.On = !m.On
		if wasOff {
			return Emit(m.BaseModule, Hi)
		} else {
			return Emit(m.BaseModule, Lo)
		}
	default:
		panic("unreachable")
	}
}

type OutputModule struct {
	BaseModule
}

func (m *OutputModule) Init() {
}

func (m *OutputModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	return 0, 0
}

type ConjunctionModule struct {
	BaseModule
	remembered map[string]Pulse
}

func (m *ConjunctionModule) Init() {
	m.remembered = make(map[string]Pulse)
	for inputName := range m.Inputs {
		m.remembered[inputName] = Lo
	}
}

func (m *ConjunctionModule) remembersAllHi() bool {
	for _, pulse := range m.remembered {
		if pulse != Hi {
			return false
		}
	}
	return true
}

func (m *ConjunctionModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	m.remembered[from] = pulse
	if m.remembersAllHi() {
		return Emit(m.BaseModule, Lo)
	} else {
		return Emit(m.BaseModule, Hi)
	}
}

func ReadInput() *State {
	s := &State{Modules: map[string]BaseModule{}}

	// Step 1: Make modules
	for _, line := range aocutil.ReadLines("input.txt") {

		parts := strings.Split(line, " -> ")
		typ := parts[0][0]
		name := parts[0]

		if typ == '%' || typ == '&' {
			name = name[1:]
		}

		dsts := strings.Split(parts[1], ", ")

		if name == "broadcaster" {
			s.Modules[name] = BaseModule{
				Name:   name,
				Type:   "broadcaster",
				Inputs: map[string]struct{}{},
				dsts:   dsts,
			}
		} else if typ == '%' {
			s.Modules[name] = BaseModule{
				Name:   name,
				Type:   "flipflop",
				Inputs: map[string]struct{}{},
				dsts:   dsts,
			}
		} else if typ == '&' {
			s.Modules[name] = BaseModule{
				Name:   name,
				Type:   "conjunction",
				Inputs: map[string]struct{}{},
				dsts:   dsts,
			}
		} else {
			panic("unreachable")
		}
	}

	// Step 2: Add a button
	s.Modules["button"] = BaseModule{
		Name:   "button",
		Type:   "button",
		Inputs: map[string]struct{}{"main": {}},
		dsts:   []string{"broadcaster"},
	}

	// Step 3: Add inputs to each module
	for moduleName, module := range s.Modules {
		for _, dstName := range module.dsts {
			if s.Modules[dstName].Name == "" {
				s.Modules[dstName] = BaseModule{
					Name:   dstName,
					Type:   "output",
					Inputs: map[string]struct{}{},
				}
			}
			s.Modules[dstName].Inputs[moduleName] = struct{}{}
		}
	}
	return s
}

// broadcaster -> a, b, c
// %a -> b
// %b -> c
// %c -> inv
// &inv -> a
