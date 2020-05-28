// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plt

import (
	"testing"
	"time"

	"github.com/cpmech/gosl/chk"
	"github.com/cpmech/gosl/io"
)

func Test_PlotSimpleCurve(tst *testing.T) {

	// title of test
	verbose()
	chk.PrintTitle("PlotSimpleCurve")

	Begin()
	Plot([]float64{0, 1, 2, 3}, []float64{0, 1, 2, 3})
	time.Sleep(3000 * time.Millisecond)
	Plot([]float64{0, 1, 2, 3}, []float64{0, 1, 4, 9}, CurveStyle{MarkerType: "*"})
	io.Pf("curves = %v\n", curves)
	Show()
}