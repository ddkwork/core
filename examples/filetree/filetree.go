// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate core generate

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"cogentcore.org/core/events"
	"cogentcore.org/core/filetree"
	"cogentcore.org/core/gi"
	"cogentcore.org/core/giv"
	"cogentcore.org/core/icons"
	"cogentcore.org/core/keyfun"
	"cogentcore.org/core/ki"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/texteditor"
	"cogentcore.org/core/units"
)

// FileBrowse is a simple file browser / viewer / editor with a file tree and
// one or more editor windows.  It is based on an early version of the Gide
// IDE framework, and remains simple to test / demo the file tree component.
type FileBrowse struct {
	gi.Frame

	// root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename
	ProjRoot gi.Filename `desc:"root directory for the project -- all projects must be organized within a top-level root directory, with all the files therein constituting the scope of the project -- by default it is the path for ProjFilename"`

	// filename of the currently-active texteditor
	ActiveFilename gi.Filename `desc:"filename of the currently-active texteditor"`

	// has the root changed?  we receive update signals from root for changes
	Changed bool `json:"-" desc:"has the root changed?  we receive update signals from root for changes"`

	// all the files in the project directory and subdirectories
	Files *filetree.Tree `desc:"all the files in the project directory and subdirectories"`

	// number of texteditors available for editing files (default 2) -- configurable with n-text-views property
	NTextEditors int `xml:"n-text-views" desc:"number of texteditors available for editing files (default 2) -- configurable with n-text-views property"`

	// index of the currently-active texteditor -- new files will be viewed in other views if available
	ActiveTextEditorIndex int `json:"-" desc:"index of the currently-active texteditor -- new files will be viewed in other views if available"`
}

func (fb *FileBrowse) Defaults() {
	fb.NTextEditors = 2
}

// todo: rewrite with direct config, as a better example

func (fb *FileBrowse) OnInit() {
	fb.Defaults()
	fb.Style(func(s *styles.Style) {
		s.Direction = styles.Column
		s.Grow.Set(1, 1)
		s.Margin.Set(units.Dp(8))
	})
	fb.OnWidgetAdded(func(w gi.Widget) {
		switch w.PathFrom(fb) {
		case "title":
			title := w.(*gi.Label)
			title.Type = gi.LabelHeadlineSmall
			w.Style(func(s *styles.Style) {
				s.Justify.Content = styles.Center
			})
		}
		if w.Parent().PathFrom(fb) == "splits" {
			if w.IndexInParent() == 0 {
				w.Style(func(s *styles.Style) {
					s.Grow.Set(1, 1)
				})
			} else {
				w.Style(func(s *styles.Style) {
					s.Grow.Set(1, 1)
					s.Min.X.Ch(20)
					s.Min.Y.Ch(10)
					s.Font.Family = string(gi.AppearanceSettings.MonoFont)
					s.Text.WhiteSpace = styles.WhiteSpacePreWrap
					s.Text.TabSize = 4
				})
			}
		}
	})
}

// UpdateFiles updates the list of files saved in project
func (fb *FileBrowse) UpdateFiles() { //gti:add
	if fb.Files == nil {
		return
	}
	fb.Files.UpdateAll()
}

// IsEmpty returns true if given FileBrowse project is empty -- has not been set to a valid path
func (fb *FileBrowse) IsEmpty() bool {
	return fb.ProjRoot == ""
}

// OpenPath opens a new browser viewer at given path, which can either be a
// specific file or a directory containing multiple files of interest -- opens
// in current FileBrowse object if it is empty, or otherwise opens a new
// window.
func (fb *FileBrowse) OpenPath(path gi.Filename) { //gti:add
	if !fb.IsEmpty() {
		NewFileBrowser(string(path))
		return
	}
	fb.Defaults()
	root, pnm, fnm, ok := ProjPathParse(string(path))
	if !ok {
		return
	}
	fb.ProjRoot = gi.Filename(root)
	fb.SetName(pnm)
	fb.UpdateProj()
	fb.Files.OpenPath(root)
	// win := fb.ParentRenderWin()
	// if win != nil {
	// 	winm := "browser-" + pnm
	// 	win.SetName(winm)
	// 	win.SetTitle(winm)
	// }
	if fnm != "" {
		fb.ViewFile(fnm)
	}
	fb.UpdateFiles()
}

// UpdateProj does full update to current proj
func (fb *FileBrowse) UpdateProj() {
	fb.StandardConfig()
	fb.SetTitle(fmt.Sprintf("FileBrowse of: %v", fb.ProjRoot)) // todo: get rid of title
	fb.UpdateFiles()
	fb.ConfigSplits()
}

