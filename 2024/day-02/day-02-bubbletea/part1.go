package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/roessland/advent-of-code/2024/day-02/day02"
)

var (
	helpStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#616161")).Render
	actionStyle     = lipgloss.NewStyle().Background(lipgloss.Color("#330033")).Padding(0, 1).Margin(1).Render
	paddedListStyle = lipgloss.NewStyle().Padding(0, 2).Render
)

func cellString(hexColor string) string {
	return lipgloss.NewStyle().Background(lipgloss.Color(hexColor)).Render(" ")
}

func main() {
	m := model{}
	p := tea.NewProgram(m)
	go part1(p.Send)
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v", err)
		os.Exit(1)
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case setLength:
		m.cellColors = make([]int, msg)
	case setPalette:
		m.palette = msg
	case increaseColor:
		m.cellColors[msg]++
	case setColor:
		m.cellColors[msg.at] = msg.col
	case tea.KeyMsg:
		// if msg.String() == "q" || msg.Type == tea.KeyCtrlC {
		return m, tea.Quit
		// }
	}
	if m.palette == nil {
		m.palette = generatePalette()
	}
	return m, nil
}

func (m model) View() string {
	cells := []string{}
	for i, col := range m.cellColors {
		newLine := ""
		if i%50 == 0 {
			newLine = "\n"
		}
		cells = append(cells, newLine+cellString(m.palette[col]))
	}
	cellsString := strings.Join(cells, "")

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.NewStyle().Render(cellsString),
		helpStyle("Press q to quit"),
	)
}

type model struct {
	cellColors []int
	palette    []string
}

func generatePalette() []string {
	palette, _ := colorful.HappyPalette(10)
	hexPalette := []string{}
	for _, c := range palette {
		hexPalette = append(hexPalette, c.Hex())
	}
	return hexPalette
}

type setPalette []string

type setLength int

type increaseColor int

type setColor struct {
	at  int
	col int
}

func sendAndSleep(send func(tea.Msg), msg tea.Msg, ms int) {
	send(msg)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

func part1(send func(tea.Msg)) {
	s := func(msg tea.Msg, sleepMs int) {
		sendAndSleep(send, msg, sleepMs)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		day02.Part2(
			"input.txt", &day02.AnimationHooks{
				SetLength: func(length int) {
					send(setLength(length))
				},
				IncreaseColor: func(index int) {
					s(increaseColor(index), 0)
				},
				SetColor: func(index, color int) {
					s(setColor{at: index, col: color}, 0)
				},
				SetPalette: func(hexPalette []string) {
					send(setPalette(hexPalette))
				},
			},
		)
		wg.Done()
	}()

	wg.Wait()

	time.Sleep(time.Second)
	send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
}
