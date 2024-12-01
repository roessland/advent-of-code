package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/roessland/advent-of-code/2024/day-01/day01"
	"github.com/roessland/gopkg/mathutil"
)

var (
	helpStyle       = lipgloss.NewStyle().Foreground(lipgloss.Color("#616161")).Render
	actionStyle     = lipgloss.NewStyle().Background(lipgloss.Color("#330033")).Padding(0, 1).Margin(1).Render
	paddedListStyle = lipgloss.NewStyle().Padding(0, 2).Render
	yellowFade      = []lipgloss.TerminalColor{}
)

func highlightedStyle(f fancyValue) string {
	return lipgloss.NewStyle().Background(yellowFade[f.highlightStrength]).Padding(0, 1).Render(fmt.Sprintf("%d", f.num))
}

const highlightSteps = 5

func init() {
	black, _ := colorful.Hex("#000000")
	yellow, _ := colorful.Hex("#ffff00")
	for i := 0; i < highlightSteps+1; i++ {
		t := float64(i) / highlightSteps
		c := black.BlendHcl(yellow, t).Clamped()
		lipglossColor := lipgloss.Color(c.Hex())
		yellowFade = append(yellowFade, lipglossColor)
	}
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
	case setLeftValues:
		m.leftValues = msg
	case setRightValues:
		m.rightValues = msg
	case setCurrentAction:
		m.currentAction = string(msg)
	case setCumSum:
		m.cumSum = msg
	case setDistances:
		m.distances = msg
	case tea.KeyMsg:
		// if msg.String() == "q" || msg.Type == tea.KeyCtrlC {
		return m, tea.Quit
		// }
	}
	return m, nil
}

func (m model) View() string {
	leftNums := []string{"Left\n────"}
	for _, n := range m.leftValues {
		leftNums = append(leftNums, highlightedStyle(n))
	}
	if m.leftValues == nil {
		leftNums = append(leftNums, "<nil>")
	}
	leftList := paddedListStyle(lipgloss.JoinVertical(lipgloss.Center, leftNums...))

	rightNums := []string{"Right\n─────"}
	for _, n := range m.rightValues {
		rightNums = append(rightNums, highlightedStyle(n))
	}
	if m.rightValues == nil {
		rightNums = append(rightNums, "<nil>")
	}
	rightList := paddedListStyle(lipgloss.JoinVertical(lipgloss.Center, rightNums...))

	var distancesList string
	distances := []string{"Dist\n─────"}
	for _, dist := range m.distances {
		distances = append(distances, highlightedStyle(dist))
	}
	distancesList = paddedListStyle(lipgloss.JoinVertical(lipgloss.Center, distances...))

	var cumSumList string
	cumSumNums := []string{"Total\n─────"}
	for _, cumSum := range m.cumSum {
		cumSumNums = append(cumSumNums, highlightedStyle(cumSum))
	}
	cumSumList = paddedListStyle(lipgloss.JoinVertical(lipgloss.Center, cumSumNums...))

	return lipgloss.JoinVertical(lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Top, leftList, rightList, distancesList, cumSumList),
		actionStyle(m.currentAction),
		helpStyle("Press q to quit"),
	)
}

type fancyValue struct {
	num               int
	highlightColor    int
	highlightStrength int
}

type model struct {
	currentAction string
	leftValues    []fancyValue
	rightValues   []fancyValue
	distances     []fancyValue
	cumSum        []fancyValue
}

type (
	setLeftValues    []fancyValue
	setRightValues   []fancyValue
	setCurrentAction string
	setDistances     []fancyValue
	setCumSum        []fancyValue
)

func sendAndSleep(send func(tea.Msg), msg tea.Msg, ms int) {
	send(msg)
	time.Sleep(time.Duration(ms) * time.Millisecond)
}

type ByID []fancyValue

