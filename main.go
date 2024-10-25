package main

import (
	"fmt"
	"os"
	"time"

	"github.com/selvamtech08/author-tex-xml/lexer"
	"github.com/selvamtech08/author-tex-xml/parser"
)

func main() {

	start := time.Now()
	defer func() {
		// just for debug purpose
		fmt.Println("elapsed:", time.Since(start))
	}()

	src, err := os.ReadFile("./example/02.tex")
	if err != nil {
		fmt.Println(err)
		return
	}

	lex := lexer.NewLexer(string(src))
	lex.Generate()
	par := parser.NewParser(lex)
	par.Parse()

}
