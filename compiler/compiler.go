package compiler

import (
	"github.com/muiscript/monkey-compiler/ast"
	"github.com/muiscript/monkey-compiler/code"
	"github.com/muiscript/monkey-compiler/object"
)

// ByteCode is byte code generated by compiler
type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

// Compiler is compiler of monkey
type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

// New returns empty compiler
func New() *Compiler {
	return &Compiler{
		code.Instructions{},
		[]object.Object{},
	}
}

// Compile ...
func (c *Compiler) Compile(node ast.Node) error {
	switch node := node.(type) {
	case *ast.Program:
		for _, stmt := range node.Statements {
			if err := c.Compile(stmt); err != nil {
				return err
			}
		}
	case *ast.ExpressionStatement:
		return c.Compile(node.Expression)
	case *ast.InfixExpression:
		if err := c.Compile(node.Left); err != nil {
			return err
		}
		if err := c.Compile(node.Right); err != nil {
			return err
		}
	case *ast.IntegerLiteral:
		integer := &object.Integer{Value: node.Value}
		c.emit(code.OpConstant, c.addConstant(integer))
	}

	return nil
}

// ByteCode ...
func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}

// returns position of start of added instruction
func (c *Compiler) emit(opcode code.Opcode, operands ...int) int {
	ins := code.Make(opcode, operands...)
	pos := c.addInstruction(ins)
	return pos
}

func (c *Compiler) addInstruction(ins []byte) int {
	pos := len(c.instructions)
	c.instructions = append(c.instructions, ins...)
	return pos
}

// returns index of added object in c.constants
func (c *Compiler) addConstant(obj object.Object) int {
	c.constants = append(c.constants, obj)
	return len(c.constants) - 1
}
