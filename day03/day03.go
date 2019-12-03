package day03

import (
	"bufio"
	"fmt"
	"github.com/brandonc/advent2019/tools"
	"math"
	"os"
	"strings"
)

type Point struct {
	X int
	Y int
}

func storeSteps(grid map[Point][]int, wireindex int, steps int, current *Point) {
	point := Point{X: current.X, Y: current.Y}

	if val, ok := grid[point]; ok {
		if val[wireindex] == 0 {
			val[wireindex] = steps
		}
	} else {
		grid[point] = make([]int, 2)
		grid[point][wireindex] = steps
	}
}

func Run() {
	line_scanner, err := tools.Readlines(os.Stdin)
	if err != nil {
		panic(err)
	}

	grid := map[Point][]int{}

	wirecount := 0
	for wire := range line_scanner {
		steps := 0
		follower := Point{X: 0, Y: 0}
		reader := strings.NewReader(wire)
		scanner := bufio.NewScanner(reader)
		scanner.Split(tools.ScanCSV)

		for scanner.Scan() {
			instruction := scanner.Text()
			direction := instruction[0]
			distance := tools.ToInt(instruction[1:])

			switch direction {
			case 'U':
				for x := 0; x < distance; x++ {
					steps++
					follower.Y += 1
					storeSteps(grid, wirecount, steps, &follower)
				}
				break
			case 'D':
				for x := 0; x < distance; x++ {
					steps++
					follower.Y -= 1
					storeSteps(grid, wirecount, steps, &follower)
				}
				break
			case 'R':
				for x := 0; x < distance; x++ {
					steps++
					follower.X += 1
					storeSteps(grid, wirecount, steps, &follower)
				}
				break
			case 'L':
				for x := 0; x < distance; x++ {
					steps++
					follower.X -= 1
					storeSteps(grid, wirecount, steps, &follower)
				}
				break
			}
		}
		wirecount += 1
	}

	closestDist := math.MaxInt32
	fewestSteps := math.MaxInt32

	for k := range grid {
		p := grid[k]
		dist := tools.Abs(k.X) + tools.Abs(k.Y)

		if p[0] > 0 && p[1] > 0 && dist < closestDist {
			closestDist = dist
		}

		if p[0] > 0 && p[1] > 0 && p[0]+p[1] < fewestSteps {
			fewestSteps = p[0] + p[1]
		}
	}

	fmt.Printf("Part 1 answer: %d\n", closestDist)
	fmt.Printf("Part 2 answer: %d\n", fewestSteps)
}
