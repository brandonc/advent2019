package day05

import (
	"fmt"
	"github.com/brandonc/advent2019/tools"
	"os"
	"strconv"
	"strings"
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

func execute(codes []int, pos int) {
	input := 5
	instruction := expandInstruction(codes[pos])

	if instruction[3] == '9' && instruction[4] == '9' {
		return
	}

	opcode := instruction[4]

	var commandLength int
	switch opcode {
	case '1':
		fmt.Println("ADD", pos)
		commandLength = 4
		p2mode := instruction[1]
		p1mode := instruction[2]

		codes[codes[pos+3]] = param(codes, p1mode, codes[pos+1]) + param(codes, p2mode, codes[pos+2])
		break
	case '2':
		fmt.Println("MUL", pos)
		commandLength = 4
		p2mode := instruction[1]
		p1mode := instruction[2]

		codes[codes[pos+3]] = param(codes, p1mode, codes[pos+1]) * param(codes, p2mode, codes[pos+2])
		break
	case '3':
		fmt.Println("INP", pos)
		commandLength = 2
		codes[codes[pos+1]] = input
		break
	case '4':
		fmt.Println("OUT", pos)
		p1mode := instruction[2]
		p1 := param(codes, p1mode, codes[pos+1])

		commandLength = 2
		fmt.Println("Output", p1)
		break
	case '5':
		p2mode := instruction[1]
		p1mode := instruction[2]

		p1 := param(codes, p1mode, codes[pos+1])
		p2 := param(codes, p2mode, codes[pos+2])

		fmt.Println("JPT", pos)
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

		fmt.Println("JPF", pos)
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

		fmt.Println("LES", pos)
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

		fmt.Println("EQL", pos)
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

	execute(codes, pos+commandLength)
}

func Run() {
	input := tools.ReadInts(os.Stdin)

	codes := make([]int, len(input))
	copy(codes, input)

	execute(codes, 0)
}
