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
	if !p.scanner.Scan() {
		p.isDone = true
		return
	}

	command := p.scanner.Text()

	// TODO: commandの整形
	fmt.Println(command)
	p.currentCommand = command
}

func (p Parser) CurrentCommand() string {
	return p.currentCommand
}
