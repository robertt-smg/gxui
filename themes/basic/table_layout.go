package basic

import (
	"github.com/robertt-smg/gxui"
	"github.com/robertt-smg/gxui/mixins"
)

func CreateTableLayout(theme *Theme) gxui.TableLayout {
	l := &mixins.TableLayout{}
	l.Init(l, theme)
	return l
}
