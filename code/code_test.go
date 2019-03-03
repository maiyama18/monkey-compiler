package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		opcode   Opcode
		operands []int
		expected []byte
	}{
		{
			"opconstant", OpConstant, []int{65534}, []byte{byte(OpConstant), 255, 254},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			instruction := Make(tc.opcode, tc.operands...)

			assert.Equal(t, len(tc.expected), len(instruction), "the length of instruction should be same")
			for i, b := range tc.expected {
				assert.Equal(t, b, instruction[i], "byte at index %d should be same", i)
			}
		})
	}
}
