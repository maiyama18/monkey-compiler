package compiler

import (
	"github.com/muiscript/monkey-compiler/ast"
	"github.com/muiscript/monkey-compiler/code"
	"github.com/muiscript/monkey-compiler/lexer"
	"github.com/muiscript/monkey-compiler/object"
	"github.com/muiscript/monkey-compiler/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

type compilerTestCase struct {
	desc                 string
	input                string
	expectedConstants    []interface{}
	expectedInstructions []code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	testCases := []compilerTestCase{
		{
			desc:              "1+2",
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: []code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			},
		},
	}

	runCompilerTests(t, testCases)
}

func runCompilerTests(t *testing.T, testCases []compilerTestCase) {
	t.Helper()

	for _, tc := range testCases {
		program := parse(tc.input)

		compiler := New()
		err := compiler.Compile(program)
		assert.Nil(t, err, "compile should return no error")

		byteCode := compiler.ByteCode()

		expectedInstructions := concatInstructions(tc.expectedInstructions)
		assert.Equal(t, expectedInstructions, byteCode.Instructions)

		assert.Equal(t, len(tc.expectedConstants), len(byteCode.Constants), "the length of constants should be same")
		for i, c := range tc.expectedConstants {
			switch c := c.(type) {
			case int:
				testIntegerObject(t, int64(c), byteCode.Constants[i])
			}
		}
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func concatInstructions(instructions []code.Instructions) code.Instructions {
	out := code.Instructions{}
	for _, ins := range instructions {
		out = append(out, ins...)
	}
	return out
}

func testIntegerObject(t *testing.T, expected int64, actual object.Object) {
	t.Helper()

	actualInteger, ok := actual.(*object.Integer)
	assert.True(t, ok, "should be converted to Integer")

	assert.Equal(t, expected, actualInteger.Value, "should be equal")
}
