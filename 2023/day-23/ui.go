package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/timer"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func init() {
	tea.LogToFile("log.txt", "")
}

var counter int

type tickMsg time.Time

type model struct {
	m             *Map
	windowWidth   int
	windowHeight  int
	part1Solution int
	part2Solution int
	timer         timer.Model
	visited       map[Pos]int
}

func initialModel(m *Map, visited map[Pos]int) model {
	return model{
		m:             m,
		windowWidth:   60,
		windowHeight:  30,
		part1Solution: 0,
		part2Solution: 0,
		timer:         timer.New(50 * time.Hour),
		visited:       visited,
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
			return m, nil
		case "down":
			return m, nil
		case " ":
			return m, m.timer.Start()
		}
	case tickMsg:
		fmt.Println("tick")
		return m, nil
	case tea.WindowSizeMsg:
		m.windowHeight = msg.Height
		m.windowWidth = msg.Width
		return m, nil
	case timer.TickMsg:
		fmt.Println("tickiiii")
		var cmd tea.Cmd
		m.timer, cmd = m.timer.Update(msg)
		return m, tea.Batch(cmd)
	}
	return m, nil
}

func (m model) View() string {
	counter++
	var buf bytes.Buffer

	visited := make(map[Pos]int)
	mu.RLock()
	for k, v := range m.visited {
		visited[k] = v
	}
	mu.RUnlock()

	for y := 0; y < m.windowHeight; y++ {
		for x := 0; x < m.windowWidth; x++ {
			t := m.m.At(Pos{y, x})
			switch t {
			case '#':
				buf.WriteString(lipgloss.NewStyle().Background(lipgloss.Color("#222222")).Render(" "))
			default:
				if visited[Pos{y, x}] == 1 {
					buf.WriteString(lipgloss.NewStyle().Background(lipgloss.Color("#9922c3")).Render(" "))
				} else {
					buf.WriteString(lipgloss.NewStyle().Background(lipgloss.Color("#993311")).Render(" "))
				}
			}
		}
		buf.WriteString(lipgloss.NewStyle().UnsetBackground().Render("\n"))
	}

	return buf.String()
}
