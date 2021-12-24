package main

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

const (
	ACommand = iota // @3
	CCommand        // D=M
	LCommand        // (xxx)
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

func (p Parser) CommandType() (commandType int, err error) {
	// TODO: 正規表現にしたい
	if p.currentCommand[0] == '@' {
		return ACommand, nil
	}

	if p.currentCommand[0] == '(' &&
		p.currentCommand[len(p.currentCommand)-1] == ')' {
		return LCommand, nil
	}

	return CCommand, nil

	// return 0, errors.New("error while getting command type")
}

func (p Parser) Symbol() (symbol string, err error) {
	commandType, err := p.CommandType()

	if err != nil {
		return "", err
	}

	switch commandType {
	case CCommand:
		return "", errors.New("command type should be A command or L command")
	case ACommand:
		return p.currentCommand[1:], nil
	case LCommand:
		return p.currentCommand[1 : len(p.currentCommand)-1], nil
	default:
		return "", errors.New("command type is invalid")
	}
}

// 改行のみの場合にtrue
func isEmptyLine(command string) bool {
	return command == ""
}
