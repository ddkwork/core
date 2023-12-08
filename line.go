// Copyright (c) 2018, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package svg

import (
	"goki.dev/mat32/v2"
)

// Line is a SVG line
type Line struct {
	NodeBase

	// position of the start of the line
	Start mat32.Vec2 `xml:"{x1,y1}"`

	// position of the end of the line
	End mat32.Vec2 `xml:"{x2,y2}"`
}

func (g *Line) SVGName() string { return "line" }

func (g *Line) OnInit() {
	g.End.Set(1, 1)
}

func (g *Line) CopyFieldsFrom(frm any) {
	fr := frm.(*Line)
	g.NodeBase.CopyFieldsFrom(&fr.NodeBase)
	g.Start = fr.Start
	g.End = fr.End
}

func (g *Line) SetPos(pos mat32.Vec2) {
	g.Start = pos
}

func (g *Line) SetSize(sz mat32.Vec2) {
	g.End = g.Start.Add(sz)
}

func (g *Line) LocalBBox() mat32.Box2 {
	bb := mat32.NewEmptyBox2()
	bb.ExpandByPoint(g.Start)
	bb.ExpandByPoint(g.End)
	hlw := 0.5 * g.LocalLineWidth()
	bb.Min.SetSubScalar(hlw)
	bb.Max.SetAddScalar(hlw)
	return bb
}

func (g *Line) Render(sv *SVG) {
	vis, pc := g.PushTransform(sv)
	if !vis {
		return
	}
	pc.Lock()
	pc.DrawLine(g.Start.X, g.Start.Y, g.End.X, g.End.Y)
	pc.Stroke()
	pc.Unlock()
	g.BBoxes(sv)

	if mrk := sv.MarkerByName(g, "marker-start"); mrk != nil {
		ang := mat32.Atan2(g.End.Y-g.Start.Y, g.End.X-g.Start.X)
		mrk.RenderMarker(sv, g.Start, ang, g.Paint.StrokeStyle.Width.Dots)
	}
	if mrk := sv.MarkerByName(g, "marker-end"); mrk != nil {
		ang := mat32.Atan2(g.End.Y-g.Start.Y, g.End.X-g.Start.X)
		mrk.RenderMarker(sv, g.End, ang, g.Paint.StrokeStyle.Width.Dots)
	}

	g.RenderChildren(sv)
	pc.PopTransformLock()
}

// ApplyTransform applies the given 2D transform to the geometry of this node
// each node must define this for itself
func (g *Line) ApplyTransform(sv *SVG, xf mat32.Mat2) {
	rot := xf.ExtractRot()
	if rot != 0 || !g.Paint.Transform.IsIdentity() {
		g.Paint.Transform = g.Paint.Transform.Mul(xf)
		g.SetProp("transform", g.Paint.Transform.String())
	} else {
		g.Start = xf.MulVec2AsPt(g.Start)
		g.End = xf.MulVec2AsPt(g.End)
		g.GradientApplyTransform(sv, xf)
	}
}

// ApplyDeltaTransform applies the given 2D delta transforms to the geometry of this node
// relative to given point.  Trans translation and point are in top-level coordinates,
// so must be transformed into local coords first.
// Point is upper left corner of selection box that anchors the translation and scaling,
// and for rotation it is the center point around which to rotate
func (g *Line) ApplyDeltaTransform(sv *SVG, trans mat32.Vec2, scale mat32.Vec2, rot float32, pt mat32.Vec2) {
	crot := g.Paint.Transform.ExtractRot()
	if rot != 0 || crot != 0 {
		xf, lpt := g.DeltaTransform(trans, scale, rot, pt, false) // exclude self
		mat := g.Paint.Transform.MulCtr(xf, lpt)
		g.Paint.Transform = mat
		g.SetProp("transform", g.Paint.Transform.String())
	} else {
		xf, lpt := g.DeltaTransform(trans, scale, rot, pt, true) // include self
		g.Start = xf.MulVec2AsPtCtr(g.Start, lpt)
		g.End = xf.MulVec2AsPtCtr(g.End, lpt)
		g.GradientApplyTransformPt(sv, xf, lpt)
	}
}

// WriteGeom writes the geometry of the node to a slice of floating point numbers
// the length and ordering of which is specific to each node type.
// Slice must be passed and will be resized if not the correct length.
func (g *Line) WriteGeom(sv *SVG, dat *[]float32) {
	SetFloat32SliceLen(dat, 4+6)
	(*dat)[0] = g.Start.X
	(*dat)[1] = g.Start.Y
	(*dat)[2] = g.End.X
	(*dat)[3] = g.End.Y
	g.WriteTransform(*dat, 4)
	g.GradientWritePts(sv, dat)
}

// ReadGeom reads the geometry of the node from a slice of floating point numbers
// the length and ordering of which is specific to each node type.
func (g *Line) ReadGeom(sv *SVG, dat []float32) {
	g.Start.X = dat[0]
	g.Start.Y = dat[1]
	g.End.X = dat[2]
	g.End.Y = dat[3]
	g.ReadTransform(dat, 4)
	g.GradientReadPts(sv, dat)
}
