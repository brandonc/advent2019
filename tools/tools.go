package tools

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

// Readlines is an iterator that returns one line of a file at a time.
func Readlines(f *os.File) (<-chan string, error) {
	scanner := bufio.NewScanner(f)
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	chnl := make(chan string)
	go func() {
		for scanner.Scan() {
			chnl <- scanner.Text()
		}
		close(chnl)
	}()

	return chnl, nil
}

func ReadInts(f *os.File) []int {
	ScanCSV := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		commaidx := bytes.IndexByte(data, ',')
		if commaidx > 0 {
			// we need to return the next position
			buffer := data[:commaidx]
			return commaidx + 1, bytes.TrimSpace(buffer), nil
		}

		// if we are at the end of the string, just return the entire buffer
		if atEOF {
			// but only do that when there is some data. If not, this might mean
			// that we've reached the end of our input CSV string
			if len(data) > 0 {
				return len(data), bytes.TrimSpace(data), nil
			}
		}

		// when 0, nil, nil is returned, this is a signal to the interface to read
		// more data in from the input reader. In this case, this input is our
		// string reader and this pretty much will never occur.
		return 0, nil, nil
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(ScanCSV)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, toInt(scanner.Text()))
	}
	return numbers
}

func toInt(s string) int {
	result, err := strconv.Atoi(s)
	check(err)
	return result
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
