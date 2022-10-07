// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package basic

import (
	"gitlab.com/fti_ticketshop_pub/gxui"
	"gitlab.com/fti_ticketshop_pub/gxui/mixins"

	"gitlab.com/fti_ticketshop_pub/gxui/math"
)

func CreateLabel(theme *Theme) gxui.Label {
	l := &mixins.Label{}
	l.Init(l, theme, theme.DefaultFont(), theme.LabelStyle.FontColor)
	l.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	return l
}
