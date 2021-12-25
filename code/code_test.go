package code_test

import (
	"testing"

	. "github.com/ksrnnb/hack-assembler-go/code"
)

func TestDest(t *testing.T) {
	correctMap := map[string]int{
		"null": 0b000,
		"M":    0b001,
		"D":    0b010,
		"MD":   0b011,
		"A":    0b100,
		"AM":   0b101,
		"AD":   0b110,
		"AMD":  0b111,
	}

	for mnemonic, value := range correctMap {
		v, err := Dest(mnemonic)

		if err != nil {
			t.Errorf("TestDest: error while testing Dest: %v\n", err)
		}

		if v != value {
			t.Errorf("TestDest: dest value is incorrect\n")
		}
	}
}

func TestJump(t *testing.T) {
	correctMap := map[string]int{
		"null": 0b000,
		"JGT":  0b001,
		"JEQ":  0b010,
		"JGE":  0b011,
		"JLT":  0b100,
		"JNE":  0b101,
		"JLE":  0b110,
		"JMP":  0b111,
	}

	for mnemonic, value := range correctMap {
		v, err := Jump(mnemonic)

		if err != nil {
			t.Errorf("TestJump: error while testing Jump: %v\n", err)
		}

		if v != value {
			t.Errorf("TestJump: jump value is incorrect\n")
		}
	}
}

func TestComp(t *testing.T) {
	correctMap := map[string]int{
		"0":   0b0101010,
		"1":   0b0111111,
		"-1":  0b0111010,
		"D":   0b0001100,
		"A":   0b0110000,
		"!D":  0b0001101,
		"-A":  0b0110011,
		"D+1": 0b0011111,
		"A+1": 0b0110111,
		"D-1": 0b0001110,
		"A-1": 0b0110010,
		"D+A": 0b0000010,
		"D-A": 0b0010011,
		"A-D": 0b0000111,
		"D&A": 0b0000000,
		"D|A": 0b0010101,
		"M":   0b1110000,
		"!M":  0b1110001,
		"-M":  0b1110011,
		"M+1": 0b1110111,
		"M-1": 0b1110010,
		"D+M": 0b1000010,
		"D-M": 0b1010011,
		"M-D": 0b1000111,
		"D&M": 0b1000000,
		"D|M": 0b1010101,
	}

	for mnemonic, value := range correctMap {
		v, err := Comp(mnemonic)

		if err != nil {
			t.Errorf("TestComp: error while testing Comp: %v\n", err)
		}

		if v != value {
			t.Errorf("TestComp: comp value is incorrect\n")
		}
	}
}