func (a ByID) Len() int           { return len(a) }
func (a ByID) Less(i, j int) bool { return a[i].num < a[j].num }
func (a ByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type animatedSort struct {
	lenFn     func() int
	lessFn    func(i, j int) bool
	swapFn    func(i, j int)
	onCompare func(a, b int)
	onSwap    func(a, b int)
}

func (a animatedSort) Len() int           { return a.lenFn() }
func (a animatedSort) Less(i, j int) bool { a.onCompare(i, j); return a.lessFn(i, j) }
func (a animatedSort) Swap(i, j int)      { a.onSwap(i, j); a.swapFn(i, j) }
func (a animatedSort) Sort()              { sort.Sort(a) }

type AnimatedSort interface {
	sort.Interface
	OnCompare(a, b int)
	OnSwap(a, b int)
}

func AnimateSort(data sort.Interface, onCompare func(a, b int), onSwap func(a, b int)) {
	sort.Stable(animatedSort{
		lenFn:     data.Len,
		lessFn:    data.Less,
		swapFn:    data.Swap,
		onCompare: onCompare,
		onSwap:    onSwap,
	})
}

func part1(send func(tea.Msg)) {
	lefts, rights := day01.ReadInput("input-ex1.txt")
	lefts, rights = SelectSubsetForAnimation(lefts, rights)

	fLeft := []fancyValue{}
	for _, id := range lefts {
		fLeft = append(fLeft, fancyValue{int(id), 0, 0})
	}

	fRight := []fancyValue{}
	for _, id := range rights {
		fRight = append(fRight, fancyValue{int(id), 0, 0})
	}

	s := func(msg tea.Msg, sleepMs int) {
		sendAndSleep(send, msg, sleepMs)
	}

	s(setCurrentAction("Sorting..."), 0)
	s(setLeftValues(fLeft), 0)
	s(setRightValues(fRight), 0)

	var wg sync.WaitGroup
	wg.Add(2)

	sortDelay := 10 + 200/len(fLeft)
	go func() {
		AnimateSort(
			ByID(fLeft),
			func(a, b int) {
				s(setLeftValues(slices.Clone(fLeft)), sortDelay)
			},
			func(a, b int) {
				fLeft[a].highlightStrength = highlightSteps
				fLeft[b].highlightStrength = highlightSteps
				s(setLeftValues(slices.Clone(fLeft)), sortDelay)
			},
		)
		wg.Done()
	}()

	go func() {
		AnimateSort(
			ByID(fRight),
			func(a, b int) {
				s(setRightValues(slices.Clone(fRight)), sortDelay)
			},
			func(a, b int) {
				fRight[a].highlightStrength = highlightSteps
				fRight[b].highlightStrength = highlightSteps
				s(setRightValues(slices.Clone(fRight)), sortDelay)
			},
		)
		wg.Done()
	}()

	// Unsafe memory access but whatever
	go func() {
		for {
			time.Sleep(10 * time.Millisecond)
			decreaseHighlights(fLeft)
			decreaseHighlights(fRight)
			s(setLeftValues(slices.Clone(fLeft)), sortDelay)
			s(setRightValues(slices.Clone(fRight)), sortDelay)
		}
	}()

	wg.Wait()

	s(setCurrentAction("Computing distances..."), 0)

	distances := []fancyValue{}
	for i := range fLeft {
		for j := range i {
			distances[j].highlightStrength = 0
		}
		dist := fancyValue{num: mathutil.AbsInt(fRight[i].num - fLeft[i].num), highlightStrength: highlightSteps}
		distances = append(distances, dist)
		s(setDistances(distances), sortDelay*5)
	}
	distances[len(distances)-1].highlightStrength = 0
	s(setDistances(distances), sortDelay*5)

	s(setCurrentAction("Summing distances..."), 10)

	cumSums := []fancyValue{}
	soFar := 0
	for i := range fLeft {
		for j := range i {
			cumSums[j].highlightStrength = 0
		}
		soFar += distances[i].num
		cumSum := fancyValue{num: soFar, highlightStrength: highlightSteps}
		cumSums = append(cumSums, cumSum)
		s(setCumSum(cumSums), sortDelay*5)
	}
	cumSums[len(cumSums)-1].highlightStrength = 0
	s(setCumSum(cumSums), sortDelay*5)

	s(setCurrentAction("Done!"), 0)

	time.Sleep(time.Second)

	send(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
}

func idsToInts(ids []day01.ID) []int {
	ints := make([]int, len(ids))
	for i, id := range ids {
		ints[i] = int(id)
	}
	return ints
}

func SelectSubsetForAnimation(lefts, rights []day01.ID) ([]day01.ID, []day01.ID) {
	if len(lefts) < 20 {
		return lefts, rights
	}

	newLefts := lefts[0:20]
	newRights := rights[0:20]
	return newLefts, newRights
}

func decreaseHighlights(f []fancyValue) []fancyValue {
	for i := range f {
		if f[i].highlightStrength > 0 {
			f[i].highlightStrength--
		}
	}
	return f
}
