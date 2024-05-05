// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"cogentcore.org/core/shell/interpreter"
)

func main() {
	// logx.UserLevel = slog.LevelDebug
	// logx.UserLevel = slog.LevelWarn
	in := interpreter.NewInterpreter()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(in.Prompt())
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		in.Eval(line)
	}
}
