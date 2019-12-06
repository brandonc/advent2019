package day06

import (
	"fmt"
	"os"
	"strings"

	"github.com/brandonc/advent2019/tools"
)

func distancefromcom(orbits map[string]string, thing string) int {
	if thing == "COM" {
		return 0
	}

	return 1 + distancefromcom(orbits, orbits[thing])

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
		total += distancefromcom(orbitlist, id)
	}

	fmt.Println("Part 1 answer:", total)
}
