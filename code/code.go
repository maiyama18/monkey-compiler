package code

import "fmt"

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
func Lookup(op byte) (*Definition, error) {
	def, ok := definitions[Opcode(op)]
	if !ok {
		return nil, fmt.Errorf("Opcode %d is not defined", op)
	}
	return def, nil
}
