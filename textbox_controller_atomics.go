package gxui

import (
	"sync/atomic"
	"unsafe"
)

type dumbUnlistener func()

func (d dumbUnlistener) Unlisten() {
	d()
}

func (t *TextBoxController) text() []rune {
	ptr := (*[]rune)(atomic.LoadPointer(&t.textPtr))
	if ptr == nil {
		return nil
	}
	ret := make([]rune, len(*ptr))
	copy(ret, *ptr)
	return ret
}

func (t *TextBoxController) setText(text []rune) {
	n := make([]rune, len(text))
	copy(n, text)
	atomic.StorePointer(&t.textPtr, unsafe.Pointer(&n))
}

func (t *TextBoxController) locationHistoryIndex() int64 {
	return atomic.LoadInt64(&t.locationHistoryIndexSrc)
}

func (t *TextBoxController) setLocationHistoryIndex(i int64) {
	atomic.StoreInt64(&t.locationHistoryIndexSrc, i)
}

func (t *TextBoxController) storeCaretLocationsNextEdit() bool {
	return atomic.LoadUint32(&t.storeCaretLocationsNextEditSrc) == 1
}

func (t *TextBoxController) setStoreCaretLocationsNextEdit(v bool) {
	var i uint32
	if v {
		i = 1
	}
	atomic.StoreUint32(&t.storeCaretLocationsNextEditSrc, i)
}

func (t *TextBoxController) indent() string {
	ptr := (*string)(atomic.LoadPointer(&t.indentPtr))
	if ptr == nil {
		return ""
	}
	return *ptr
}

func (t *TextBoxController) setIndent(indent string) {
	atomic.StorePointer(&t.indentPtr, unsafe.Pointer(&indent))
}

func (t *TextBoxController) lineStarts() []int {
	ptr := (*[]int)(atomic.LoadPointer(&t.lineStartsPtr))
	if ptr == nil {
		return nil
	}
	ret := make([]int, len(*ptr))
	copy(ret, *ptr)
	return ret
}

func (t *TextBoxController) setLineStarts(l []int) {
	n := make([]int, len(l))
	copy(n, l)
	atomic.StorePointer(&t.lineStartsPtr, unsafe.Pointer(&n))
}

func (t *TextBoxController) lineEnds() []int {
	ptr := (*[]int)(atomic.LoadPointer(&t.lineEndsPtr))
	if ptr == nil {
		return nil
	}
	ret := make([]int, len(*ptr))
	copy(ret, *ptr)
	return ret
}

func (t *TextBoxController) setLineEnds(l []int) {
	n := make([]int, len(l))
	copy(n, l)
	atomic.StorePointer(&t.lineEndsPtr, unsafe.Pointer(&n))
}

func (t *TextBoxController) locationHistory() [][]int {
	ptr := (*[][]int)(atomic.LoadPointer(&t.locationHistoryPtr))
	if ptr == nil {
		return nil
	}
	v := *ptr
	ret := make([][]int, len(v))
	for i := range ret {
		ret[i] = make([]int, len(v[i]))
		copy(ret[i], v[i])
	}
	return ret
}

func (t *TextBoxController) setLocationHistory(h [][]int) {
	n := make([][]int, len(h))
	for i := range n {
		n[i] = make([]int, len(h[i]))
		copy(n[i], h[i])
	}
	atomic.StorePointer(&t.locationHistoryPtr, unsafe.Pointer(&n))
}

func (t *TextBoxController) selections() []TextSelection {
	ptr := (*[]TextSelection)(atomic.LoadPointer(&t.selectionsPtr))
	if ptr == nil {
		return []TextSelection{{}} // there must always be one!
	}
	ret := make([]TextSelection, len(*ptr))
	copy(ret, *ptr)
	if len(ret) == 0 {
		ret = append(ret, TextSelection{})
	}
	return ret
}

func (t *TextBoxController) setSelections(s []TextSelection) {
	n := make([]TextSelection, len(s))
	copy(n, s)
	atomic.StorePointer(&t.selectionsPtr, unsafe.Pointer(&n))
}

func (t *TextBoxController) nextSelectionChangedID() int64 {
	return atomic.AddInt64(&t.nextSelectionChangedIDSrc, 1)
}

func (t *TextBoxController) onSelectionChanged() map[int64]func() {
	ptr := (*map[int64]func())(atomic.LoadPointer(&t.onSelectionChangedPtr))
	if ptr == nil {
		return nil
	}
	ret := make(map[int64]func())
	for i, f := range *ptr {
		ret[i] = f
	}
	return ret
}

func (t *TextBoxController) selectionChanged() {
	for _, f := range t.onSelectionChanged() {
		f()
	}
}

func (t *TextBoxController) setOnSelectionChanged(s map[int64]func()) {
	n := make(map[int64]func())
	for i, f := range s {
		n[i] = f
	}
	atomic.StorePointer(&t.onSelectionChangedPtr, unsafe.Pointer(&n))
}

func (t *TextBoxController) nextTextChangedID() int64 {
	return atomic.AddInt64(&t.nextTextChangedIDSrc, 1)
}

func (t *TextBoxController) onTextChanged() map[int64]func([]TextBoxEdit) {
	ptr := (*map[int64]func([]TextBoxEdit))(atomic.LoadPointer(&t.onTextChangedPtr))
	if ptr == nil {
		return nil
	}
	ret := make(map[int64]func([]TextBoxEdit))
	for i, f := range *ptr {
		ret[i] = f
	}
	return ret
}

func (t *TextBoxController) textChanged(e []TextBoxEdit) {
	for _, f := range t.onTextChanged() {
		f(e)
	}
}

func (t *TextBoxController) setOnTextChanged(s map[int64]func([]TextBoxEdit)) {
	n := make(map[int64]func([]TextBoxEdit))
	for i, f := range s {
		n[i] = f
	}
	atomic.StorePointer(&t.onTextChangedPtr, unsafe.Pointer(&n))
}
