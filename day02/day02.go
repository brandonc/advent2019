package day02

import (
	"fmt"
	"github.com/brandonc/advent2019/tools"
	"os"
)

func execute(codes []int, pos int) {
	instruction := codes[pos]

	if instruction == 99 {
		return
	}

	first := codes[codes[pos+1]]
	second := codes[codes[pos+2]]
	output_pos := codes[pos+3]

	switch instruction {
	case 1:
		codes[output_pos] = first + second
		break
	case 2:
		codes[output_pos] = first * second
	default:
		panic("Something went wrong")
	}

	execute(codes, pos+4)
}

func Run() {
	input := tools.ReadInts(os.Stdin)

	codes := make([]int, len(input))
	copy(codes, input)
	codes[1] = 12
	codes[2] = 2

	execute(codes, 0)

	fmt.Printf("Part 1 Position 0: %d\n", codes[0])

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			codes_part2 := make([]int, len(input))
			copy(codes_part2, input)
			codes_part2[1] = noun
			codes_part2[2] = verb

			execute(codes_part2, 0)

			if codes_part2[0] == 19690720 {
				fmt.Printf("Part 2 Answer: 100 * %d + %d = %d\n", noun, verb, 100*noun+verb)
				os.Exit(0)
			}
		}
	}
	fmt.Printf(("Part 2 not found!\n"))
}
