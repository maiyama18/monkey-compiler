package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

			assert.Equal(t, len(tc.expected), len(instruction), "the length of instruction should be same")
			for i, b := range tc.expected {
				assert.Equal(t, b, instruction[i], "byte at index %d should be same", i)
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

	assert.Equal(t, expected, concatenated.String())
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
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			instruction := Make(test.opcode, test.operands...)

			def, err := Lookup(byte(test.opcode))
			assert.Nil(t, err)

			operandsRead, n := ReadOperands(def, instruction[1:])
			assert.Equal(t, test.bytesRead, n)
			assert.Equal(t, test.operands, operandsRead)
		})
	}
}