// ProjPathParse parses given project path into a root directory (which could
// be the path or just the directory portion of the path, depending in whether
// the path is a directory or not), and a bool if all is good (otherwise error
// message has been reported). projnm is always the last directory of the path.
func ProjPathParse(path string) (root, projnm, fnm string, ok bool) {
	if path == "" {
		return "", "blank", "", false
	}
	info, err := os.Lstat(path)
	if err != nil {
		emsg := fmt.Errorf("ProjPathParse: Cannot open at given path: %q: Error: %v", path, err)
		log.Println(emsg)
		return
	}
	dir, fn := filepath.Split(path)
	pathIsDir := info.IsDir()
	if pathIsDir {
		root = path
	} else {
		root = dir
		fnm = fn
	}
	_, projnm = filepath.Split(root)
	ok = true
	return
}

//////////////////////////////////////////////////////////////////////////////////////
//   TextEditors

// ActiveTextEditor returns the currently-active TextEditor
func (fb *FileBrowse) ActiveTextEditor() *texteditor.Editor {
	return fb.TextEditorByIndex(fb.ActiveTextEditorIndex)
}

// SetActiveTextEditor sets the given view index as the currently-active
// TextEditor -- returns that texteditor
func (fb *FileBrowse) SetActiveTextEditor(idx int) *texteditor.Editor {
	if idx < 0 || idx >= fb.NTextEditors {
		log.Printf("FileBrowse SetActiveTextEditor: text view index out of range: %v\n", idx)
		return nil
	}
	fb.ActiveTextEditorIndex = idx
	av := fb.ActiveTextEditor()
	if av.Buffer != nil {
		fb.ActiveFilename = av.Buffer.Filename
	}
	av.SetFocusEvent()
	return av
}

// NextTextEditor returns the next text view available for viewing a file and
// its index -- if the active text view is empty, then it is used, otherwise
// it is the next one
func (fb *FileBrowse) NextTextEditor() (*texteditor.Editor, int) {
	av := fb.TextEditorByIndex(fb.ActiveTextEditorIndex)
	if av.Buffer == nil {
		return av, fb.ActiveTextEditorIndex
	}
	nxt := (fb.ActiveTextEditorIndex + 1) % fb.NTextEditors
	return fb.TextEditorByIndex(nxt), nxt
}

// SaveActiveView saves the contents of the currently-active texteditor
func (fb *FileBrowse) SaveActiveView() { //gti:add
	tv := fb.ActiveTextEditor()
	if tv.Buffer != nil {
		tv.Buffer.Save() // todo: errs..
		fb.UpdateFiles()
	}
}

// SaveActiveViewAs save with specified filename the contents of the
// currently-active texteditor
func (fb *FileBrowse) SaveActiveViewAs(filename gi.Filename) { //gti:add
	tv := fb.ActiveTextEditor()
	if tv.Buffer != nil {
		tv.Buffer.SaveAs(filename)
	}
}

// ViewFileNode sets the next text view to view file in given node (opens
// buffer if not already opened)
func (fb *FileBrowse) ViewFileNode(fn *filetree.Node) {
	if _, err := fn.OpenBuf(); err == nil {
		nv, nidx := fb.NextTextEditor()
		if nv.Buffer != nil && nv.Buffer.IsNotSaved() { // todo: save current changes?
			fmt.Printf("Changes not saved in file: %v before switching view there to new file\n", nv.Buffer.Filename)
		}
		nv.SetBuffer(fn.Buffer)
		fn.Buffer.Hi.Style = "emacs" // todo prefs
		fb.SetActiveTextEditor(nidx)
		fb.UpdateFiles()
	}
}

// ViewFile sets the next text view to view given file name -- include as much
// of name as possible to disambiguate -- will use the first matching --
// returns false if not found
func (fb *FileBrowse) ViewFile(fnm string) bool {
	fn, ok := fb.Files.FindFile(fnm)
	if !ok {
		return false
	}
	fb.ViewFileNode(fn)
	return true
}

//////////////////////////////////////////////////////////////////////////////////////
//   GUI configs

// StandardFrameConfig returns a Config for configuring a standard Frame
// -- can modify as desired before calling ConfigChildren on Frame using this
func (fb *FileBrowse) StandardFrameConfig() ki.Config {
	config := ki.Config{}
	config.Add(gi.LabelType, "title")
	config.Add(gi.SplitsType, "splits")
	return config
}

// StandardConfig configures a standard setup of the overall Frame.
// It returns whether any modifications were made.
func (fb *FileBrowse) StandardConfig() bool {
	config := fb.StandardFrameConfig()
	return fb.ConfigChildren(config)
}

// SetTitle sets the optional title and updates the Title label
func (fb *FileBrowse) SetTitle(title string) {
	lab, _ := fb.TitleWidget()
	if lab != nil {
		lab.Text = title
	}
}

// Title returns the title label widget, and its index, within frame -- nil,
// -1 if not found
func (fb *FileBrowse) TitleWidget() (*gi.Label, int) {
	idx, ok := fb.Children().IndexByName("title", 0)
	if !ok {
		return nil, -1
	}
	return fb.Child(idx).(*gi.Label), idx
}

