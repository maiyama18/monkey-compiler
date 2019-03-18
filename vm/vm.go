package vm

import (
	"errors"
	"monkey-compiler/code"
	"monkey-compiler/compiler"
	"monkey-compiler/object"
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
	if vm.sp == 0 {
		return nil
	}
	return vm.stack[vm.sp-1]
}

func (vm *VM) Run() error {
	for ip := 0; ip < len(vm.instructions); ip++ {
		opcode := code.Opcode(vm.instructions[ip])

		switch opcode {
		case code.OpConstant:
			index := code.ReadUint16(vm.instructions[ip+1:])
			if err := vm.push(vm.constants[index]); err != nil {
				return err
			}
			ip += 2
		case code.OpAdd:
			right := vm.pop()
			left := vm.pop()
			added := right.(*object.Integer).Value + left.(*object.Integer).Value
			if err := vm.push(&object.Integer{Value: added}); err != nil {
				return err
			}
		case code.OpPop:
			vm.pop()
		}
	}

	return nil
}

func (vm *VM) push(o object.Object) error {
	if vm.sp >= StackSize {
		return errors.New("stack overflow")
	}

	vm.stack[vm.sp] = o
	vm.sp++

	return nil
}

func (vm *VM) pop() object.Object {
	if vm.sp == 0 {
		return nil
	}

	o := vm.stack[vm.sp-1]
	vm.sp--
	return o
}

func (vm *VM) LastPopped() object.Object {
	return vm.stack[vm.sp]
}
