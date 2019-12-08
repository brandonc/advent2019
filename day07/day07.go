package day07

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brandonc/advent2019/tools"
)

func expandInstruction(code int) string {
	var instruction strings.Builder
	instructionRaw := strconv.Itoa(code)

	for i := len(instructionRaw); i < 5; i++ {
		instruction.WriteString("0")
	}

	instruction.WriteString(instructionRaw)
	return instruction.String()
}

func param(codes []int, mode byte, v int) int {
	if mode == '0' {
		return codes[v]
	} else {
		return v
	}
}

func execute(input []int, output func(int), codes []int, pos int) []int {
	instruction := expandInstruction(codes[pos])

	if instruction[3] == '9' && instruction[4] == '9' {
		return input
	}

	opcode := instruction[4]

	var commandLength int
	switch opcode {
	case '1':
		commandLength = 4
		p2mode := instruction[1]
		p1mode := instruction[2]

		codes[codes[pos+3]] = param(codes, p1mode, codes[pos+1]) + param(codes, p2mode, codes[pos+2])
		break
	case '2':
		commandLength = 4
		p2mode := instruction[1]
		p1mode := instruction[2]

		codes[codes[pos+3]] = param(codes, p1mode, codes[pos+1]) * param(codes, p2mode, codes[pos+2])
		break
	case '3':
		commandLength = 2
		nextInput := input[0]
		input = input[1:len(input)]
		codes[codes[pos+1]] = nextInput
		break
	case '4':
		p1mode := instruction[2]
		p1 := param(codes, p1mode, codes[pos+1])

		commandLength = 2
		output(p1)
		break
	case '5':
		p2mode := instruction[1]
		p1mode := instruction[2]

		p1 := param(codes, p1mode, codes[pos+1])
		p2 := param(codes, p2mode, codes[pos+2])

		commandLength = 3
		if p1 != 0 {
			commandLength = 0
			pos = p2
		}
		break
	case '6':
		p2mode := instruction[1]
		p1mode := instruction[2]

		p1 := param(codes, p1mode, codes[pos+1])
		p2 := param(codes, p2mode, codes[pos+2])

		commandLength = 3

		if p1 == 0 {
			commandLength = 0
			pos = p2
		}
	case '7':
		p2mode := instruction[1]
		p1mode := instruction[2]

		p1 := param(codes, p1mode, codes[pos+1])
		p2 := param(codes, p2mode, codes[pos+2])

		commandLength = 4

		if p1 < p2 {
			codes[codes[pos+3]] = 1
		} else {
			codes[codes[pos+3]] = 0
		}
		break
	case '8':
		p2mode := instruction[1]
		p1mode := instruction[2]

		p1 := param(codes, p1mode, codes[pos+1])
		p2 := param(codes, p2mode, codes[pos+2])

		commandLength = 4

		if p1 == p2 {
			codes[codes[pos+3]] = 1
		} else {
			codes[codes[pos+3]] = 0
		}
		break

	default:
		panic("Something is wrong")
	}

	return execute(input, output, codes, pos+commandLength)
}

func permutationsEmit(f func([]int), k int, array []int) {
	if k == 1 {
		c := make([]int, len(array))
		copy(c, array)
		f(c)
	} else {
		permutationsEmit(f, k-1, array)

		for i := 0; i < k-1; i++ {
			if k%2 == 0 {
				t := array[i]
				array[i] = array[k-1]
				array[k-1] = t
			} else {
				t := array[0]
				array[0] = array[k-1]
				array[k-1] = t
			}
			permutationsEmit(f, k-1, array)
		}
	}
}

func permutations(array []int) [][]int {
	result := make([][]int, 0)
	collectPermutations := func(p []int) {
		result = append(result, p)
	}
	permutationsEmit(collectPermutations, len(array), array)
	return result
}

func part1(program []int) int {
	phasesPermutations := permutations([]int{0, 1, 2, 3, 4})

	max := 0
	for i := range phasesPermutations {
		phases := phasesPermutations[i]

		// Input 0 at first position
		phases = append(phases, 0)
		copy(phases[2:], phases[1:])
		phases[1] = 0

		output := 0
		for {
			codes := make([]int, len(program))
			copy(codes, program)

			phases = execute(phases, func(o int) { output = o }, codes, 0)

			if len(phases) == 0 {
				break
			}

			phases = append(phases, 0)
			copy(phases[2:], phases[1:])
			phases[1] = output
		}

		if output > max {
			max = output
		}
	}

	return max
}

func part2(program []int) int {
	phasesPermutations := permutations([]int{5, 6, 7, 8, 9})

	max := 0
	for i := range phasesPermutations {
		phases := phasesPermutations[i]

		codesA := make([]int, len(program))
		copy(codesA, program)

		codesB := make([]int, len(program))
		copy(codesB, program)

		codesC := make([]int, len(program))
		copy(codesC, program)

		codesD := make([]int, len(program))
		copy(codesD, program)

		codesE := make([]int, len(program))
		copy(codesE, program)

		allCodes := [][]int{codesA, codesB, codesC, codesD, codesE}
		pointer := 0

		var handleOutput func(int)

		lastOutput := 0
		iter := 0
		handleOutput = func(output int) {
			if iter == 11 {
				os.Exit(0)
			}
			iter += 1
			pointer += 1
			pointer %= 5

			if pointer == 0 {
				lastOutput = output
			}

			fmt.Printf("%d -> amp %s\n", output, []string{"A", "B", "C", "D", "E"}[pointer])

			execute([]int{phases[pointer], output}, handleOutput, allCodes[pointer], 0)
		}

		fmt.Printf("0 -> amp A\n")
		execute([]int{0, phases[0]}, handleOutput, allCodes[0], 0)

		if lastOutput > max {
			max = lastOutput
		}
	}

	return max
}

func Run() {
	program := tools.ReadInts(os.Stdin)

	fmt.Println("Part 1 answer", part1(program))
	fmt.Println("Part 2 answer", part2(program))
}
