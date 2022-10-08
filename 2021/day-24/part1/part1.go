package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"github.com/roessland/gopkg/sliceutil"
	"io"
	"log"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

var a, b, c []int

func init() {
	allInstrs := ParseFile("input.txt")
	a, b, c = ExtractABC(allInstrs)
}

type InstructionType int

const (
	Inp InstructionType = iota
	AddC
	AddP
	MulC
	MulP
	DivC
	DivP
	ModC
	ModP
	EqlC
	EqlP
)

type Reg struct {
	W, X, Y, Z int
}

func (r *Reg) Set(name int, val int) {
	switch name {
	case 'w':
		r.W = val
	case 'x':
		r.X = val
	case 'y':
		r.Y = val
	case 'z':
		r.Z = val
	default:
		panic("no such register")
	}
}

func (r *Reg) Get(name int) int {
	switch name {
	case 'w':
		return r.W
	case 'x':
		return r.X
	case 'y':
		return r.Y
	case 'z':
		return r.Z
	default:
		panic("no such register")
	}
}

func Compute(instrs []Instruction, inps []int, reg Reg) Reg {
	inpCount := 0
	for _, instr := range instrs {
		switch instr.Type {
		case Inp:
			reg.Set(instr.A, inps[inpCount])
			inpCount++
		case AddC:
			reg.Set(instr.A, reg.Get(instr.A)+instr.B)
		case AddP:
			reg.Set(instr.A, reg.Get(instr.A)+reg.Get(instr.B))
		case MulC:
			reg.Set(instr.A, reg.Get(instr.A)*instr.B)
		case MulP:
			reg.Set(instr.A, reg.Get(instr.A)*reg.Get(instr.B))
		case DivC:
			reg.Set(instr.A, reg.Get(instr.A)/instr.B)
		case DivP:
			reg.Set(instr.A, reg.Get(instr.A)/reg.Get(instr.B))
		case ModC:
			reg.Set(instr.A, reg.Get(instr.A)%instr.B)
		case ModP:
			reg.Set(instr.A, reg.Get(instr.A)%reg.Get(instr.B))
		case EqlC:
			if reg.Get(instr.A) == instr.B {
				reg.Set(instr.A, 1)
			} else {
				reg.Set(instr.A, 0)
			}
		case EqlP:
			if reg.Get(instr.A) == reg.Get(instr.B) {
				reg.Set(instr.A, 1)
			} else {
				reg.Set(instr.A, 0)
			}
		}
	}
	return reg
}

func Compute2(instrs []Instruction, inps []int) Reg {
	reg := Reg{}

	inpCount := 0
	lastModified := Reg{}
	for i, instr := range instrs {
		fmt.Printf("%d[\"%s\"]\n", i, instr.Text)
		deps := []int{}
		switch instr.Type {
		case Inp:
			reg.Set(instr.A, inps[inpCount])
			deps = append(deps, 1000+inpCount)
			lastModified.Set(instr.A, i)
			inpCount++
		case AddC:
			reg.Set(instr.A, reg.Get(instr.A)+instr.B)
			deps = append(deps, lastModified.Get(instr.A))
			lastModified.Set(instr.A, i)
		case AddP:
			reg.Set(instr.A, reg.Get(instr.A)+reg.Get(instr.B))
			deps = append(deps, lastModified.Get(instr.A))
			deps = append(deps, lastModified.Get(instr.B))
			lastModified.Set(instr.A, i)
		case MulC:
			reg.Set(instr.A, reg.Get(instr.A)*instr.B)
			if instr.B != 0 {
				deps = append(deps, lastModified.Get(instr.A))
			}
			lastModified.Set(instr.A, i)
		case MulP:
			reg.Set(instr.A, reg.Get(instr.A)*reg.Get(instr.B))
			deps = append(deps, lastModified.Get(instr.A))
			deps = append(deps, lastModified.Get(instr.B))
			lastModified.Set(instr.A, i)
		case DivC:
			reg.Set(instr.A, reg.Get(instr.A)/instr.B)
			deps = append(deps, lastModified.Get(instr.A))
			lastModified.Set(instr.A, i)
		case DivP:
			reg.Set(instr.A, reg.Get(instr.A)/reg.Get(instr.B))
			deps = append(deps, lastModified.Get(instr.A))
			deps = append(deps, lastModified.Get(instr.B))
			lastModified.Set(instr.A, i)
		case ModC:
			reg.Set(instr.A, reg.Get(instr.A)%instr.B)
			deps = append(deps, lastModified.Get(instr.A))
			lastModified.Set(instr.A, i)
		case ModP:
			reg.Set(instr.A, reg.Get(instr.A)%reg.Get(instr.B))
			deps = append(deps, lastModified.Get(instr.A))
			deps = append(deps, lastModified.Get(instr.B))
			lastModified.Set(instr.A, i)
		case EqlC:
			if reg.Get(instr.A) == instr.B {
				reg.Set(instr.A, 1)
			} else {
				reg.Set(instr.A, 0)
			}
			deps = append(deps, lastModified.Get(instr.A))
			lastModified.Set(instr.A, i)
		case EqlP:
			if reg.Get(instr.A) == reg.Get(instr.B) {
				reg.Set(instr.A, 1)
			} else {
				reg.Set(instr.A, 0)
			}
			deps = append(deps, lastModified.Get(instr.A))
			deps = append(deps, lastModified.Get(instr.B))
			lastModified.Set(instr.A, i)
		default:
			panic("bruh")
		}

		if len(deps) == 1 {
			fmt.Printf("%d --> %d\n", i, deps[0])
		}
		if len(deps) == 2 {
			fmt.Printf("%d --> %d & %d\n", i, deps[0], deps[1])
		}

	}
	fmt.Printf("z --> %d\n", lastModified.Z)

	os.Exit(0)
	return reg
}

