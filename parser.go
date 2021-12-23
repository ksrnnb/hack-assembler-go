package main

import (
	"bufio"
	"io"
	"strings"
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
	cmdRemovedSpace := strings.ReplaceAll(command, " ", "")
	cmd := strings.Split(cmdRemovedSpace, "//")[0]

	if isEmptyLine(cmd) {
		p.Advance()
		return
	}

	p.currentCommand = cmd
}

func (p Parser) CurrentCommand() string {
	return p.currentCommand
}

// 改行のみの場合にtrue
func isEmptyLine(command string) bool {
	return command == ""
}
