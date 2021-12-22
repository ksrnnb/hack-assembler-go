package main

import (
	"bufio"
	"fmt"
	"io"
)

type Parser struct {
	scanner        *bufio.Scanner
	currentCommand string
	isDone         bool
}

func NewParser(r io.Reader) *Parser {
	scanner := bufio.NewScanner(r)
	return &Parser{scanner: scanner, isDone: false}
}

func (p Parser) HasMoreCommands() bool {
	return !p.isDone
}

func (p *Parser) Advance() {
	p.scanner.Scan()

	command := p.scanner.Text()

	fmt.Println(command)
	// TODO: commandの整形
}
