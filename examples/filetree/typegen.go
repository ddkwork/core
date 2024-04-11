// Code generated by "core generate"; DO NOT EDIT.

package main

import (
	"cogentcore.org/core/core"
	"cogentcore.org/core/filetree"
	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

// FileBrowseType is the [types.Type] for [FileBrowse]
var FileBrowseType = types.AddType(&types.Type{Name: "main.FileBrowse", IDName: "file-browse", Doc: "FileBrowse is a simple file browser / viewer / editor with a file tree and\none or more editor windows.  It is based on an early version of the Gide\nIDE framework, and remains simple to test / demo the file tree component.", Methods: []types.Method{{Name: "UpdateFiles", Doc: "UpdateFiles updates the list of files saved in project", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "OpenPath", Doc: "OpenPath opens a new browser viewer at given path, which can either be a\nspecific file or a directory containing multiple files of interest -- opens\nin current FileBrowse object if it is empty, or otherwise opens a new\nwindow.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"path"}}, {Name: "SaveActiveView", Doc: "SaveActiveView saves the contents of the currently-active texteditor", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}, {Name: "SaveActiveViewAs", Doc: "SaveActiveViewAs save with specified filename the contents of the\ncurrently-active texteditor", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"filename"}}, {Name: "ConfigToolbar", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Args: []string{"tb"}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "ProjRoot", Doc: "root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename"}, {Name: "ActiveFilename", Doc: "filename of the currently-active texteditor"}, {Name: "Changed", Doc: "has the root changed?  we receive update signals from root for changes"}, {Name: "Files", Doc: "all the files in the project directory and subdirectories"}, {Name: "NTextEditors", Doc: "number of texteditors available for editing files (default 2) -- configurable with n-text-views property"}, {Name: "ActiveTextEditorIndex", Doc: "index of the currently-active texteditor -- new files will be viewed in other views if available"}}, Instance: &FileBrowse{}})

// NewFileBrowse adds a new [FileBrowse] with the given name to the given parent:
// FileBrowse is a simple file browser / viewer / editor with a file tree and
// one or more editor windows.  It is based on an early version of the Gide
// IDE framework, and remains simple to test / demo the file tree component.
func NewFileBrowse(parent tree.Node, name ...string) *FileBrowse {
	return parent.NewChild(FileBrowseType, name...).(*FileBrowse)
}

// NodeType returns the [*types.Type] of [FileBrowse]
func (t *FileBrowse) NodeType() *types.Type { return FileBrowseType }

// New returns a new [*FileBrowse] value
func (t *FileBrowse) New() tree.Node { return &FileBrowse{} }

// SetProjRoot sets the [FileBrowse.ProjRoot]:
// root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename
func (t *FileBrowse) SetProjRoot(v core.Filename) *FileBrowse { t.ProjRoot = v; return t }

// SetActiveFilename sets the [FileBrowse.ActiveFilename]:
// filename of the currently-active texteditor
func (t *FileBrowse) SetActiveFilename(v core.Filename) *FileBrowse { t.ActiveFilename = v; return t }

// SetChanged sets the [FileBrowse.Changed]:
// has the root changed?  we receive update signals from root for changes
func (t *FileBrowse) SetChanged(v bool) *FileBrowse { t.Changed = v; return t }

// SetFiles sets the [FileBrowse.Files]:
// all the files in the project directory and subdirectories
func (t *FileBrowse) SetFiles(v *filetree.Tree) *FileBrowse { t.Files = v; return t }

// SetNTextEditors sets the [FileBrowse.NTextEditors]:
// number of texteditors available for editing files (default 2) -- configurable with n-text-views property
func (t *FileBrowse) SetNTextEditors(v int) *FileBrowse { t.NTextEditors = v; return t }

// SetActiveTextEditorIndex sets the [FileBrowse.ActiveTextEditorIndex]:
// index of the currently-active texteditor -- new files will be viewed in other views if available
func (t *FileBrowse) SetActiveTextEditorIndex(v int) *FileBrowse {
	t.ActiveTextEditorIndex = v
	return t
}

// SetTooltip sets the [FileBrowse.Tooltip]
func (t *FileBrowse) SetTooltip(v string) *FileBrowse { t.Tooltip = v; return t }
