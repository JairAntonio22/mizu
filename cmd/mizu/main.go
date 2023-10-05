package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JairAntonio22/pkg/lexer"
)

func main() {
	file, err := os.Open("examples/hello.mz")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	l := lexer.NewLexer(reader)

	for {
		token := l.Read()
		fmt.Println(token)
		break
	}
}
