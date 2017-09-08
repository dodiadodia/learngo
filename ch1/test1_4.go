package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)

	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "dup2 must input file name\n")
		return
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}

			countLine(f, counts)
			f.Close()
		}
	}

	for line, sets := range counts {
		if len(sets) > 1 {
			fmt.Printf("%s are existing at file %s\n", line, strings.Join(sets, " "))
		}
	}
}

func countLine(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
}
