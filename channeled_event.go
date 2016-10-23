// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gxui

import (
	"reflect"
	"sync"
)

type EventQueue interface {
	Inject(func())
}

type ChanneledEvent struct {
	sync.RWMutex
	base  EventBase
	queue EventQueue
}

func CreateChanneledEvent(signature interface{}, queue EventQueue) Event {
	e := &ChanneledEvent{
		queue: queue,
	}
	e.base.init(signature)
	baseUnlisten := e.base.unlisten
	e.base.unlisten = func(id int) {
		e.RLock()
		baseUnlisten(id)
		e.RUnlock()
	}
	return e
}

func (e *ChanneledEvent) Fire(args ...interface{}) {
	e.base.VerifyArguments(args)
	e.queue.Inject(func() {
		e.RLock()
		e.base.InvokeListeners(args)
		e.RUnlock()
	})
}

func (e *ChanneledEvent) Listen(listener interface{}) EventSubscription {
	e.Lock()
	res := e.base.Listen(listener)
	e.Unlock()
	return res
}

func (e *ChanneledEvent) ParameterTypes() []reflect.Type {
	return e.base.ParameterTypes()
}
