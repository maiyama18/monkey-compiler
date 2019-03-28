package vm

import (
	"monkey-compiler/ast"
	"monkey-compiler/compiler"
	"monkey-compiler/lexer"
	"monkey-compiler/object"
	"monkey-compiler/parser"
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
		{"1 + 2", 3},
		{"1 - 2", -1},
		{"2 * 3", 6},
		{"4 / 2", 2},
		{"4 / 2 * 2 + 10 - 5", 9},
		{"2 * (2 + 3)", 10},
		{"2 * 2 + 3", 7},
		{"-1", -1},
		{"-10 + 30 + -10", 10},
	}

	runVmTests(t, testCases)
}

func TestBooleanExpression(t *testing.T) {
	testCases := []vmTestCase{
		{"true;", true},
		{"false;", false},
		{"5 > 3;", true},
		{"5 < 3;", false},
		{"3 > 5;", false},
		{"3 < 5;", true},
		{"3 == 3;", true},
		{"3 != 3;", false},
		{"3 == 5;", false},
		{"3 != 5;", true},
		{"true == true;", true},
		{"true != true;", false},
		{"true == false;", false},
		{"true != false;", true},
		{"!true;", false},
		{"!false;", true},
		{"!!true;", true},
	}

	runVmTests(t, testCases)
}

func TestConditionals(t *testing.T) {
	testCases := []vmTestCase{
		{"if (true) { 10 }", 10},
		{"if (true) { 10 } else { 20 }", 10},
		{"if (false) { 10 } else { 20 }", 20},
		{"if (1) { 10 } else { 20 }", 10},
		{"if (1 < 2) { 10 }", 10},
		{"if (1 < 2) { 10 } else { 20 }", 10},
		{"if (1 > 2) { 10 } else { 20 }", 20},
	}

	runVmTests(t, testCases)
}

func runVmTests(t *testing.T, testCases []vmTestCase) {
	t.Helper()

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			c := compiler.New()
			program := parse(tc.input)

			if err := c.Compile(program); err != nil {
				t.Errorf("compiler error: %s", err)
			}

			vm := New(c.ByteCode())
			if err := vm.Run(); err != nil {
				t.Errorf("vm error: %s", err)
			}

			elem := vm.LastPopped()
			testObject(t, tc.expected, elem)
		})
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
	case bool:
		testBooleanObject(t, expected, actual)
	}
}

func testIntegerObject(t *testing.T, expected int64, actual object.Object) {
	t.Helper()

	actualInteger, ok := actual.(*object.Integer)
	if !ok {
		t.Errorf("could not convert to Integer: %+v", actual)
	}

	if actualInteger.Value != expected {
		t.Errorf("Integer valud wrong. want=%d, got=%d", expected, actualInteger.Value)
	}
}

func testBooleanObject(t *testing.T, expected bool, actual object.Object) {
	t.Helper()

	actualBoolean, ok := actual.(*object.Boolean)
	if !ok {
		t.Errorf("could not convert to Boolean: %+v", actual)
	}

	if actualBoolean.Value != expected {
		t.Errorf("Boolean valud wrong. want=%t, got=%t", expected, actualBoolean.Value)
	}
}
