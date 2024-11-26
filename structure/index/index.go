/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package index

import "sync"

// Index 索引
type Index[T any, I comparable] struct {
	idx   []I
	value []T
	sync.RWMutex
}

func NewIndex[T any, I comparable]() *Index[T, I] {
	return &Index[T, I]{
		idx:   make([]I, 0),
		value: make([]T, 0),
	}
}

func (i *Index[T, I]) Add(idx I, res T) {
	i.idx = append(i.idx, idx)
	i.value = append(i.value, res)
}

func (i *Index[T, I]) Get(idx I) T {
	i.RLock()
	defer i.RUnlock()
	for j, v := range i.idx {
		if v == idx {
			return i.value[j]
		}
	}
	return *new(T)
}

func (i *Index[T, I]) Set(idx I, v T) {
	i.RLock()
	defer i.RUnlock()
	for j, x := range i.idx {
		if x == idx {
			i.value[j] = v
		}
		return
	}
}

func (i *Index[T, I]) Remove(idx I) {
	i.Lock()
	defer i.Unlock()
	for j, v := range i.idx {
		if v == idx {
			i.idx = append(i.idx[0:j], i.idx[j:]...)
			i.value = append(i.value[0:j], i.value[j:]...)
			return
		}
	}
	return
}
