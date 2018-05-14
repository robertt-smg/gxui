// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !js

package gl

import (
	"runtime"
	"strings"
)

// discoverUIGoRoutine finds and stores the program counter of the
// function 'applicationLoop' that must be in the callstack. The
// PC is stored so that AssertUIGoroutine can verify that the call
// came from the application loop (the UI go-routine).
func (d *driver) discoverUIGoRoutine() {
	for _, pc := range d.pcs[:runtime.Callers(2, d.pcs)] {
		name := runtime.FuncForPC(pc).Name()
		if strings.HasSuffix(name, "applicationLoop") {
			d.uiPC = pc
			return
		}
	}
	panic("applicationLoop was not found in the callstack")
}

func (d *driver) isUIGoroutine() bool {
	for _, pc := range d.pcs[:runtime.Callers(2, d.pcs)] {
		if pc == d.uiPC {
			return true
		}
	}
	return false
}

// AssertUIGoroutine will panic if d.Debug() == true *and* it is
// called from a goroutine that is not the UI goroutine.
func (d *driver) AssertUIGoroutine() {
	if !d.Debug() {
		return
	}
	if !d.isUIGoroutine() {
		panic("AssertUIGoroutine called on a go-routine that was not the UI go-routine")
	}
}
