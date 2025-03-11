package loader

import (
	"github.com/fsnotify/fsnotify"
	"github.com/hopeio/utils/log"
	"io"
	"os"
	"slices"
	"time"
)

type ConfigLoader struct {
	// 间隔大于1秒采用timer定时加载，小于1秒用fsnotify
	ReloadInterval time.Duration
	Paths          []string
	watcher        *fsnotify.Watcher
	timer          *time.Ticker
	modTime        time.Time
}

// New initialize a Loader
func NewConfigLoader(interval time.Duration, filepaths ...string) *ConfigLoader {
	return &ConfigLoader{ReloadInterval: interval, Paths: filepaths}
}

func (ld *ConfigLoader) Close() error {
	if ld.watcher != nil {
		return ld.watcher.Close()
	}
	if ld.timer != nil {
		ld.timer.Stop()
	}
	return nil
}

// Load will unmarshal configurations to struct from files that you provide
func (ld *ConfigLoader) Handle(handle func(io.Reader)) (err error) {
	err = load(handle, ld.Paths...)
	if err != nil {
		return err
	}
	ld.modTime = time.Now()
	if ld.ReloadInterval > 0 {
		if ld.ReloadInterval >= time.Second {
			ld.timer = time.NewTicker(ld.ReloadInterval)
			go ld.watchTimer(handle)
		} else {
			watcher, err := fsnotify.NewWatcher()
			if err != nil {
				ld.timer = time.NewTicker(time.Second)
				go ld.watchTimer(handle)
				return nil
			} else {
				for _, filepath := range ld.Paths {
					err = watcher.Add(filepath)
					if err != nil {
						return err
					}
				}
				ld.watcher = watcher
				go ld.watchNotify(handle)
			}
		}
	}

	return
}

func (ld *ConfigLoader) watchNotify(handle func(reader io.Reader)) {
	for {
		select {
		case event, ok := <-ld.watcher.Events:
			if !ok {
				return
			}
			now := time.Now()
			if now.Sub(ld.modTime) < time.Second {
				continue
			}
			ld.modTime = now
			if event.Op&fsnotify.Write == fsnotify.Write {
				idx := slices.Index(ld.Paths, event.Name)
				if err := load(handle, ld.Paths[idx:]...); err != nil {
					log.Errorf("failed to reload data from %v, got error %v\n", ld.Paths, err)
				}
			}
		case err, ok := <-ld.watcher.Errors:
			if !ok {
				return
			}
			log.Error(err)
		}
	}
}

func (ld *ConfigLoader) watchTimer(handle func(reader io.Reader)) {

	for range ld.timer.C {
		for i := range ld.Paths {
			file := ld.Paths[i]
			// check configuration
			if fileInfo, err := os.Stat(file); err == nil && fileInfo.Mode().IsRegular() {
				if fileInfo.ModTime().After(ld.modTime) {
					ld.modTime = fileInfo.ModTime()
					if err := load(handle, ld.Paths[i:]...); err != nil {
						log.Error("failed to reload data from %v, got error %v\n", ld.Paths, err)
					}
					break
				}
			}
		}
	}
}
