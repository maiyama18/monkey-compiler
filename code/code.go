package code

import (
	"bytes"
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

func (ins Instructions) String() string {
	out := bytes.Buffer{}
	offset := 0

	for offset < len(ins) {
		opcodeOffset := offset
		opcode := Opcode(ins[offset])
		offset++

		def, err := LookUp(opcode)
		if err != nil {
			return ""
		}

		operands, read := ReadOperands(def, ins[offset:])
		offset += read

		fmt.Fprintf(&out, "%04d %s", opcodeOffset, def.Name)
		for _, operand := range operands {
			fmt.Fprintf(&out, " %d", operand)
		}
		fmt.Fprintf(&out, "\n")
	}

	return out.String()
}

func ReadOperands(def *Definition, ins Instructions) ([]int, int) {
	operands := make([]int, len(def.OperandWidths))
	offset := 0

	for i, width := range def.OperandWidths {
		switch width {
		case 2:
			operands[i] = int(ReadUint16(ins[offset:]))
		}

		offset += width
	}

	return operands, offset
}

func ReadUint16(b []byte) uint16 {
	return binary.BigEndian.Uint16(b)
}

