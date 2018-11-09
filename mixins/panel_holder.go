// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mixins

import (
	"fmt"

	"github.com/nelsam/gxui"
	"github.com/nelsam/gxui/math"
	"github.com/nelsam/gxui/mixins/base"
)

type PanelTab interface {
	gxui.Control
	SetText(string)
	SetActive(bool)
}

type PanelTabCreater interface {
	CreatePanelTab() PanelTab
}

type PanelHolderOuter interface {
	base.ContainerNoControlOuter
	gxui.PanelHolder
	PanelTabCreater
}

type PanelEntry struct {
	Tab                   PanelTab
	Panel                 gxui.Control
	MouseDownSubscription gxui.EventSubscription
}

type PanelHolder struct {
	base.Container

	outer PanelHolderOuter

	theme     gxui.Theme
	tabLayout gxui.LinearLayout
	entries   []PanelEntry
	selected  PanelEntry
	left      PanelTab
	right     PanelTab

	begin int
	end   int
}

func insertIndex(holder gxui.PanelHolder, at math.Point) int {
	count := holder.End()
	bestIndex := count - 1
	bestScore := float32(1e20)
	score := func(point math.Point, index int) {
		score := point.Sub(at).Len()
		if score < bestScore {
			bestIndex = index
			bestScore = score
		}
	}
	for i := holder.Begin(); i < count; i++ {
		tab := holder.Tab(i)
		size := tab.Size()
		ml := math.Point{Y: size.H / 2}
		mr := math.Point{Y: size.H / 2, X: size.W}
		score(gxui.TransformCoordinate(ml, tab, holder), i)
		score(gxui.TransformCoordinate(mr, tab, holder), i+1)
	}
	return bestIndex
}

func beginTabDragging(holder gxui.PanelHolder, panel gxui.Control, name string, window gxui.Window) {
	var mms, mos gxui.EventSubscription
	mms = window.OnMouseMove(func(ev gxui.MouseEvent) {
		for _, c := range gxui.TopControlsUnder(ev.WindowPoint, ev.Window) {
			if over, ok := c.C.(gxui.PanelHolder); ok {
				insertAt := insertIndex(over, c.P)
				if over == holder {
					if insertAt > over.PanelIndex(panel) {
						insertAt--
					}
				}
				holder.RemovePanel(panel)
				holder = over
				holder.AddPanelAt(panel, name, insertAt)
				holder.Select(insertAt)
			}
		}
	})
	mos = window.OnMouseUp(func(gxui.MouseEvent) {
		mms.Unlisten()
		mos.Unlisten()
	})
}

func (p *PanelHolder) Init(outer PanelHolderOuter, theme gxui.Theme) {
	p.Container.Init(outer, theme)

	p.outer = outer
	p.theme = theme

	p.tabLayout = theme.CreateLinearLayout()
	p.tabLayout.SetDirection(gxui.LeftToRight)

	p.left = p.outer.CreatePanelTab()
	p.left.SetText("<-")
	p.left.OnClick(func(ev gxui.MouseEvent) {
		p.SelectPrev()
	})
	p.tabLayout.AddChild(p.left)

	p.right = p.outer.CreatePanelTab()
	p.right.SetText("->")
	p.right.OnClick(func(ev gxui.MouseEvent) {
		p.SelectNext()
	})
	p.tabLayout.AddChild(p.right)

	p.Container.AddChild(p.tabLayout)
	p.SetMargin(math.Spacing{L: 1, T: 2, R: 1, B: 1})
	p.SetMouseEventTarget(true) // For drag-drop targets

	// Interface compliance test
	_ = gxui.PanelHolder(p)
}

func (p *PanelHolder) LayoutChildren() {
	s := p.Size()
	tabHeight := p.tabLayout.DesiredSize(math.ZeroSize, s).H
	panelRect := math.CreateRect(0, tabHeight, s.W, s.H).Contract(p.Padding())

	for _, child := range p.Children() {
		if child.Control == p.tabLayout {
			child.Control.SetSize(math.Size{W: s.W, H: tabHeight})
			child.Offset = math.ZeroPoint
		} else {
			rect := panelRect.Contract(child.Control.Margin())
			child.Control.SetSize(rect.Size())
			child.Offset = rect.Min
		}
	}
}

func (p *PanelHolder) DesiredSize(min, max math.Size) math.Size {
	return max
}

func (p *PanelHolder) SelectedPanel() gxui.Control {
	return p.selected.Panel
}

