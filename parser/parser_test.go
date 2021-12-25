package parser_test

import (
	"bytes"
	"testing"

	. "github.com/ksrnnb/hack-assembler-go/parser"
)

func TestHasMoreCommands(t *testing.T) {
	reader := bytes.NewReader([]byte("@2"))
	parser := NewParser(reader)

	if !parser.HasMoreCommands() {
		t.Error("parser should not be done")
	}

	parser.Advance()
	parser.Advance()

	if parser.HasMoreCommands() {
		t.Error("parser should be done")
	}
}

func TestAdvance(t *testing.T) {
	reader := bytes.NewReader([]byte("@2\nD=M"))
	parser := NewParser(reader)

	parser.Advance()

	if parser.CurrentCommand() != "@2" {
		t.Error("1st current command is invalid")
	}

	parser.Advance()

	if parser.CurrentCommand() != "D=M" {
		t.Error("2nd current command is invalid")
	}
}

func TestAdvanceWithComment(t *testing.T) {
	reader := bytes.NewReader([]byte("// test comment / this line will be ignored @3\n@2 // test comment after \nD=M"))
	parser := NewParser(reader)

	parser.Advance()

	if parser.CurrentCommand() != "@2" {
		t.Error("1st current command is invalid")
	}

	parser.Advance()

	if parser.CurrentCommand() != "D=M" {
		t.Error("2nd current command is invalid")
	}
}

func TestAdvanceWithEmptyLine(t *testing.T) {
	reader := bytes.NewReader([]byte("\n\n\n\n@2\n\n\n\nD=M"))
	parser := NewParser(reader)

	parser.Advance()

	if parser.CurrentCommand() != "@2" {
		t.Error("1st current command is invalid")
	}

	parser.Advance()

	if parser.CurrentCommand() != "D=M" {
		t.Error("2nd current command is invalid")
	}
}

func TestAdvanceWithSpace(t *testing.T) {
	reader := bytes.NewReader([]byte("    @   2   \n   D    =   M    "))
	parser := NewParser(reader)

	parser.Advance()

	if parser.CurrentCommand() != "@2" {
		t.Error("1st current command is invalid")
	}

	parser.Advance()

	if parser.CurrentCommand() != "D=M" {
		t.Error("2nd current command is invalid")
	}
}

func TestDest(t *testing.T) {
	reader := bytes.NewReader([]byte("@2\nM=M+1"))
	parser := NewParser(reader)

	parser.Advance()

	if _, err := parser.Dest(); err == nil {
		t.Error("A command cannot get dest")
	}

	parser.Advance()

	dest, err := parser.Dest()
	if err != nil {
		t.Error("C command should be executable Dest()")
	}

	if dest != "M" {
		t.Error("Dest is invalid")
	}
}

func TestSymbol(t *testing.T) {
	reader := bytes.NewReader([]byte("@2\nD=M\n(xxx)"))
	parser := NewParser(reader)

	parser.Advance()
	symbol, err := parser.Symbol()

	if err != nil {
		t.Error("symbol error")
	}

	if symbol != "2" {
		t.Error("A Command parse error")
	}

	parser.Advance()
	_, err = parser.Symbol()

	if err == nil {
		t.Error("C command should not be called")
	}

	parser.Advance()
	lSymbol, err := parser.Symbol()

	if err != nil {
		t.Error("symbol error")
	}

	if lSymbol != "xxx" {
		t.Error("L Command parse error")
	}
}

func TestComp(t *testing.T) {
	reader := bytes.NewReader([]byte("@2\nM=M+1"))
	parser := NewParser(reader)

	parser.Advance()

	if _, err := parser.Comp(); err == nil {
		t.Error("A command cannot get comp")
	}

	parser.Advance()

	dest, err := parser.Comp()
	if err != nil {
		t.Error("C command should be executable Comp()")
	}

	if dest != "M+1" {
		t.Error("Comp is invalid")
	}
}

func TestJump(t *testing.T) {
	reader := bytes.NewReader([]byte("@2\nM=M+1\nD;JMP"))
	parser := NewParser(reader)

	parser.Advance()

	if _, err := parser.Jump(); err == nil {
		t.Error("A command cannot get comp")
	}

	parser.Advance()
	jump, err := parser.Jump()
	if err != nil {
		t.Error("C command should be executable Jump()")
	}

	if jump != "" {
		t.Error("jump should be empty string for 'M=M+1'")
	}

	parser.Advance()
	jump, err = parser.Jump()

	if err != nil {
		t.Error("C command should be executable Jump()")
	}

	if jump != "JMP" {
		t.Error("Comp is invalid")
	}
}
