/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package fsnotify

import (
	"github.com/hopeio/utils/os/fs/watch"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/hopeio/utils/log"
)

type Watch struct {
	*fsnotify.Watcher
	interval time.Duration
	done     chan struct{}
	handler  watch.Handler
}

func New(interval time.Duration) (*Watch, error) {
	watcher, err := fsnotify.NewWatcher()
	w := &Watch{
		Watcher:  watcher,
		interval: interval,
		done:     make(chan struct{}, 1),
		//1.map和数组做取舍
		handler: make(watch.Handler),
		//Handler:  make(map[string]map[fsnotify.Op]func()),
		//2.提高时间复杂度，用event做key，然后每次事件循环取值
		//Handler:  make(map[fsnotify.Event]func()),
	}

	if err == nil {
		go w.run()
	}

	return w, err
}

func (w *Watch) Add(path string, op fsnotify.Op, callback func(string)) error {
	path = filepath.Clean(path)
	handle, ok := w.handler[path]
	if !ok {
		err := w.Watcher.Add(path)
		if err != nil {
			return err
		}
		handle = &watch.Callback{}
		w.handler[path] = handle
	}
	handle.Callbacks[op-1] = callback
	return nil
}

func (w *Watch) run() {
	ev := &fsnotify.Event{}
OuterLoop:
	for {
		select {
		case event, ok := <-w.Watcher.Events:
			if !ok {
				return
			}
			log.Info("event:", event)
			ev = &event
			if handle, ok := w.handler[event.Name]; ok {
				now := time.Now()
				if now.Sub(handle.LastModTime) < w.interval && event == *ev {
					continue
				}
				handle.LastModTime = now
				for i := range handle.Callbacks {
					if event.Op&fsnotify.Op(i+1) == fsnotify.Op(i+1) && handle.Callbacks[i] != nil {
						handle.Callbacks[i](event.Name)
					}
				}
			}
		case err, ok := <-w.Watcher.Errors:
			if !ok {
				return
			}
			log.Error("error:", err)
		case <-w.done:
			break OuterLoop
		}
	}
}

func (w *Watch) Close() error {
	w.done <- struct{}{}
	close(w.done)
	return w.Watcher.Close()
}
