// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package discovery

import (
	"sync"

	"v.io/v23/context"
	"v.io/v23/discovery"
	"v.io/v23/verror"
)

const pkgPath = "v.io/x/ref/runtime/internal/discovery"

var (
	errDiscoveryClosed = verror.Register(pkgPath+".errDiscoveryClosed", verror.NoRetry, "{1:}{2:} discovery closed")
)

// ds is an implementation of discovery.T.
type ds struct {
	plugins []Plugin

	mu     sync.Mutex
	closed bool                  // GUARDED_BY(mu)
	tasks  map[*context.T]func() // GUARDED_BY(mu)
	wg     sync.WaitGroup

	ads map[string]struct{} // GUARDED_BY(mu)
}

func (ds *ds) Close() {
	ds.mu.Lock()
	if ds.closed {
		ds.mu.Unlock()
		return
	}
	for _, cancel := range ds.tasks {
		cancel()
	}
	ds.closed = true
	ds.mu.Unlock()
	ds.wg.Wait()
}

func (ds *ds) addTask(ctx *context.T) (*context.T, func(), error) {
	ds.mu.Lock()
	if ds.closed {
		ds.mu.Unlock()
		return nil, nil, verror.New(errDiscoveryClosed, ctx)
	}
	ctx, cancel := context.WithCancel(ctx)
	ds.tasks[ctx] = cancel
	ds.wg.Add(1)
	ds.mu.Unlock()
	return ctx, cancel, nil
}

func (ds *ds) removeTask(ctx *context.T) {
	ds.mu.Lock()
	if _, exist := ds.tasks[ctx]; exist {
		delete(ds.tasks, ctx)
		ds.wg.Done()
	}
	ds.mu.Unlock()
}

// New returns a new Discovery instance initialized with the given plugins.
//
// Mostly for internal use. Consider to use factory.New.
func NewWithPlugins(plugins []Plugin) discovery.T {
	if len(plugins) == 0 {
		panic("no plugins")
	}
	ds := &ds{
		plugins: make([]Plugin, len(plugins)),
		tasks:   make(map[*context.T]func()),
		ads:     make(map[string]struct{}),
	}
	copy(ds.plugins, plugins)
	return ds
}
