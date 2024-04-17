// Copyright 2019 Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Initially copied from G3N: github.com/g3n/engine/math32
// Copyright 2016 The G3N Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// with modifications needed to suit Cogent Core functionality.

package math32

//"math"

type Spline struct {
	points []Vector3
}

func NewSpline(points []Vector3) Spline {
	sp := Spline{}
	sp.points = make([]Vector3, len(points))
	copy(sp.points, points)
	return sp
}

func (sp *Spline) InitFromArray(a []float32) {
	// PEND array of what ?
	//this.points = [];
	//for ( var i = 0; i < a.length; i ++ ) {
	//    this.points[ i ] = { x: a[ i ][ 0 ], y: a[ i ][ 1 ], z: a[ i ][ 2 ] };
	//}
}