// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package shell

// Shell represents one running shell context.
type Shell struct{}

func NewShell() *Shell {
	sh := &Shell{}
	return sh
}