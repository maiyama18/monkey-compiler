package code

import (
	"encoding/binary"
	"fmt"
)

type Instructions []byte

type Opcode byte

type Definition struct {
	Name          string
	OperandWidths []int
}

const (
	OpConstant Opcode = iota
)

var definitions = map[Opcode]*Definition {
	OpConstant: &Definition{Name: "OpConstant", OperandWidths: []int{2}},
}

func LookUp(opcode Opcode) (*Definition, error) {
	def, ok := definitions[opcode]
	if !ok {
		return nil, fmt.Errorf("opcode %d undefined", opcode)
	}

	return def, nil
}

func Make(opcode Opcode, operands ...int) []byte {
	def, err := LookUp(opcode)
	if err != nil {
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