// gxui.PanelHolder compliance
func (p *PanelHolder) AddPanel(panel gxui.Control, name string) {
	p.AddPanelAt(panel, name, len(p.entries))
}

func (p *PanelHolder) AddPanelAt(panel gxui.Control, name string, index int) {
	if index < 0 || index > p.PanelCount() {
		panic(fmt.Errorf("Index %d is out of bounds. Acceptable range: [%d - %d]",
			index, 0, p.PanelCount()))
	}
	tab := p.outer.CreatePanelTab()
	tab.SetText(name)
	mds := tab.OnMouseDown(func(ev gxui.MouseEvent) {
		p.Select(p.PanelIndex(panel))
		beginTabDragging(p.outer, panel, name, ev.Window)
	})

	p.entries = append(p.entries, PanelEntry{})
	copy(p.entries[index+1:], p.entries[index:])
	p.entries[index] = PanelEntry{
		Panel: panel,
		Tab:   tab,
		MouseDownSubscription: mds,
	}

	if p.selected.Panel == nil {
		p.Select(index)
	}
}

func (p *PanelHolder) RemovePanel(panel gxui.Control) {
	index := p.PanelIndex(panel)
	if index < 0 {
		panic("PanelHolder does not contain panel")
	}

	entry := p.entries[index]
	entry.MouseDownSubscription.Unlisten()
	p.entries = append(p.entries[:index], p.entries[index+1:]...)

	if panel == p.selected.Panel {
		if p.PanelCount() > 0 {
			p.Select(math.Max(index-1, 0))
		} else {
			p.Select(-1)
		}
	}
}

func (p *PanelHolder) Select(index int) {
	if index >= p.PanelCount() {
		panic(fmt.Errorf("Index %d is out of bounds. Acceptable range: [%d - %d]",
			index, -1, p.PanelCount()-1))
	}
	if p.selected.Panel != nil {
		p.selected.Tab.SetActive(false)
		p.Container.RemoveChild(p.selected.Panel)
	}

	if index >= 0 {
		p.selected = p.entries[index]
	} else {
		p.selected = PanelEntry{}
	}

	if p.selected.Panel != nil {
		p.Container.AddChild(p.selected.Panel)
		p.update()
		for !p.isSelectedShow() {
			if index > p.begin {
				p.begin++
			} else if index < p.begin {
				p.begin--
			} else {
				p.update()
				break
			}
			p.update()
		}
		p.selected.Tab.SetActive(true)
	} else {
		p.update()
	}
}

func (p *PanelHolder) SelectNext() {
	ind := p.PanelIndex(p.selected.Panel)
	N := p.PanelCount()
	if N != 0 {
		ind = ind + 1
		if ind == N {
			ind = 0
		}
		p.Select(ind)
	}
}

func (p *PanelHolder) SelectPrev() {
	ind := p.PanelIndex(p.selected.Panel)
	N := p.PanelCount()
	if N != 0 {
		ind = ind - 1
		if ind < 0 {
			ind = N - 1
		}
		p.Select(ind)
	}
}

func (p *PanelHolder) PanelCount() int {
	return len(p.entries)
}

func (p *PanelHolder) PanelIndex(panel gxui.Control) int {
	for i, e := range p.entries {
		if e.Panel == panel {
			return i
		}
	}
	return -1
}

func (p *PanelHolder) Panel(index int) gxui.Control {
	return p.entries[index].Panel
}

func (p *PanelHolder) Tab(index int) gxui.Control {
	return p.entries[index].Tab
}

func (p *PanelHolder) Begin() int {
	return p.begin
}

func (p *PanelHolder) End() int {
	return p.end
}

func (p *PanelHolder) update() {
	for len(p.tabLayout.Children()) != 2 {
		p.tabLayout.RemoveChildAt(1)
	}
	p.end = p.begin
	for i := p.begin; i < len(p.entries); i++ {
		index := i - p.begin + 1
		ctrl := p.Tab(i)
		p.tabLayout.AddChildAt(index, ctrl)
		p.end++
		s := p.Size().W
		ds := p.tabLayout.DesiredSize(math.ZeroSize, p.Size()).W
		if s != 0 && s == ds {
			p.tabLayout.RemoveChildAt(index)
			p.end--
			break
		}
	}
}

func (p *PanelHolder) isSelectedShow() bool {
	idx := p.PanelIndex(p.selected.Panel)
	return (idx >= p.begin && idx < p.end)
}
