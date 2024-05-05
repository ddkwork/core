// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package interpreter

import (
	"fmt"
	"reflect"
	"strings"

	"cogentcore.org/core/base/errors"
	"cogentcore.org/core/shell"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

// Interpreter represents one running shell context
type Interpreter struct {
	Shell *shell.Shell

	Interp *interp.Interpreter
}

func NewInterpreter() *Interpreter {
	in := &Interpreter{}
	in.Interp = interp.New(interp.Options{}) // GoPath: "/Users/oreilly/go/src"})
	stdlib.Symbols["cogentcore.org/core/shell"] = map[string]reflect.Value{
		"Exec": reflect.ValueOf(shell.Exec),
	}
	in.Interp.Use(stdlib.Symbols) // this causes symbols to crash
	// errors.Log1(in.Interp.Eval(`import "cogentcore.org/core/shell"`))
	in.Shell = shell.NewShell()
	return in
}

func (in *Interpreter) Eval(ln string) string {
	eln := in.Shell.TranspileLine(ln)
	errors.Log1(in.Interp.Eval(eln))
	fmt.Println("exec:", eln)
	return eln
}

// SymbolByName returns the reflect.Value for given symbol name
// from the current Globals, Symbols (must call GetSymbols first)
func (in *Interpreter) SymbolByName(name string) (bool, reflect.Value) {
	globs := in.Interp.Globals()
	syms := in.Interp.Symbols("main") // note: cannot use ""

	nmpath := ""
	dotIdx := strings.Index(name, ".")
	if dotIdx > 0 {
		nmpath = name[:dotIdx]
		name = name[dotIdx+1:]
	}
	for path, sy := range syms {
		if nmpath != "" && path != nmpath {
			continue
		}
		for nm, v := range sy {
			if nm == name {
				return true, v
			}
		}
	}

	for nm, v := range globs {
		if nm == name {
			return true, v
		}
	}
	return false, reflect.Value{}
}