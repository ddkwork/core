// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package paint

import (
	"fmt"

	"goki.dev/colors"
	"goki.dev/girl/styles"
	"goki.dev/mat32/v2"
)

// DrawBox calls DrawBorder with position, size and border parameters
// as a convenience method for DrawStdBox
func (pc *Paint) DrawBox(rs *State, pos mat32.Vec2, sz mat32.Vec2, bs styles.Border) {
	pc.DrawBorder(rs, pos.X, pos.Y, sz.X, sz.Y, bs)
}

// DrawStdBox draws the CSS "standard box" model using given style.
// This is used for rendering widgets such as buttons, textfields, etc in a GUI.
func (pc *Paint) DrawStdBox(rs *State, st *styles.Style, pos mat32.Vec2, sz mat32.Vec2, surroundBgColor *colors.Full) {
	mpos := pos.Add(st.EffMargin().Pos())
	msz := sz.Sub(st.EffMargin().Size())
	rad := st.Border.Radius.Dots()

	// the background color we actually use
	bg := st.BackgroundColor
	sbg := surroundBgColor
	if bg.IsNil() {
		// we need to do this to prevent
		// elements from rendering over themselves
		// (see https://github.com/goki/gi/issues/565)
		bg = *sbg
	}

	// If a state layer is not specified but our state indicates we should have one,
	// we set it to the one specified by the state. This allows for state layers to
	// work for states automatically by default, but allow overriding it to custom values
	// if wanted. The user can always set StateLayer to -1 to get no state layer.
	if ssl := st.State.StateLayer(); st.StateLayer == 0 && ssl != 0 {
		st.StateLayer = ssl
	}
	// TODO: support state layers on gradient backgrounds
	if st.StateLayer > 0 && st.BackgroundColor.Gradient == nil {
		bg.Solid = colors.AlphaBlend(bg.Solid, colors.SetAF32(st.Color, st.StateLayer/100))
		fmt.Println(bg.Solid, st.Color, st.StateLayer/100)
	}

	// We need to fill the whole box where the
	// box shadows / element can go to prevent growing
	// box shadows and borders. We couldn't just
	// do this when there are box shadows, as they
	// may be removed and then need to be covered up.
	// This also fixes https://github.com/goki/gi/issues/579.
	// This isn't an ideal solution because of performance,
	// so TODO: maybe come up with a better solution for this.
	// We need to use raw LayState data because we need to clear
	// any box shadow that may have gone in margin.
	mspos, mssz := st.BoxShadowPosSize(pos, sz)
	pc.FillBox(rs, mspos, mssz, sbg)

	// first do any shadow
	if st.HasBoxShadow() {
		for _, shadow := range st.BoxShadow {
			pc.StrokeStyle.SetColor(nil)
			pc.FillStyle.SetColor(shadow.Color)

			// TODO: better handling of opacity?
			prevOpacity := pc.FillStyle.Opacity
			pc.FillStyle.Opacity = float32(shadow.Color.A) / 255
			// we only want radius for border, no actual border
			pc.DrawBox(rs, shadow.BasePos(mpos), shadow.BaseSize(msz), styles.Border{Radius: st.Border.Radius})
			// pc.FillStyle.Opacity = 1.0
			if shadow.Blur.Dots != 0 {
				// must divide by 2^2 (4) like CSS
				pc.BlurBox(rs, shadow.Pos(mpos), shadow.Size(msz), shadow.Blur.Dots/4)
			}
			pc.FillStyle.Opacity = prevOpacity
		}
	}

	// then draw the box over top of that.
	// need to set clipping to box first.. (?)
	// we need to draw things twice here because we need to clear
	// the whole area with the background color first so the border
	// doesn't render weirdly
	if rad.IsZero() {
		pc.FillBox(rs, mpos, msz, &bg)
	} else {
		pc.FillStyle.SetFullColor(&bg)
		// no border -- fill only
		pc.DrawRoundedRectangle(rs, mpos.X, mpos.Y, msz.X, msz.Y, rad)
		pc.Fill(rs)
	}

	// pc.StrokeStyle.SetColor(&st.Border.Color)
	// pc.StrokeStyle.Width = st.Border.Width
	// pc.FillStyle.SetFullColor(&st.BackgroundColor)
	mpos.SetAdd(st.Border.Width.Dots().Pos().MulScalar(0.5))
	msz.SetSub(st.Border.Width.Dots().Size().MulScalar(0.5))
	pc.FillStyle.SetColor(nil)
	// now that we have drawn background color
	// above, we can draw the border
	pc.DrawBox(rs, mpos, msz, st.Border)
}
