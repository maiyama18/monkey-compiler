package code

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMake(t *testing.T) {
	testCases := []struct {
		opcode   Opcode
		operands []int
		expected []byte
	}{
		{
			opcode: OpConstant,
			operands: []int{65534},
			expected: []byte{byte(OpConstant), 255, 254},
		},
	}

	for _, tC := range testCases {
        instruction := Make(tC.opcode, tC.operands...)

        assert.Equal(t, instruction, tC.expected)
	}
}
