package basic

import (
	"gitlab.com/fti_ticketshop_pub/gxui"
	"gitlab.com/fti_ticketshop_pub/gxui/mixins"
)

func CreateTableLayout(theme *Theme) gxui.TableLayout {
	l := &mixins.TableLayout{}
	l.Init(l, theme)
	return l
}
