package advent

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func day4() error {
	i, err := os.Open("day4.input")
	if err != nil {
		return err
	}
	defer i.Close()

	logs := []string{}

	scanner := bufio.NewScanner(i)
	for scanner.Scan() {
		logs = append(logs, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	sort.Strings(logs)

	var guardID = regexp.MustCompile(`.*#([0-9]+).*$`)
	var time = regexp.MustCompile(`^\[1518-([0-9]{2}-[0-9]{2}) [0-9]{2}:([0-9]{2})\].*$`)

	sleepSchedule := map[string][]int{}
	sleepyGuards := map[string]int{}
	var guard string
	var sleep int
	for _, line := range logs {
		t := time.FindStringSubmatch(line)
		minute, _ := strconv.Atoi(t[2])

		if guardID.FindString(line) != "" {
			guard = guardID.FindStringSubmatch(line)[1]
		} else if strings.Contains(line, "falls asleep") {
			sleep = minute
		} else if strings.Contains(line, "wakes up") {
			minute--

			timeAsleep := minute - sleep
			sleepyGuards[guard] += timeAsleep

			if _, exist := sleepSchedule[guard]; !exist {
				sleepSchedule[guard] = make([]int, 60)
			}

			for i := sleep; i <= minute; i++ {
				sleepSchedule[guard][i]++
			}
		}
	}

	type sleepy struct {
		Guard   string
		Minutes int
	}

	var ss []sleepy
	for k, v := range sleepyGuards {
		ss = append(ss, sleepy{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Minutes > ss[j].Minutes
	})

	//fmt.Printf("%+v\n", ss[0].Guard)
	//fmt.Printf("%+v\n", sleepSchedule[ss[0].Guard])

	m := 0
	minute := 0
	for min, times := range sleepSchedule[ss[0].Guard] {
		if min == 0 || times > m {
			m = times
			minute = min
		}
	}
	//fmt.Printf("%d = %+v\n", minute, sleepSchedule[ss[0].Guard][minute])
	gNum, _ := strconv.Atoi(ss[0].Guard)
	fmt.Printf("Day 1 part 1: %d\n", gNum*minute)

	m = 0
	minute = 0
	guard = ""
	for g := range sleepSchedule {
		//fmt.Printf("%s - %+v\n", g, sleepSchedule[g])
		for min, times := range sleepSchedule[g] {
			if times > m {
				m = times
				minute = min
				guard = g
			}
		}
	}

	gNum, _ = strconv.Atoi(guard)
	fmt.Printf("Day 1 part 2: %d\n", gNum*minute)

	return nil
}
