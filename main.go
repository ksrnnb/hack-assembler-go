package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/ksrnnb/hack-assembler-go/code"
	"github.com/ksrnnb/hack-assembler-go/parser"
	"github.com/ksrnnb/hack-assembler-go/symbol"
)

func main() {
	inputFile, err := os.Open("max.asm")

	if err != nil {
		panic(err)
	}

	defer inputFile.Close()
	err = symbol.RegisterSymbolTable(inputFile)

	if err != nil {
		panic(err)
	}

	// ファイルの読み込み位置を先頭位置に戻す
	_, err = inputFile.Seek(0, 0)
	if err != nil {
		panic(err)
	}

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
			fmt.Printf("error while parsing command type: %v\n", err)
			break
		}

		if cmdType == parser.ACommand {
			if err := doACommand(p, out); err != nil {
				fmt.Printf("error while doing A command: %v\n", err)
				break
			}
			continue
		}

		if cmdType == parser.CCommand {
			if err := doCCommand(p, out); err != nil {
				fmt.Printf("error while doing C command: %v\n", err)
				break
			}
			continue
		}
	}
}

func doACommand(parser *parser.Parser, out io.Writer) error {
	s, err := parser.Symbol()

	if err != nil {
		return err
	}

	intSymbol, err := strconv.Atoi(s)

	if err == nil {
		fmt.Fprintf(out, "0%015b\n", intSymbol)
		return nil
	}

	if symbol.Contains(s) {
		address, err := symbol.GetAddress(s)

		if err != nil {
			return err
		}

		fmt.Fprintf(out, "0%015b\n", address)
	} else {
		symbol.AddRAMEntry(s)
	}

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
