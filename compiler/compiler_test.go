package compiler

import (
	"fmt"
	"github.com/muiscript/monkey-compiler/ast"
	"github.com/muiscript/monkey-compiler/code"
	"github.com/muiscript/monkey-compiler/lexer"
	"github.com/muiscript/monkey-compiler/object"
	"github.com/muiscript/monkey-compiler/parser"
	"testing"
)

type compilerTestCase struct {
	desc                 string
	input                string
	expectedConstants    []interface{}
	expectedInstructions code.Instructions
}

func TestIntegerArithmetic(t *testing.T) {
	tests := []compilerTestCase{
		{
			desc:              "Add",
			input:             "1 + 2",
			expectedConstants: []interface{}{1, 2},
			expectedInstructions: concatInstructions([]code.Instructions{
				code.Make(code.OpConstant, 0),
				code.Make(code.OpConstant, 1),
			}),
		},
	}

	runCompilerTests(t, tests)
}

func runCompilerTests(t *testing.T, tests []compilerTestCase) {
	t.Helper()

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			program := parse(tt.input)

			compiler := New()
			if err := compiler.Compile(program); err != nil {
				t.Fatalf("compiler error: %s\n", err)
			}

			bytecode := compiler.ByteCode()

			if err := testInstructions(tt.expectedInstructions, bytecode.Instructions); err != nil {
				t.Fatalf("testInstructions failed: %s", err)
			}

			if err := testConstants(tt.expectedConstants, bytecode.Constants); err != nil {
				t.Fatalf("testConstants failed: %s", err)
			}
		})
	}
}

func testInstructions(expected, actual code.Instructions) error {
	if len(actual) != len(expected)	{
		return fmt.Errorf("wrong instructions length. \nexpected=%q\nactual=%q", expected, actual)
	}

	for i, e := range expected {
		if actual[i] != e {
			return fmt.Errorf("wrong instruction at %d.\nexpected=%q\nactual=%q", i, e, actual[i])
		}
	}

	return nil
}

func testConstants(expected []interface{}, actual []object.Object) error {
	if len(actual) != len(expected)	{
		return fmt.Errorf("wrong instructions length. \nexpected=%d\nactual=%d", len(expected), len(actual))
	}

	for i, expectedConst := range expected {
		switch expectedConst := expectedConst.(type) {
		case int:
			if err := testIntegerObject(int64(expectedConst), actual[i]); err != nil {
				return fmt.Errorf("constant %d - testIntegerObject failed: %s", i, err)
			}
		}
	}

	return nil
}

func testIntegerObject(expected int64, actual object.Object) error {
	result, ok := actual.(*object.Integer)
	if !ok {
		return fmt.Errorf("object is not Integer. actual=%T (%+v)", actual, actual)
	}

	if result.Value != expected {
		return fmt.Errorf("object has wrong value. expected=%d, got=%d", expected, result.Value)
	}

	return nil
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
