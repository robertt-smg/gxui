package basic

import (
	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/mixins"
)

func CreateTableLayout(theme *Theme) gxui.TableLayout {
	l := &mixins.TableLayout{}
	l.Init(l, theme)
	return l
}
