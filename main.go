package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ksrnnb/hack-assembler-go/code"
	"github.com/ksrnnb/hack-assembler-go/parser"
)

func main() {
	inputFile, err := os.Open("test.asm")

	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	p := parser.NewParser(inputFile)

	out, err := os.Create("test.hack")

	if err != nil {
		panic(err)
	}

	for {
		p.Advance()

		if !p.HasMoreCommands() {
			break
		}

		cmdType, err := p.CommandType()

		if err != nil {
			fmt.Printf("error while parsing command type: %v", err)
			break
		}

		if cmdType == parser.ACommand {
			if err := doACommand(p, out); err != nil {
				fmt.Printf("error while doing A command: %v", err)
				break
			}
			continue
		}

		if cmdType == parser.CCommand {
			if err := doCCommand(p, out); err != nil {
				fmt.Printf("error while doing C command: %v", err)
				break
			}
			continue
		}
	}
}

func doACommand(parser *parser.Parser, out io.Writer) error {
	symbol, err := parser.Symbol()

	if err != nil {
		return err
	}

	intSymbol, err := strconv.Atoi(symbol)

	if err != nil {
		return err
	}

	fmt.Fprintf(out, "0%015b\n", intSymbol)
	return nil
}

func doCCommand(parser *parser.Parser, out io.Writer) error {
	dest, err := parser.Dest()

	if err != nil {
		return err
	}

	comp, err := parser.Comp()

	if err != nil {
		return err
	}

	jump, err := parser.Jump()

	if err != nil {
		return err
	}

	destBinary, err := code.Dest(dest)

	if err != nil {
		return err
	}

	compBinary, err := code.Comp(comp)

	if err != nil {
		return err
	}

	jumpBinary, err := code.Jump(jump)

	if err != nil {
		return err
	}

	fmt.Fprintf(out, "111%07b%03b%03b\n", compBinary, destBinary, jumpBinary)
	return nil
}
