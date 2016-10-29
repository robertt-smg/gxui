// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

type TextSelection struct {
	storedStart, storedEnd int
	start, end             int
	caretAtStart           bool
}

func CreateTextSelection(start, end int, caretAtStart bool) TextSelection {
	sel := TextSelection{start: start, end: end, caretAtStart: caretAtStart}
	if sel.start > sel.end {
		sel.start, sel.end = sel.end, sel.start
	}
	return sel.Store()
}

func (i TextSelection) Store() TextSelection {
	i.storedStart = i.start
	i.storedEnd = i.end
	return i
}
func (i TextSelection) Stored() (start, end int) { return i.storedStart, i.storedEnd }
func (i TextSelection) Length() int              { return i.end - i.start }
func (i TextSelection) Range() (start, end int)  { return i.start, i.end }
func (i TextSelection) Start() int               { return i.start }
func (i TextSelection) End() int                 { return i.end }
func (i TextSelection) First() int               { return i.start }
func (i TextSelection) Last() int                { return i.end - 1 }
func (i TextSelection) CaretAtStart() bool       { return i.caretAtStart }

func (t TextSelection) Offset(i int) TextSelection {
	t.start += i
	t.end += i
	return t
}

func (i TextSelection) Caret() int {
	if i.caretAtStart {
		return i.start
	}
	return i.end
}

func (i TextSelection) From() int { // TODO: Think of a better name for this function
	if i.caretAtStart {
		return i.end
	}
	return i.start
}

func (i TextSelection) Span() (start, end uint64) {
	return uint64(i.start), uint64(i.end)
}
