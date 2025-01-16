/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package syncmap

import "sync"

type Map[K comparable, V any] struct {
	m map[K]V
	l sync.RWMutex
}

func (s *Map[K, V]) Set(key K, value V) {
	s.l.Lock()
	defer s.l.Unlock()
	s.m[key] = value
}

func (s *Map[K, V]) Get(key K) V {
	s.l.RLock()
	defer s.l.RUnlock()
	return s.m[key]
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: make(map[K]V)}
}
