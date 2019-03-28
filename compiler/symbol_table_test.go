package compiler

import "testing"

func TestDefine(t *testing.T) {
	expected := map[string]Symbol{
		"a": {Name: "a", Scope: GlobalScope, Index: 0},
		"b": {Name: "b", Scope: GlobalScope, Index: 1},
	}

	global := NewSymbolTable()

	a := global.Define("a")
	if a != expected["a"] {
		t.Fatalf("want a=%+v, got a=%+v", expected["a"], a)
	}
}

func TestResolveGlobal(t *testing.T) {
	global := NewSymbolTable()
	global.Define("a")
	global.Define("b")

	expected := []Symbol{
		{Name: "a", Scope: GlobalScope, Index: 0},
		{Name: "b", Scope: GlobalScope, Index: 1},
	}

	for _, expSym := range expected {
		actual, ok := global.Resolve(expSym.Name)
		if !ok {
			t.Fatalf("name '%s' could not be resolved", expSym.Name)
		}
		if actual != expSym {
			t.Fatalf("resolved '%s' wrong. want=%+v, got=%+v", expSym.Name, expSym, actual)
		}
	}
}
