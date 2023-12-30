package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/roessland/advent-of-code/2023/aocutil"
	"github.com/roessland/gopkg/mathutil/crt"
)

var buttonPressed int

/*
&cs -> rx
cs must become hhhh.

- &cs
  - &kh
    - &sk
      -
  - &lz
    - &sd
      -
  - &tg
    - &pl
      -
  - &hn
    - &zv
      -
*/

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

func printStateRow(order []string, modules map[string]Module) {
	widths := make(map[string]int)
	states := make(map[string]string)

	// Get state and column width
	for _, name := range order {
		states[name], widths[name] = modules[name].State()
	}

	// Print state
	for _, name := range order {
		width := widths[name]
		fmt.Printf("| %"+strconv.Itoa(width)+"s ", states[name])
	}
	fmt.Println()
}

func printStateHeaders(order []string, modules map[string]Module) {
	widths := make(map[string]int)
	states := make(map[string]string)

	// Get state and column width
	for _, name := range order {
		states[name], widths[name] = modules[name].State()
	}

	// Print table headers
	for _, name := range order {
		if name == "broadcaster" {
			name = "bcr"
		}
		if name == "button" {
			name = "btn"
		}
		width := widths[name]
		fmt.Printf("| %"+strconv.Itoa(width)+"s ", name)
	}
	fmt.Println()
	for _, name := range order {
		fmt.Print(strings.Repeat("-", widths[name]+3))
	}
	fmt.Println()
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

	printStateHeaders(input.ModuleNames, modules)
	numLo, numHi := 0, 0
	for i := 0; i < 1000000000; i++ {
		pq = append(pq, QueueItem{
			Process: func() (int, int) { return modules["button"].Process(Lo, "main") },
			PulseID: GetPulseID(),
		})
		buttonPressed++
		for len(pq) > 0 {
			qi := pq[0]
			pq = pq[1:]
			l, h := qi.Process()
			numLo += l
			numHi += h
			FixPQ()
		}

		if i%10000000 == 0 {
			printStateHeaders(input.ModuleNames, modules)
		}
		if i%1000000 == 0 {
			printStateRow(input.ModuleNames, modules)
		}
	}

	fmt.Println("numHi:", numHi)
	fmt.Println("numLo:", numLo)
	fmt.Println("Part 1:", numLo*numHi)
}

type State struct {
	ModuleNames []string
	Modules     map[string]BaseModule
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
	State() (string, int)
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

type QueueItem struct {
	Process func() (int, int)
	PulseID int
}

var nextPulseID int

func GetPulseID() int {
	nextPulseID++
	return nextPulseID
}

// Note: The priority queue must be stable!
// If adding a, b, c with the same priority, they must be popped in the same order.
// It has a max length of ~100 so just stablesorting it on every push is fine.
var pq []QueueItem

func FixPQ() {
	slices.SortStableFunc(pq, func(a, b QueueItem) int {
		return cmp.Compare(a.PulseID, b.PulseID)
	})
}

func Emit(m BaseModule, pulse Pulse) (numLo, numHi int) {
	if m.Name == "hn" && pulse == Hi {
		fmt.Println("hn emitted hi", buttonPressed, buttonPressed%4013) // 409326
	}
	if m.Name == "lz" && pulse == Hi {
		fmt.Println("lz emitted hi", buttonPressed, buttonPressed%3917) // 411285
	}
	if m.Name == "kh" && pulse == Hi {
		fmt.Println("kh emitted hi", buttonPressed, buttonPressed%3889) // 408345
	}
	if m.Name == "tg" && pulse == Hi {
		fmt.Println("tg emitted hi", buttonPressed, buttonPressed%3769) // 414590
	}

	// x = 409326 mod 4013
	// x = 411285 mod 3917
	// x = 408345 mod 3889
	// x = 414590 mod 3769
	x, N := crt.CRTCoerce([]int64{409326, 411285, 408345, 414590}, []int64{4013, 3917, 3889, 3769})
	fmt.Println("Part 2:", x+N)

	fmt.Println("Part 2 (without CRT): ", 4013*3917*3889*3769)
	os.Exit(0)

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

type ButtonModule struct {
	BaseModule
}

func (m *ButtonModule) Init() {
}

func (m *ButtonModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	return Emit(m.BaseModule, pulse)
}

func (m *ButtonModule) State() (string, int) {
	return "", 3
}

var _ Module = &ButtonModule{}

type BroadcastModule struct {
	BaseModule
}

// Init implements Module.
func (*BroadcastModule) Init() {
}

func (m *BroadcastModule) State() (string, int) {
	return "", 3
}

var _ Module = &BroadcastModule{}

func (m *BroadcastModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	debugReceive(m.Name, "flipflop", pulse)
	return Emit(m.BaseModule, pulse)
}

type FlipFlopModule struct {
	BaseModule
	On bool
}

var _ Module = &FlipFlopModule{}

func (m *FlipFlopModule) Init() {
	m.On = false
}

func (m *FlipFlopModule) State() (string, int) {
	if m.On {
		return "on", 3
	} else {
		return "off", 3
	}
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

func (m *OutputModule) State() (string, int) {
	return "", 3
}

func (m *OutputModule) Process(pulse Pulse, from string) (numLo, numHi int) {
	return 0, 0
}

var _ Module = &OutputModule{}

type ConjunctionModule struct {
	BaseModule
	remembered     map[string]Pulse
	rememberedSrcs []string
}

var _ Module = &ConjunctionModule{}

func (m *ConjunctionModule) Init() {
	m.remembered = make(map[string]Pulse)
	for inputName := range m.Inputs {
		m.remembered[inputName] = Lo
		m.rememberedSrcs = append(m.rememberedSrcs, inputName)
	}
	sort.Strings(m.rememberedSrcs)
}

func (m *ConjunctionModule) State() (string, int) {
	s := make([]byte, len(m.rememberedSrcs))
	for i, src := range m.rememberedSrcs {
		if m.remembered[src] == Hi {
			s[i] = 'h'
		} else {
			s[i] = 'l'
		}
	}
	width := len(m.rememberedSrcs)
	if width < 3 {
		width = 3
	}
	return string(s), width
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

		s.ModuleNames = append(s.ModuleNames, name)
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
	s.ModuleNames = append(s.ModuleNames, "button")
	s.Modules["button"] = BaseModule{
		Name:   "button",
		Type:   "button",
		Inputs: map[string]struct{}{"main": {}},
		dsts:   []string{"broadcaster"},
	}

	// Step 3: Add inputs to each module.
	for moduleName, module := range s.Modules {
		for _, dstName := range module.dsts {
			if s.Modules[dstName].Name == "" {
				s.ModuleNames = append(s.ModuleNames, dstName)
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