// Splits returns the main Splits
func (fb *FileBrowse) Splits() (*gi.Splits, int) {
	idx, ok := fb.Children().IndexByName("splits", 2)
	if !ok {
		return nil, -1
	}
	return fb.Child(idx).(*gi.Splits), idx
}

// TextEditorByIndex returns the TextEditor by index, nil if not found
func (fb *FileBrowse) TextEditorByIndex(idx int) *texteditor.Editor {
	if idx < 0 || idx >= fb.NTextEditors {
		log.Printf("FileBrowse: text view index out of range: %v\n", idx)
		return nil
	}
	split, _ := fb.Splits()
	stidx := 1 // 0 = file browser -- could be collapsed but always there.
	if split != nil {
		svk := split.Child(stidx + idx)
		return svk.(*texteditor.Editor)
	}
	return nil
}

func (fb *FileBrowse) ConfigToolbar(tb *gi.Toolbar) { //gti:add
	giv.NewFuncButton(tb, fb.UpdateFiles).SetIcon(icons.Refresh).SetShortcut("Command+U")
	op := giv.NewFuncButton(tb, fb.OpenPath).SetKey(keyfun.Open)
	op.Args[0].SetValue(fb.ActiveFilename)
	// op.Args[0].SetTag("ext", ".json")
	giv.NewFuncButton(tb, fb.SaveActiveView).SetKey(keyfun.Save)
	// save.SetUpdateFunc(func() {
	// 	save.SetEnabledUpdate(fb.Changed && ge.Filename != "")
	// })
	sa := giv.NewFuncButton(tb, fb.SaveActiveViewAs).SetKey(keyfun.SaveAs)
	sa.Args[0].SetValue(fb.ActiveFilename)
	// sa.Args[0].SetTag("ext", ".json")
}

// SplitsConfig returns a Config for configuring the Splits
func (fb *FileBrowse) SplitsConfig() ki.Config {
	config := ki.Config{}
	config.Add(gi.FrameType, "filetree-fr")
	for i := 0; i < fb.NTextEditors; i++ {
		config.Add(texteditor.EditorType, fmt.Sprintf("texteditor-%v", i))
	}
	// todo: tab view
	return config
}

// ConfigSplits configures the Splits.
func (fb *FileBrowse) ConfigSplits() {
	split, _ := fb.Splits()
	if split == nil {
		return
	}
	split.SetSplits(.2, .4, .4)

	config := fb.SplitsConfig()
	if split.ConfigChildren(config) {
		ftfr := split.Child(0).(*gi.Frame)
		fb.Files = filetree.NewTree(ftfr, "filetree")
		fb.Files.OnSelect(func(e events.Event) {
			e.SetHandled()
			if len(fb.Files.SelectedNodes) > 0 {
				sn, ok := fb.Files.SelectedNodes[0].This().(*filetree.Node)
				if ok {
					fb.FileNodeSelected(sn)
				}
			}
		})
		fb.Files.DoubleClickFun = func(e events.Event) {
			e.SetHandled()
			if len(fb.Files.SelectedNodes) > 0 {
				sn, ok := fb.Files.SelectedNodes[0].This().(*filetree.Node)
				if ok {
					fb.FileNodeOpened(sn)
				}
			}
		}
	}
}

func (fb *FileBrowse) FileNodeSelected(fn *filetree.Node) {
	fmt.Println("selected:", fn.FPath)
}

func (fb *FileBrowse) FileNodeOpened(fn *filetree.Node) {
	if fn.IsDir() {
		if !fn.HasChildren() {
			fn.OpenEmptyDir()
		} else {
			fn.ToggleClose()
		}
	} else {
		fb.ViewFileNode(fn)
		fn.UpdateNode()
	}
}

//////////////////////////////////////////////////////////////////////////////////////
//   Project window

// NewFileBrowser creates a new FileBrowse window with a new FileBrowse project for given
// path, returning the window and the path
func NewFileBrowser(path string) (*FileBrowse, *gi.Stage) {
	_, projnm, _, _ := ProjPathParse(path)

	b := gi.NewBody("Browser: " + projnm)
	fb := NewFileBrowse(b, "browser")
	b.AddAppBar(fb.ConfigToolbar)
	fb.OpenPath(gi.Filename(path))
	return fb, b.NewWindow().Run()
}

//////////////////////////////////////////////////////////////////////////////////////
//  main

func main() {
	var path string

	// process command args
	if len(os.Args) > 1 {
		flag.StringVar(&path, "path", "", "path to open -- can be to a directory or a filename within the directory")
		// todo: other args?
		flag.Parse()
		if path == "" {
			if flag.NArg() > 0 {
				path = flag.Arg(0)
			}
		}
	}
	if path == "" {
		path = "./"
	}
	if path != "" {
		path, _ = filepath.Abs(path)
	}
	fmt.Println("path:", path)
	_, st := NewFileBrowser(path)
	st.Wait()
}
