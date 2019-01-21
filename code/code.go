package code

import "fmt"

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