type Instruction struct {
	Type InstructionType
	A    int
	B    int
	Text string
}

func (i Instruction) String() string {
	return i.Text
}

func ParseFile(filename string) []Instruction {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	return ParseInput(f)
}

func ParseString(s string) []Instruction {
	r := strings.NewReader(s)
	return ParseInput(r)
}

func ParseInput(r io.Reader) []Instruction {
	var instructions []Instruction
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		codeComment := strings.SplitN(line, "//", 2)
		parts := strings.Split(strings.TrimSpace(codeComment[0]), " ")
		name := parts[0]
		if len(parts) == 1 {
			continue
		}
		a := parts[1]
		b := ""
		var instr Instruction
		instr.Text = line
		var isPtr bool
		var bInt int
		if len(parts) > 2 {
			b = parts[2]
			if len(b) == 0 {
				panic(fmt.Sprintf("wtf: %v", parts))
			}
			isPtr = 'a' <= b[0] && b[0] <= 'z'
			bInt, _ = strconv.Atoi(b)
		}

		type Op struct {
			Name  string
			IsPtr bool
		}
		op := Op{name, isPtr}
		switch op {
		case Op{"inp", false}:
			instr.Type = Inp
			instr.A = int(a[0])
		case Op{"add", false}:
			instr.Type = AddC
			instr.A = int(a[0])
			instr.B = bInt
		case Op{"add", true}:
			instr.Type = AddP
			instr.A = int(a[0])
			instr.B = int(b[0])
		case Op{"mul", false}:
			instr.Type = MulC
			instr.A = int(a[0])
			instr.B = bInt
		case Op{"mul", true}:
			instr.Type = MulP
			instr.A = int(a[0])
			instr.B = int(b[0])
		case Op{"div", false}:
			instr.Type = DivC
			instr.A = int(a[0])
			instr.B = bInt
		case Op{"div", true}:
			instr.Type = DivP
			instr.A = int(a[0])
			instr.B = int(b[0])
		case Op{"mod", false}:
			instr.Type = ModC
			instr.A = int(a[0])
			instr.B = bInt
		case Op{"mod", true}:
			instr.Type = ModP
			instr.A = int(a[0])
			instr.B = int(b[0])
		case Op{"eql", false}:
			instr.Type = EqlC
			instr.A = int(a[0])
			instr.B = bInt
		case Op{"eql", true}:
			instr.Type = EqlP
			instr.A = int(a[0])
			instr.B = int(b[0])
		}
		instructions = append(instructions, instr)
	}
	return instructions
}

type Pair struct {
	A, B int
}

type Triplet struct {
	A, B, C int
}

type ModuleFunc map[int]func(z int, d int) int
type CumulativeFunc map[int]func([]int) int

func CompareModules(allInstrs []Instruction) {
	for k := -1; k < 14; k++ {
		fmt.Printf("|k=%d", k)
	}
	fmt.Println("|")
	for k := -1; k < 14; k++ {
		fmt.Printf("|----")
	}
	fmt.Println("|")

	for i := 0; i < 18; i++ {
		fmt.Printf("|%d|", i)
		for k := 0; k < 14; k++ {
			instr := allInstrs[k*18+i]
			fmt.Print(instr, "|")
		}
		fmt.Println()
	}
}

