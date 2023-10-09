package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/JairAntonio22/pkg/mizu/scanner"
	"github.com/JairAntonio22/pkg/mizu/token"
)

func main() {
	file, err := os.Open("examples/primes.mz")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	s := scanner.New(reader)

	start := time.Now()

	for t := s.Read(); t != token.Eof; t = s.Read() {
		if t == token.Eol {
			fmt.Printf("\n")
		} else {
			fmt.Printf("%v ", t)
		}
	}

	fmt.Printf("Time taken: %v\n", time.Since(start))
}
