package main_test

import (
	"bytes"
	"testing"

	. "github.com/ksrnnb/hack-assembler-go"
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
