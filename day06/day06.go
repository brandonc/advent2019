package day06

import (
	"fmt"
	"os"
	"strings"

	"github.com/brandonc/advent2019/tools"
)

func distancefromCOM(orbits map[string]string, thing string) int {
	if thing == "COM" {
		return 0
	}

	return 1 + distancefromCOM(orbits, orbits[thing])
}

func heirarchy(orbits map[string]string, start string) []string {
	result := make([]string, 0)
	parent := orbits[start]
	for {
		result = append(result, parent)
		if parent == "COM" {
			break
		}
		parent = orbits[parent]
	}

	return result
}

func Run() {
	reader, err := tools.Readlines(os.Stdin)

	if err != nil {
		panic(err)
	}

	orbitlist := map[string]string{}

	for line := range reader {
		tuple := strings.Split(line, ")")
		orbitlist[tuple[1]] = tuple[0]
	}

	total := 0
	for id := range orbitlist {
		total += distancefromCOM(orbitlist, id)
	}

	fmt.Println("Part 1 answer:", total)

	myheirarchy := heirarchy(orbitlist, "YOU")
	santasheirarchy := heirarchy(orbitlist, "SAN")

	// Finds the orbital parent that SANTA and YOU have in common
	transfers := 0
	santatransfers := 0
	for myparent := range myheirarchy {
		santatransfers = 0
		for santasparent := range santasheirarchy {
			if santasheirarchy[santasparent] == myheirarchy[myparent] {
				fmt.Println("Part 2 answer:", transfers+santatransfers)
				os.Exit(0)
			}
			santatransfers += 1
		}
		transfers += 1
	}
}
