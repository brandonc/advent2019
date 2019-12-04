package day04

import (
	"fmt"
	"github.com/brandonc/advent2019/tools"
	"os"
	"strconv"
	"strings"
)

func neverDecreases(n string) bool {
	for pos := 0; pos < len(n)-1; pos++ {
		if n[pos] > n[pos+1] {
			return false
		}
	}
	return true
}

func hasDouble(n string) bool {
	groups := map[byte]int{}
	i := 0
	for i = 0; i < len(n); i++ {
		b := n[i]
		if _, ok := groups[b]; ok {
			groups[b] += 1
		} else {
			groups[b] = 1
		}
	}
	for _, value := range groups {
		if value == 2 {
			return true
		}
	}
	return false
}

func hasRepeated(n string) bool {
	for pos := 0; pos < len(n)-1; pos++ {
		if n[pos] == n[pos+1] {
			return true
		}
	}
	return false
}

func Run() {
	reader, err := tools.Readlines(os.Stdin)

	if err != nil {
		panic(err)
	}

	input := <-reader
	inputa := strings.Split(input, "-")
	first := tools.ToInt(inputa[0])
	last := tools.ToInt(inputa[1])
	candidatesPart1 := 0
	candidatesPart2 := 0

	for current := first; current <= last; current++ {
		currents := strconv.Itoa(current)

		if neverDecreases(currents) && hasDouble(currents) {
			candidatesPart2 += 1
			candidatesPart1 += 1
		} else if neverDecreases(currents) && hasRepeated(currents) {
			candidatesPart1 += 1
		}
	}

	fmt.Printf("Part 1 answer: %d\n", candidatesPart1)
	fmt.Printf("Part 2 answer: %d\n", candidatesPart2)
}
