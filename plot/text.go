// Copyright (c) 2024, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plot

import (
	"cogentcore.org/core/math32"
	"cogentcore.org/core/paint"
	"cogentcore.org/core/styles"
	"cogentcore.org/core/styles/units"
)

// DefaultFontFamily specifies a default font for plotting.
// if not set, the standard Cogent Core default font is used.
var DefaultFontFamily = ""

// TextStyle specifies styling parameters for Text elements
type TextStyle struct {
	styles.FontRender

	// how to align text along the relevant dimension for the text element
	Align styles.Aligns

	// Padding is used in a case-dependent manner to add space around text elements
	Padding units.Value

	// rotation of the text
	Rotation float32
}

func (ts *TextStyle) Defaults() {
	ts.FontRender.Defaults()
	ts.Align = styles.Center
	if DefaultFontFamily != "" {
		ts.FontRender.Family = DefaultFontFamily
	}
}

func (ts *TextStyle) openFont(pt *Plot) {
	if ts.Font.Face == nil {
		paint.OpenFont(&ts.FontRender, &pt.Paint.UnitContext) // calls SetUnContext after updating metrics
	}
}

// Text specifies a single text element in a plot
type Text struct {

	// text string, which can use HTML formatting
	Text string

	// styling for this text element
	Style TextStyle

	// paintText is the [paint.Text] for the text.
	paintText paint.Text
}

func (tx *Text) Defaults() {
	tx.Style.Defaults()
}

// config is called during the layout of the plot, prior to drawing
func (tx *Text) config(pt *Plot) {
	uc := &pt.Paint.UnitContext
	fs := &tx.Style.FontRender
	fs.ToDots(uc)
	tx.Style.Padding.ToDots(uc)
	txln := float32(len(tx.Text))
	fht := fs.Size.Dots
	hsz := float32(12) * txln
	txs := &pt.StdTextStyle
	txs.OrientationHoriz = tx.Style.Rotation
	txs.Align = tx.Style.Align

	tx.paintText.SetHTML(tx.Text, fs, txs, uc, nil)
	tx.paintText.Layout(txs, fs, uc, math32.Vector2{X: hsz, Y: fht})
}

// startPosX returns the starting position for a horizontally-aligned text element,
// based on given width.  Text must have been config'd already.
func (tx *Text) startPosX(width float32) math32.Vector2 {
	pos := math32.Vector2{}
	pos.X = styles.AlignFactor(tx.Style.Align) * width
	switch tx.Style.Align {
	case styles.Center:
		pos.X -= 0.5 * tx.paintText.Size.X
	case styles.End:
		pos.X -= tx.paintText.Size.X
	}
	return pos
}

// Draw renders the text at given upper left position
func (tx *Text) draw(pt *Plot, pos math32.Vector2) {
	tx.paintText.Render(pt.Paint, pos)
}