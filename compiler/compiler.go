package compiler

import (
	"github.com/muiscript/monkey-compiler/ast"
	"github.com/muiscript/monkey-compiler/code"
	"github.com/muiscript/monkey-compiler/object"
)

type ByteCode struct {
	Instructions code.Instructions
	Constants    []object.Object
}

type Compiler struct {
	instructions code.Instructions
	constants    []object.Object
}

func (c *Compiler) Compile(node ast.Node) error {
	return nil
}

func (c *Compiler) ByteCode() *ByteCode {
	return &ByteCode{
		Instructions: c.instructions,
		Constants:    c.constants,
	}
}
