package timer

import (
	"github.com/hopeio/utils/io/fs/watch"
	"github.com/hopeio/utils/log"
	"os"
	"path/filepath"
	"time"
)

// only support Create,Remove,Write
type Watch struct {
	interval time.Duration
	done     chan struct{}
	handler  watch.Handler
}

func New(interval time.Duration) (*Watch, error) {
	return &Watch{
		interval: interval,
		done:     make(chan struct{}),
		handler:  make(watch.Handler),
	}, nil
}

func (w *Watch) Add(path string, op watch.Op, callback func(string)) error {
	path = filepath.Clean(path)
	var modTime time.Time
	info, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	if err == nil {
		modTime = info.ModTime()
	}

	handle, ok := w.handler[path]
	if !ok {
		handle = &watch.Callback{
			LastModTime: modTime,
		}
		w.handler[path] = handle
	}
	handle.Callbacks[op-1] = callback
	return nil
}

func (w *Watch) run(interval time.Duration) {
	timer := time.NewTicker(interval)
	for {
		select {
		case <-timer.C:
			for path, handle := range w.handler {
				var modTime time.Time
				info, err := os.Stat(path)
				if err != nil && !os.IsNotExist(err) {
					log.Error(err)
				}
				if err == nil {
					modTime = info.ModTime()
				}

				if !handle.LastModTime.IsZero() {
					if modTime.IsZero() {
						handle.LastModTime = modTime
						handle.Callbacks[watch.Remove-1](path)
					} else {
						if modTime.Sub(handle.LastModTime) > time.Second {
							handle.LastModTime = modTime
							handle.Callbacks[watch.Write-1](path)
						}
					}
				} else {
					if modTime.After(handle.LastModTime) {
						handle.LastModTime = modTime
						handle.Callbacks[watch.Create-1](path)
					}
				}

			}
		case <-w.done:
			timer.Stop()
			return
		}
	}
}

func (w *Watch) Close() error {
	close(w.done)
	return nil
}
