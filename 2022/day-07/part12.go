package main

import (
	"bufio"
	"embed"
	_ "embed"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt input-ex.txt
var inputFiles embed.FS

// part 1: 55 min
// part 2: 5 min
func main() {
	// Part 1: 55 min spent
	input := ReadInput()
	device := NewDevice()
	for _, exec := range input {
		device.HandleCommand(exec.Cmd, exec.Arg, exec.Output)
	}

	device.RootDir.ComputeTotalSizes()

	answerPart1 := 0
	device.RootDir.Walk(func(node *Node) {
		if node.Type == NodeTypeDirectory && node.TotalSize <= 100000 {
			answerPart1 += node.TotalSize
		}
	})
	fmt.Println(answerPart1)

	// Part 2: 5 min spent
	diskSize := 70000000
	diskUsed := device.RootDir.TotalSize
	requiredUnused := 30000000
	answerPart2 := math.MaxInt32
	device.RootDir.Walk(func(node *Node) {
		if node.Type != NodeTypeDirectory {
			return
		}
		if diskSize-diskUsed+node.TotalSize >= requiredUnused {
			if node.TotalSize < answerPart2 {
				answerPart2 = node.TotalSize
			}
		}
	})
	fmt.Println(answerPart2)
}

type Device struct {
	CurrDir *Node
	RootDir *Node
}

func (device *Device) HandleCommand(cmd string, arg string, output []string) {
	switch cmd {
	case "cd":
		device.HandleCd(arg, output)
	case "ls":
		device.HandleLs(arg, output)
	default:
		panic("unknown command " + cmd)
	}
}

func (device *Device) HandleCd(dst string, output []string) {
	if len(output) > 0 {
		panic("cd shouldn't have output")
	}
	if dst == "/" {
		device.CurrDir = device.RootDir
		return
	}
	if dst == ".." {
		device.CurrDir = device.CurrDir.Parent
		return
	}
	device.CurrDir = device.CurrDir.Children[dst]
}

func (device *Device) HandleLs(arg string, output []string) {
	if len(arg) > 0 {
		panic("ls shouldn't have arg")
	}
	for _, line := range output {
		parts := strings.Split(line, " ")
		if parts[0] == "dir" {
			device.AddNode(Node{
				Name: parts[1],
				Type: NodeTypeDirectory,
			})
		} else {
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				panic("invalid file size")
			}
			device.AddNode(Node{
				Name: parts[1],
				Type: NodeTypeFile,
				Size: size,
			})
		}
	}
}

func (device *Device) AddNode(node Node) {
	if device.CurrDir.Children[node.Name] != nil {
		log.Print("dir already exists")
		return
	}
	if device.CurrDir.Children == nil {
		device.CurrDir.Children = make(map[string]*Node)
	}
	device.CurrDir.Children[node.Name] = &Node{
		node.Name,
		node.Size,
		node.Type,
		device.CurrDir,
		nil,
		node.Size,
	}
}

func NewDevice() *Device {
	rootNode := &Node{
		Name: "/",
		Type: NodeTypeDirectory,
	}
	return &Device{CurrDir: rootNode, RootDir: rootNode}
}

type Node struct {
	Name      string
	Size      int
	Type      NodeType
	Parent    *Node
	Children  map[string]*Node
	TotalSize int
}

func (node *Node) ComputeTotalSizes() {
	if node.Type == NodeTypeFile {
		return
	}
	totalSize := 0
	for _, child := range node.Children {
		child.ComputeTotalSizes()
		totalSize += child.TotalSize
	}
	node.TotalSize = totalSize
}

func (node *Node) Walk(f func(*Node)) {
	f(node)
	for _, child := range node.Children {
		child.Walk(f)
	}
}

type NodeType string

const NodeTypeDirectory = "dir"
const NodeTypeFile = "f"

type Execution struct {
	Cmd    string
	Arg    string
	Output []string
}

func ReadInput() []Execution {
	f, err := inputFiles.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var input []Execution
	var currOutput []string
	currCmd, currArg := "", ""
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println("parsing", line)
		if line[0] == '$' {
			if currCmd != "" {
				input = append(input, Execution{currCmd, currArg, currOutput})
			}
			parts := strings.Split(line, " ")
			currCmd = parts[1]
			if len(parts) >= 3 {
				currArg = parts[2]
			} else {
				currArg = ""
			}
			currOutput = nil
		} else {
			currOutput = append(currOutput, line)
		}
	}
	input = append(input, Execution{currCmd, currArg, currOutput})
	return input
}
