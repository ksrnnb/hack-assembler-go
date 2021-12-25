package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("test.asm")

	if err != nil {
		panic(err)
	}

	defer inputFile.Close()

	parser := NewParser(inputFile)

	out, err := os.Create("test.hack")

	if err != nil {
		panic(err)
	}

	for {
		parser.Advance()

		if !parser.HasMoreCommands() {
			break
		}

		cmdType, err := parser.CommandType()

		if err != nil {
			fmt.Printf("error while parsing command type: %v", err)
			break
		}

		if cmdType == ACommand {
			doACommand(parser, out)
			continue
		}

		if cmdType == CCommand {
			doCCommand(parser, out)
			continue
		}
	}
}

func doACommand(parser *Parser, out io.Writer) error {
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

func doCCommand(parser *Parser, out io.Writer) {

}
