package main

import (
	"bufio"
	"os"
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
