package code

import (
	"testing"
)

func TestMake(t *testing.T) {
	tests := []struct {
		desc     string
		opcode   Opcode
		operands []int
		expected []byte
	}{
		{
			desc:     "OpConstant",
			opcode:   OpConstant,
			operands: []int{65534},
			expected: []byte{byte(OpConstant), 255, 254},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			instruction := Make(tt.opcode, tt.operands...)

			if len(instruction) != len(tt.expected) {
				t.Fatalf("instruction has wrong length. expected=%q, actual=%q", tt.expected, instruction)
			}

			for i, e := range tt.expected {
				if instruction[i] != e {
					t.Fatalf("instruction has wrong byte at %d. expected=%d, actual=%d", i, e, instruction[i])
				}
			}
		})
	}
}

func TestInstructionsString(t *testing.T) {
	instructions := concatInstructions([]Instructions{
		Make(OpConstant, 1),
		Make(OpConstant, 2),
		Make(OpConstant, 65535),
	})

	expected := `0000 OpConstant 1
0003 OpConstant 2
0006 OpConstant 65535
`

	if instructions.String() != expected {
		t.Errorf("instructions wrongly formatted.\nexpected=%q\nactual=%q", expected, instructions.String())
	}
}

func concatInstructions(instructions []Instructions) Instructions {
	out := Instructions{}

	for _, ins := range instructions {
		out = append(out, ins...)
	}

	return out
}
