package main

import (
	"fmt"
	"os"
	"strconv"
)

func module_fuel(weight int) (int, int) {
	var fuel int = weight/3 - 2

	var more_fuel = 0

	var t = fuel/3 - 2
	for {
		if t <= 0 {
			break
		}

		more_fuel = more_fuel + t
		t = t/3 - 2
	}
	return fuel, more_fuel
}

func main() {
	scanner, err := Readlines(os.Stdin)

	if err != nil {
		fmt.Println(err)
	}

	var part1_total int = 0
	var part2_total int = 0

	for line := range scanner {
		var (
			module    int
			fuel      int
			more_fuel int = 0
			err       error
		)

		module, err = strconv.Atoi(line)
		if err != nil {
			fmt.Print("Not a number!\n")
			break
		}
		fuel, more_fuel = module_fuel(module)

		part1_total = part1_total + fuel
		part2_total = part2_total + fuel + more_fuel
	}

	fmt.Printf("Part 1 Total: %d\n", part1_total)
	fmt.Printf("Part 2 Total: %d\n", part2_total)
}
