package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/JairAntonio22/pkg/mizu"
	"github.com/JairAntonio22/pkg/mizu/tokens"
)

func main() {
	file, err := os.Open("examples/primes.mz")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	lexer := mizu.NewLexer(reader)
	token := lexer.Read()

	start := time.Now()

	for token != tokens.Eof {
		if token == tokens.Eol {
			fmt.Println()
		} else {
			fmt.Printf("%v ", token)
		}

		token = lexer.Read()
	}

	fmt.Printf("Time taken: %v\n", time.Since(start))
}
