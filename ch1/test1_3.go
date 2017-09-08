package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	s, sep := "", ""

	start := time.Now()

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	time.Sleep(1000)

	end := time.Now()

	result := end.Sub(start).Nanoseconds() / 1000000

	fmt.Println(result)

	fmt.Println(s)
}
