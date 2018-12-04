package main

import "fmt"
import "strconv"
import "sort"
import "log"
import "bufio"
import "os"
import "strings"
import "io/ioutil"
import "encoding/csv"
import "github.com/roessland/gopkg/disjointset"

type GuardMinute struct {
	Guard  int
	Minute int
}

func main() {
	guardSleepMinutes := map[int]int{}
	guardSleepMinuteFreqs := map[GuardMinute]int{}
	_ = guardSleepMinutes
	var currGuard int
	var isSleeping bool
	_ = isSleeping
	var fellAsleepTime int
	var wakeupTime int
	buf, _ := ioutil.ReadFile("input.txt")
	for _, line := range strings.Split(string(buf), "\n") {
		if len(line) == 0 {
			continue
		}
		line = strings.Replace(line, "-", " ", -1)
		line = strings.Replace(line, "#", "", -1)
		line = strings.Replace(line, "[", "", -1)
		line = strings.Replace(line, "]", "", -1)
		line = strings.Replace(line, ":", " ", -1)
		var month, day, hour, minute, guard int
		if strings.Contains(line, "Guard") {
			_, err := fmt.Sscanf(line, "1518 %d %d %d %d Guard %d begins shift", &month, &day, &hour, &minute, &guard)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(month, day, hour, minute, guard)
			currGuard = guard
			isSleeping = false
			continue
		}

		if strings.Contains(line, "falls") {
			_, err := fmt.Sscanf(line, "1518 %d %d %d %d falls asleep", &month, &day, &hour, &minute)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("asleep", month, day, hour, minute, guard)
			isSleeping = true
			fellAsleepTime = minute
			continue
		}

		if strings.Contains(line, "wakes") {
			_, err := fmt.Sscanf(line, "1518 %d %d %d %d wakes up", &month, &day, &hour, &minute)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("wakes", month, day, hour, minute, guard)
			wakeupTime = minute
			isSleeping = false
			guardSleepMinutes[currGuard] += wakeupTime - fellAsleepTime + 1
			for i := fellAsleepTime; i < wakeupTime; i++ {
				gm := GuardMinute{currGuard, i}
				guardSleepMinuteFreqs[gm]++
			}
			continue
		}

	}

	fmt.Println(guardSleepMinutes)
	sortedBySleep := []string{}
	for guardId, sleepTime := range guardSleepMinutes {
		sortedBySleep = append(sortedBySleep, fmt.Sprintf("%d - %d guard\n", sleepTime, guardId))
	}
	sort.Strings(sortedBySleep)
	fmt.Println(sortedBySleep)

	// # 641
	maxgm := GuardMinute{}
	maxfreq := 0
	for gm, freq := range guardSleepMinuteFreqs {
		if freq > maxfreq {
			maxfreq = freq
			maxgm = gm
		}
		fmt.Println("BOB", freq, gm.Guard, gm.Minute)
	}

	fmt.Println(guardSleepMinuteFreqs)
	fmt.Println("most sleep at minute", maxgm.Minute, "is guard", maxgm.Guard)
	fmt.Println("who slept for", maxfreq, "minutes at that minute")
	fmt.Println("nas is", maxgm.Guard*maxgm.Minute)

}

var _ = fmt.Println
var _ = strconv.Atoi
var _ = log.Fatal
var _ = bufio.NewScanner // (os.Stdin) -> scanner.Scan(), scanner.Text()
var _ = os.Stdin
var _ = strings.Split    // "str str" -> []string{"str", "str"}
var _ = ioutil.ReadFile  // ("input.txt") -> (buf, err)
var _ = csv.NewReader    // (os.Stdin)
var _ = disjointset.Make // (10) -> ds. ds.Union(a,b), ds.Connected(a,b), ds.Count

func atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
