// Code generated by "core generate"; DO NOT EDIT.

package pages

import (
	"io/fs"

	"cogentcore.org/core/tree"
	"cogentcore.org/core/types"
)

// PageType is the [types.Type] for [Page]
var PageType = types.AddType(&types.Type{Name: "cogentcore.org/core/pages.Page", IDName: "page", Doc: "Page represents a content page with support for navigating\nto other pages within the same source content.", Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "Source", Doc: "Source is the filesystem in which the content is located."}, {Name: "Context", Doc: "Context is the page's [htmlview.Context]."}, {Name: "History", Doc: "The history of URLs that have been visited. The oldest page is first."}, {Name: "HistoryIndex", Doc: "HistoryIndex is the current place we are at in the History"}, {Name: "PagePath", Doc: "PagePath is the fs path of the current page in [Page.Source]"}, {Name: "URLToPagePath", Doc: "URLToPagePath is a map between user-facing page URLs and underlying\nFS page paths."}}, Instance: &Page{}})

// NewPage adds a new [Page] with the given name to the given parent:
// Page represents a content page with support for navigating
// to other pages within the same source content.
func NewPage(parent tree.Node, name ...string) *Page {
	return parent.NewChild(PageType, name...).(*Page)
}

// NodeType returns the [*types.Type] of [Page]
func (t *Page) NodeType() *types.Type { return PageType }

// New returns a new [*Page] value
func (t *Page) New() tree.Node { return &Page{} }

// SetSource sets the [Page.Source]:
// Source is the filesystem in which the content is located.
func (t *Page) SetSource(v fs.FS) *Page { t.Source = v; return t }

// SetTooltip sets the [Page.Tooltip]
func (t *Page) SetTooltip(v string) *Page { t.Tooltip = v; return t }
