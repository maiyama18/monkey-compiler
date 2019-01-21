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
