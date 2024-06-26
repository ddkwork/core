// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slicesx

// ToAny converts a slice of a certain type to a []any slice.
func ToAny[E any](s []E) []any {
	as := make([]any, len(s))
	for i, v := range s {
		as[i] = v
	}
	return as
}
