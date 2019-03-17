package vm

import (
	"github.com/muiscript/monkey-compiler/code"
	"github.com/muiscript/monkey-compiler/compiler"
	"github.com/muiscript/monkey-compiler/object"
)

const StackSize = 2048

type VM struct {
	constants    []object.Object
	instructions code.Instructions

	stack []object.Object
	sp    int // stack pointer. top of the stack is stack[sp-1]
}

func New(byteCode *compiler.ByteCode) *VM {
	return &VM{
		constants:    byteCode.Constants,
		instructions: byteCode.Instructions,

		stack: make([]object.Object, StackSize),
		sp:    0,
	}
}

func (vm *VM) StackTop() object.Object {
	return nil
}

func (vm *VM) Run() error {
	return nil
}
