package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/roessland/gopkg/priorityqueue2"
)

func init() {
	tea.LogToFile("log.txt", "")
}

var counter int

type tickMsg time.Time

type model struct {
	bricks        []Brick
	maxZ          int
	maxY          int
	minY          int
	minX, maxX    int
	windowWidth   int
	windowHeight  int
	scrollZ       int
	part1Solution int
	part2Solution int
}

func initialModel(bricks []Brick) model {
	// Useful for rendering
	maxZ := 0
	maxY := 0
	minY := bricks[0].Origin.Y
	minX := bricks[0].Origin.X
	maxX := 0
	for _, b := range bricks {
		if b.Origin.Z+b.Size.Z > maxZ {
			maxZ = b.Origin.Z + b.Size.Z
		}
		if b.Origin.Y+b.Size.Y > maxY {
			maxY = b.Origin.Y + b.Size.Y
		}
		if b.Origin.Y < minY {
			minY = b.Origin.Y
		}
		if b.Origin.X < minX {
			minX = b.Origin.X
		}
		if b.Origin.X+b.Size.X > maxX {
			maxX = b.Origin.X + b.Size.X
		}

	}

	return model{
		bricks:        bricks,
		maxZ:          maxZ,
		minY:          minY,
		maxY:          maxY,
		minX:          minX,
		maxX:          maxX,
		windowWidth:   60,
		windowHeight:  30,
		scrollZ:       0,
		part1Solution: 0,
		part2Solution: 0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.scrollZ++
			return m, nil
		case "down":
			m.scrollZ--
			return m, nil
		case "pgup":
			m.scrollZ += 10
			return m, nil
		case "pgdown":
			m.scrollZ -= 10
			return m, nil
		case " ":
			var supportSets map[int]map[int]bool
			m.bricks, m.part1Solution, supportSets = Fall(m.bricks)
			m.part2Solution = Part2(m.bricks, supportSets)
			return m, nil
		}
	case tickMsg:
		return m, nil
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width
		return m, nil
	}
	return m, nil
}

func Part2(bricks []Brick, supportSets map[int]map[int]bool) int {
	// Find essential bricks, those that are the sole support in any support set
	essentialBricks := map[int]bool{}
	for _, supportSet := range supportSets {
		if len(supportSet) == 1 {
			for supportID := range supportSet {
				essentialBricks[supportID] = true
			}
		}
	}

	// supportSets[brID] contains what brID is standing on
	//
	// upstream[brID] contains what is standing on brID (and maybe others too!)

	upstream := map[int]map[int]bool{}
	for brID, supportSet := range supportSets {
		for supportID := range supportSet {
			if upstream[supportID] == nil {
				upstream[supportID] = map[int]bool{}
			}
			upstream[supportID][brID] = true
		}
	}

	var numDestroyed func(int) int

	destroyed := map[int]bool{}
	var toCheck *priorityqueue2.PriorityQueue[int, int]
	numDestroyed = func(startingAt int) int {
		sum := 0
		for toCheck.Len() > 0 {
			// Check if still supported
			brID := toCheck.Pop()
			log.Printf("Checking %v", brID)
			numSupport := 0
			for supportID := range supportSets[brID] {
				if !destroyed[supportID] {
					numSupport++
				}
			}

			// Destroy if unsupported, and queue its upstream.
			// Process low Z first.
			if numSupport == 0 || brID == startingAt {
				if !destroyed[brID] {
					destroyed[brID] = true
					sum++
				}

				for upstreamBrickID := range upstream[brID] {
					if !destroyed[upstreamBrickID] {
						toCheck.Push(upstreamBrickID, bricks[upstreamBrickID].Origin.Z)
					}
				}
			}
		}

		return sum
	}

	totalDamage := 0
	for brickIDToDestroy := range essentialBricks {
		destroyed = map[int]bool{brickIDToDestroy: true}
		toCheck = priorityqueue2.New[int, int]()
		toCheck.Push(brickIDToDestroy, 0)
		log.Print("Destroying brick", brickIDToDestroy, "also destroys", upstream[brickIDToDestroy])
		totalDamage += numDestroyed(brickIDToDestroy)
		log.Println("Total damage so far:", totalDamage)
	}
	return totalDamage
}

func (m model) View() string {
	counter++
	var buf bytes.Buffer

	lastVisibleZ := m.scrollZ
	firstVisibleZ := lastVisibleZ + m.windowHeight + m.scrollZ - 2

	bgStyle := lipgloss.NewStyle().Background(lipgloss.Color("#222222")).Foreground(lipgloss.Color("#222222"))

	for z := firstVisibleZ; z >= lastVisibleZ; z-- {
		for x := m.minX; x <= m.maxX; x++ {
			xz := Coord2{x, z}
			foundBrick := false
			for _, br := range m.bricks {
				if br.ContainsXZ(xz) {
					color := getColor(br.ID, br.Origin.Y)
					buf.WriteString(lipgloss.NewStyle().Background(color).Render(br.Name))
					buf.WriteString(lipgloss.NewStyle().UnsetBackground().Render())
					foundBrick = true
					break
				}
			}

			if !foundBrick {
				var fillStr string
				if counter%2 == 0 {
					fillStr = "22"
				} else {
					fillStr = "11"
				}
				buf.WriteString(bgStyle.Render(fillStr))
			}
		}
		buf.WriteString(lipgloss.NewStyle().UnsetBackground().Render("\n"))
	}
	buf.WriteString(fmt.Sprintf("Part 1 solution: %d\n", m.part1Solution))
	buf.WriteString(fmt.Sprintf("Part 2 solution: %d\n", m.part2Solution))

	return buf.String()
}

func main() {
	bricks := ReadInput()

	// p := tea.NewProgram(initialModel(bricks))
	p := tea.NewProgram(initialModel(bricks), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %s\n", err)
		os.Exit(1)
	}
}
