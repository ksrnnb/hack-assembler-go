package parser

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/ksrnnb/hack-assembler-go/utils"
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

var jumpMnemonic utils.StringSlice = []string{"JGT", "JEQ", "JGE", "JLT", "JNE", "JLE", "JMP"}

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

	if len(strings.Split(p.currentCommand, "=")) == 2 ||
		len(strings.Split(p.currentCommand, ";")) == 2 {
		return CCommand, nil
	}

	return 0, errors.New("error while getting command type")
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

// destを返す
// dest=comp;jump
func (p Parser) Dest() (dest string, err error) {
	cmdType, err := p.CommandType()

	if err != nil {
		return "", err
	}

	if cmdType != CCommand {
		return "", errors.New("Dest: command type should be c command")
	}

	cmds := strings.Split(p.currentCommand, "=")

	if len(cmds) == 2 {
		return cmds[0], nil
	}

	return "null", nil
}

// compを返す
// dest=comp;jump
func (p Parser) Comp() (comp string, err error) {
	cmdType, err := p.CommandType()

	if err != nil {
		return "", err
	}

	if cmdType != CCommand {
		return "", errors.New("Comp: command type should be c command")
	}

	cmds := strings.Split(p.currentCommand, "=")
	if len(cmds) == 2 {
		// dest=comp;jump
		cmds = strings.Split(cmds[1], ";")

		if len(cmds) == 2 {
			// comp, jump
			return cmds[1], nil
		}

		// dest, jump
		return cmds[0], nil
	}

	// destが無い場合 => comp;jump
	cmds = strings.Split(p.currentCommand, ";")

	if len(cmds) == 2 {
		// comp, jump
		return cmds[0], nil
	}

	return "", errors.New("Comp: c command should include '=' or ';'")
}

// jumpを返す
// dest=comp;jump
func (p Parser) Jump() (comp string, err error) {
	cmdType, err := p.CommandType()

	if err != nil {
		return "", err
	}

	if cmdType != CCommand {
		return "", errors.New("Jump: command type should be c command")
	}

	cmds := strings.Split(p.currentCommand, ";")

	if len(cmds) < 2 {
		return "null", nil
	}

	if !jumpMnemonic.Contains(cmds[1]) {
		return "", errors.New("Jump: mnemonic is invalid")
	}

	return cmds[1], nil
}

// 改行のみの場合にtrue
func isEmptyLine(command string) bool {
	return command == ""
}
