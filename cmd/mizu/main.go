package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JairAntonio22/pkg/lexer"
)

func main() {
	file, err := os.Open("examples/primes.mz")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	l := lexer.NewLexer(reader)

	for i := 0; i < 20; i++ {
		token := l.Read()
		fmt.Println(token)
	}
}
