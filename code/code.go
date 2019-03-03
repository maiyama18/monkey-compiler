package code

import (
	"encoding/binary"
	"fmt"
)

// Instructions is byte array representing code
type Instructions []byte

// Opcode is a byte corresponding a instruction
type Opcode byte

const (
	// OpConstant register literal in monkey code
	OpConstant Opcode = iota
)

// Definition defines Opcode
type Definition struct {
	Name          string
	OperandWidths []int
}

var definitions = map[Opcode]*Definition{
	OpConstant: {"OpConstant", []int{2}},
}

// Lookup returns definition of passed opcode
func Lookup(opcode byte) (*Definition, error) {
	def, ok := definitions[Opcode(opcode)]
	if !ok {
		return nil, fmt.Errorf("Opcode %d is not defined", opcode)
	}
	return def, nil
}

// Make makes instruction byte array from opcode and operands
func Make(opcode Opcode, operands ...int) []byte {
	def, ok := definitions[opcode]
	if !ok {
		return []byte{}
	}

	instructionLen := 1
	for _, w := range def.OperandWidths {
		instructionLen += w
	}

	instruction := make([]byte, instructionLen)
	instruction[0] = byte(opcode)

	offset := 1
	for i, operand := range operands {
		width := def.OperandWidths[i]
		switch width {
		case 2:
			binary.BigEndian.PutUint16(instruction[offset:], uint16(operand))
		}
		offset += width
	}

	return instruction
}
