package code

import (
	"testing"
)

func TestMake(t *testing.T) {
	testCases := []struct {
		desc     string
		opcode   Opcode
		operands []int
		expected []byte
	}{
		{
			"opconstant", OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254},
		},
		{
			"oppop", OpPop, []int{}, []byte{byte(OpPop)},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			instruction := Make(tc.opcode, tc.operands...)

			if len(tc.expected) != len(instruction) {
				t.Errorf("the length of instruction wrong. want=%d, got=%d", len(tc.expected), len(instruction))
			}
			for i, b := range tc.expected {
				if b != instruction[i] {
					t.Errorf("the instruction wrong. want=%s, got=%s", tc.expected, instruction)
					break
				}
			}
		})
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := []Instructions{
		Make(OpConstant, 1),
		Make(OpConstant, 2),
		Make(OpAdd),
		Make(OpConstant, 65535),
	}

	expected := `0000 OpConstant 1
0003 OpConstant 2
0006 OpAdd
0007 OpConstant 65535
`

	concatenated := concatInstructions(instructions)

	if expected != concatenated.String() {
		t.Errorf("Instructions.String() wrong.\nwant=%s\ngot=%s", expected, concatenated)
	}
}

func concatInstructions(instructions []Instructions) Instructions {
	out := Instructions{}
	for _, ins := range instructions {
		out = append(out, ins...)
	}
	return out
}

func TestReadOperands(t *testing.T) {
	tests := []struct {
		name      string
		opcode    Opcode
		operands  []int
		bytesRead int
	}{
		{
			"opconstant", OpConstant, []int{65535}, 2,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			instruction := Make(tc.opcode, tc.operands...)

			def, err := Lookup(byte(tc.opcode))
			if err != nil {
				t.Errorf("lookup error: %s", err.Error())
			}

			operandsRead, n := ReadOperands(def, instruction[1:])
			if n != tc.bytesRead {
				t.Errorf("number of bytes of operand wrong. want=%d, got=%d", tc.bytesRead, n)
			}
			if len(operandsRead) != len(tc.operands) {
				t.Errorf("operand read wrong. want=%v, got=%v", tc.operands, operandsRead)
			}
			for i, o := range operandsRead {
				if o != tc.operands[i] {
					t.Errorf("operand read wrong. want=%v, got=%v", tc.operands, operandsRead)
					break
				}
			}
		})
	}
}
