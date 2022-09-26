package main

import (
	"bufio"
	"fmt"
	"github.com/roessland/gopkg/mathutil"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

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
	return i.Text + "\n"
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

func main() {

	allInstrs := ParseFile("input.txt")

	fk := map[int]func(z int, d int) int{}
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

nextInput:
	for {
		input := rand.Intn(99999999999999)
		inps := mathutil.ToDigitsInt(input, 10)
		if len(inps) != 14 {
			continue
		}
		for i := range inps {
			if inps[i] == 0 {f
				continue nextInput
			}
		}

		z := F[13](inps)
		if z == 0 {
			fmt.Println("OHLY BALLS", inps, z)
		}
	}

}