func ExtractABC(allInstrs []Instruction) (as, bs, cs []int) {
	for k := 0; k < 14; k++ {
		instr := allInstrs[k*18+4]
		as = append(as, instr.B)
	}

	for k := 0; k < 14; k++ {
		instr := allInstrs[k*18+5]
		bs = append(bs, instr.B)
	}

	for k := 0; k < 14; k++ {
		instr := allInstrs[k*18+15]
		cs = append(cs, instr.B)
	}

	return as, bs, cs
}

func GetModulesWithMemoization(allInstrs []Instruction) (fk ModuleFunc, Fk CumulativeFunc) {
	fk = make(ModuleFunc)
	cache := map[Triplet]int{}
	for i := 0; i < 14; i++ {
		fk[i] = (func(i_ int) func(z, in int) int {
			return func(z, in int) int {
				cached, ok := cache[Triplet{i_, z, in}]
				if ok {
					return cached
				}
				instrs := allInstrs[i_*18 : (i_+1)*18]
				ret := Compute(instrs, []int{in}, Reg{Z: z}).Z
				cache[Triplet{i_, z, in}] = ret
				return ret
			}
		})(i)
	}

	F := map[int]func([]int) int{}
	F[0] = func(is []int) int {
		return fk[0](0, is[0])
	}
	for i := 1; i < 14; i++ {
		F[i] = func(i_ int) func([]int) int {
			return func(is []int) int {
				return fk[i_](F[i_-1](is), is[i_])
			}
		}(i)
	}

	return fk, F
}

func RandomizeInput(inps []int) {
	for i := 0; i < len(inps); i++ {
		inps[i] = 1 + rand.Intn(9)
	}
}

func PartialRandomizeInput(inps []int) {
	inps[rand.Intn(len(inps))] = 1 + rand.Intn(9)
}

func FirstAlgo() {
	validInputs := map[int]map[int]bool{}
	validInputs[14] = map[int]bool{0: true}

	zMax := 10760000
	for k := 13; k >= 6; k-- {
		validInputs[k] = map[int]bool{}
		for zk := 0; zk < zMax; zk++ {
			for dk := 1; dk <= 9; dk++ {
				zNext := f(zk, dk, k)
				if validInputs[k+1][zNext] {
					validInputs[k][zk] = true
				}
			}
		}
	}

	for k := 0; k < 14; k++ {
		if zk, ok := validInputs[k]; ok {
			fmt.Println(k, len(zk), " ")
		}
	}
}

func asInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func f(z, i, k int) int {
	a := a[k]
	b := b[k]
	c := c[k]
	return z/a*(25*(asInt(z%26+b != i))+1) + (i+c)*asInt(z%26+b != i)
}

func Codegen(a, b, c, i []int) (z int) {
	for k := 0; k < 14; k++ {
		z = f(z, i[k], k)
	}
	return z
}

func computeZ1() {
	// [{1 13} {2 346} {3 9005} {4 234136} {5 9005} {6 234147} {7 6087836} {8 234147} {9 9005} {10 346} {11 9005} {12 346} {13 13} {14 0}] [1 1 8 4 1 2 3 1 1 1 7 1 8 9]
	//11841231117189
	// 1 1 8
	for d := 1; d <= 9; d++ {
		fmt.Println(d, f(346, d, 13))
	}
}

func F13(i []int) (z int) {
	asInt := func(b bool) int {
		if b {
			return 1
		}
		return 0
	}

	z = z/1*(25*(asInt(z%26+14 != i[0]))+1) + (i[0]+12)*asInt(z%26+14 != i[0])
	z = z/1*(25*(asInt(z%26+15 != i[1]))+1) + (i[1]+7)*asInt(z%26+15 != i[1])
	z = z/1*(25*(asInt(z%26+12 != i[2]))+1) + (i[2]+1)*asInt(z%26+12 != i[2])
	z = z/1*(25*(asInt(z%26+11 != i[3]))+1) + (i[3]+2)*asInt(z%26+11 != i[3])
	z = z/26*(25*(asInt(z%26+-5 != i[4]))+1) + (i[4]+4)*asInt(z%26+-5 != i[4])
	z = z/1*(25*(asInt(z%26+14 != i[5]))+1) + (i[5]+15)*asInt(z%26+14 != i[5])
	z = z/1*(25*(asInt(z%26+15 != i[6]))+1) + (i[6]+11)*asInt(z%26+15 != i[6])
	z = z/26*(25*(asInt(z%26+-13 != i[7]))+1) + (i[7]+5)*asInt(z%26+-13 != i[7])
	z = z/26*(25*(asInt(z%26+-16 != i[8]))+1) + (i[8]+3)*asInt(z%26+-16 != i[8])
	z = z/26*(25*(asInt(z%26+-8 != i[9]))+1) + (i[9]+9)*asInt(z%26+-8 != i[9])
	z = z/1*(25*(asInt(z%26+15 != i[10]))+1) + (i[10]+2)*asInt(z%26+15 != i[10])
	z = z/26*(25*(asInt(z%26+-8 != i[11]))+1) + (i[11]+3)*asInt(z%26+-8 != i[11])
	z = z/26*(25*(asInt(z%26+0 != i[12]))+1) + (i[12]+3)*asInt(z%26+0 != i[12])
	z = z/26*(25*(asInt(z%26+-4 != i[13]))+1) + (i[13]+11)*asInt(z%26+-4 != i[13])
	return z
}

