package symbol

import (
	"errors"
	"io"

	"github.com/ksrnnb/hack-assembler-go/parser"
)

var st map[string]int
var RAMAddress = 16

func init() {
	st = map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"R0":     0,
		"R1":     1,
		"R2":     2,
		"R3":     3,
		"R4":     4,
		"R5":     5,
		"R6":     6,
		"R7":     7,
		"R8":     8,
		"R9":     9,
		"R10":    10,
		"R11":    11,
		"R12":    12,
		"R13":    13,
		"R14":    14,
		"R15":    15,
		"SCREEN": 16384,
		"KBD":    245576,
	}
}

func AddEntry(symbol string, address int) {
	st[symbol] = address
}

func AddRAMEntry(symbol string) {
	st[symbol] = RAMAddress
	RAMAddress++
}

func Contains(symbol string) bool {
	_, ok := st[symbol]

	return ok
}

func GetAddress(symbol string) (address int, err error) {
	address, ok := st[symbol]

	if !ok {
		return 0, errors.New("symbol is not registered")
	}

	return address, nil
}

func RegisterSymbolTable(input io.Reader) error {
	p := parser.NewParser(input)

	ROMAddress := 0
	for {
		p.Advance()

		if !p.HasMoreCommands() {
			break
		}

		cmdType, err := p.CommandType()

		if err != nil {
			return err
		}

		if cmdType != parser.LCommand {
			ROMAddress++
			continue
		}

		symbol, err := p.Symbol()

		if err != nil {
			return err
		}

		st[symbol] = ROMAddress + 1
	}

	return nil
}
