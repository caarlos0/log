package log

import (
	"iter"
	"maps"
	"slices"
	"sync"
)

type orderedMap struct {
	keys  []string
	inner map[string]any
	mu    sync.RWMutex
}

func newOrderedMap() *orderedMap {
	return &orderedMap{
		keys:  []string{},
		inner: map[string]any{},
	}
}

func (o *orderedMap) Copy() *orderedMap {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return &orderedMap{
		keys:  slices.Clone(o.keys),
		inner: maps.Clone(o.inner),
	}
}

func (o *orderedMap) Keys() []string {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return slices.Clone(o.keys)
}

func (o *orderedMap) Set(k string, v any) {
	o.mu.Lock()
	defer o.mu.Unlock()
	if !slices.Contains(o.keys, k) {
		o.keys = append(o.keys, k)
	}
	o.inner[k] = v
}

func (o *orderedMap) All() iter.Seq2[string, any] {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return func(yield func(string, any) bool) {
		for _, k := range o.keys {
			if !yield(k, o.inner[k]) {
				return
			}
		}
	}
}
