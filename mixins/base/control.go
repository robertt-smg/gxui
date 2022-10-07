// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

import (
	"gitlab.com/fti_ticketshop_pub/gxui"
	"gitlab.com/fti_ticketshop_pub/gxui/mixins/outer"
	"gitlab.com/fti_ticketshop_pub/gxui/mixins/parts"

	"gitlab.com/fti_ticketshop_pub/gxui/math"
)

type ControlOuter interface {
	gxui.Control
	outer.Painter
	outer.Redrawer
	outer.Relayouter
}

type Control struct {
	parts.Attachable
	parts.DrawPaint
	parts.InputEventHandler
	parts.Layoutable
	parts.Parentable
	parts.Visible
}

func (c *Control) Init(outer ControlOuter, theme gxui.Theme) {
	c.Attachable.Init(outer)
	c.DrawPaint.Init(outer, theme)
	c.Layoutable.Init(outer, theme)
	c.InputEventHandler.Init(outer)
	c.Parentable.Init(outer)
	c.Visible.Init(outer)

	// Interface compliance test
	_ = gxui.Control(c)
}

func (c *Control) DesiredSize(min, max math.Size) math.Size {
	return max
}

func (c *Control) ContainsPoint(p math.Point) bool {
	return c.IsVisible() && c.Size().Rect().Contains(p)
}
