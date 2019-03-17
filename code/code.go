package code

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// Instructions is byte array representing code
type Instructions []byte

func (ins Instructions) String() string {
	var out bytes.Buffer
	offset := 0

	for offset < len(ins) {
		def, err := Lookup(ins[offset])
		if err != nil {
			_, _ = fmt.Fprintf(&out, "ERROR: %s\n", err)
			continue
		}

		operands, read := ReadOperands(def, ins[offset+1:])

		_, _ = fmt.Fprintf(&out, "%04d %s\n", offset, fmtInstruction(def, operands))

		offset += 1 + read
	}

	return out.String()
}

func fmtInstruction(def *Definition, operands []int) string {
	operandCount := len(def.OperandWidths)
	if len(operands) != operandCount {
		return fmt.Sprintf("ERROR: number of operands wrong: want=%d, got=%d", operandCount, len(operands))
	}

	switch operandCount {
	case 1:
		return fmt.Sprintf("%s %d", def.Name, operands[0])
	}

	return fmt.Sprintf("ERROR: unexpected number of operands for %s: %d", def.Name, operandCount)
}

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
		return nil, fmt.Errorf("opcode %d is not defined", opcode)
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

func ReadUint16(ins Instructions) uint16 {
	return binary.BigEndian.Uint16(ins)
}
