package vm

import (
	"github.com/muiscript/monkey-compiler/ast"
	"github.com/muiscript/monkey-compiler/compiler"
	"github.com/muiscript/monkey-compiler/lexer"
	"github.com/muiscript/monkey-compiler/object"
	"github.com/muiscript/monkey-compiler/parser"
	"github.com/stretchr/testify/assert"
	"testing"
)

type vmTestCase struct {
	input    string
	expected interface{}
}

func TestIntegerArithmetic(t *testing.T) {
	testCases := []vmTestCase{
		{"1", 1},
		{"2", 2},
		{"1 + 2", 2}, // TODO: fix
	}

	runVmTests(t, testCases)
}

func runVmTests(t *testing.T, testCases []vmTestCase) {
	t.Helper()

	for _, tc := range testCases {
		c := compiler.New()
		program := parse(tc.input)

		if err := c.Compile(program); err != nil {
			t.Fatalf("compiler error: %v", err)
		}

		vm := New(c.ByteCode())
		if err := vm.Run(); err != nil {
			t.Fatalf("vm error: %v", err)
		}

		assert.Equal(t, tc.expected, vm.StackTop())
	}
}

func parse(input string) *ast.Program {
	l := lexer.New(input)
	p := parser.New(l)
	return p.ParseProgram()
}

func testObject(t *testing.T, expected interface{}, actual object.Object) {
	t.Helper()

	switch expected := expected.(type) {
	case int:
		testIntegerObject(t, int64(expected), actual)
	}
}

func testIntegerObject(t *testing.T, expected int64, actual object.Object) {
	t.Helper()

	actualInteger, ok := actual.(*object.Integer)
	assert.True(t, ok, "should be converted to Integer")

	assert.Equal(t, expected, actualInteger.Value, "should be equal")
}
