package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/JairAntonio22/pkg/scan"
	"github.com/JairAntonio22/pkg/scan/tokens"
)

func main() {
	file, err := os.Open("examples/primes.mz")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	scanner := scan.NewScanner(reader)
	token := scanner.Read()

	start := time.Now()

	for token != tokens.Eof {
		if token == tokens.Eol {
			fmt.Println()
		} else {
			fmt.Printf("%v ", token)
		}

		token = scanner.Read()
	}

	fmt.Printf("Time taken: %v\n", time.Since(start))
}
