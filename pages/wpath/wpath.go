// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package wpath handles pages paths.
package wpath

import (
	"path"
	"slices"
	"strings"
	"unicode"

	"cogentcore.org/core/base/strcase"
)

// Draft returns whether the given path is a draft page that
// should be ignored in released builds, which is the case
// if the path starts with a dash.
func Draft(p string) bool {
	return strings.HasPrefix(path.Base(p), "-")
}

// Format formats the given path into a correct pages path
// by removing all `{digit(s)}-` prefixes at the start of path
// segments, which are used for ordering files and folders and
// thus should not be displayed.
func Format(path string) string {
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if !strings.Contains(part, "-") {
			continue
		}
		parts[i] = strings.TrimLeftFunc(part, func(r rune) bool {
			return unicode.IsDigit(r) || r == '-'
		})
	}
	return strings.Join(parts, "/")
}

// Label returns a user friendly label for the given page URL,
// with the given backup name if the URL is blank.
func Label(u string, backup string) string {
	res := ""
	if u == "" {
		return backup
	}
	parts := strings.Split(u, "/")
	for i, part := range parts {
		parts[i] = strcase.ToSentence(part)
	}
	slices.Reverse(parts)
	res = strings.Join(parts, " • ")
	return res
}