func Bruteforce() {
	//for nthread := 0; nthread < 1; nthread++ {
	//	go func() {
	inps := []int{1, 1, 9, 9, 6, 2, 3, 1, 1, 1, 7, 1, 8, 9}
	//inps := []int{1, 1, 8, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5}
	for i := 0; ; i++ {
		PartialRandomizeInput(inps[5:])
		z := F13(inps)
		if z == 0 {
			fmt.Println("z", z, "inps", inps)

			inps64 := make([]int64, len(inps))
			for d := 0; d < len(inps64); d++ {
				inps64[d] = int64(inps[d])
			}
			input10 := mathutil.FromDigits(inps64[0:], 10)
			input26 := mathutil.ToDigitsInt(input10, 26)
			fmt.Println("base 26", input26)
		}

	}
	//	}()
	//}
	//time.Sleep(1000 * time.Second)
}

type State struct {
	K, Z int
}

func ComputePruningLimits() []int {
	limits := make([]int, len(a)+1)
	limits[13] = 26
	for k := 12; k >= 0; k-- {
		if a[k] == 26 {
			limits[k] = limits[k+1]*26 + 26
		} else {
			limits[k] = limits[k+1] + 26
		}
	}
	return limits
}

func GraphSearch(max bool) int {
	limits := ComputePruningLimits()
	visited := map[State]bool{}
	S := map[State]struct{}{}
	S[State{0, 0}] = struct{}{}
	prev := map[State]State{}
	digit := map[State]int{}
	var results []int

	for len(S) > 0 {
		var nodeCurr State
		for s := range S {
			nodeCurr = s
			break
		}
		delete(S, nodeCurr)
		visited[nodeCurr] = true

		k := nodeCurr.K
		z := nodeCurr.Z

		// Pruning branches with too high / unrecoverable z value.
		if z > limits[k] {
			continue
		}

		if k == 14 {
			if z == 0 {
				digits := []int{}
				for nodeCurr != (State{0, 0}) {
					digits = append(digits, digit[nodeCurr])
					nodeCurr = prev[nodeCurr]
				}
				sliceutil.ReverseInt(digits)
				results = append(results, mathutil.FromDigitsInt(digits, 10))
			}
			continue
		}
		for i := 1; i <= 9; i++ {
			zNext := f(z, i, k)
			nodeNext := State{k + 1, zNext}
			switchedDigit := false
			if digit[nodeNext] == 0 {
				prev[nodeNext] = nodeCurr
				digit[nodeNext] = i
				switchedDigit = true
			} else if max && i >= digit[nodeNext] {
				prev[nodeNext] = nodeCurr
				digit[nodeNext] = i
				switchedDigit = true
			} else if !max && i <= digit[nodeNext] {
				prev[nodeNext] = nodeCurr
				digit[nodeNext] = i
				switchedDigit = true
			}
			if switchedDigit { //
				S[nodeNext] = struct{}{}
			}
		}
	}
	sort.Ints(results)
	if max {
		return results[len(results)-1]
	} else {
		return results[0]
	}
}

func main() {
	var t0 time.Time

	t0 = time.Now()
	fmt.Print("Part 1: ")
	a := GraphSearch(true)
	fmt.Println(a == 12996997829399, time.Since(t0))

	t0 = time.Now()
	fmt.Print("Part 2: ")
	b := GraphSearch(false)
	fmt.Println(b == 11841231117189, time.Since(t0))
}

// 12996997829399 part 1
// 11841231117189 part 2
