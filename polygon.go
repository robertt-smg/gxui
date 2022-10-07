// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

import (
	"gitlab.com/fti_ticketshop_pub/gxui/math"
)

type PolygonVertex struct {
	Position      math.Point
	RoundedRadius float32
}

type Polygon []PolygonVertex
