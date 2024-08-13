package loader

import (
	"github.com/fsnotify/fsnotify"
	"github.com/hopeio/utils/log"
	"os"
	"time"
)

type Loader struct {
	AutoReloadType ReloadType `json:"autoReloadType" comment:"none,fsnotify,timer"` // 本地分为Watch和AutoReload，Watch采用系统调用通知，AutoReload定时器去查文件是否变更
	TimerInterval  time.Duration
}

type ReloadType string

const (
	ReloadTypeNone     = "none"
	ReloadTypeFsNotify = "fsnotify"
	ReloadTypeTimer    = "timer"
)

// New initialize a Loader
func New() *Loader {
	return &Loader{}
}

func (ld *Loader) SetAutoReloadType(autoReloadType ReloadType, interval time.Duration) {
	ld.AutoReloadType = autoReloadType
	if autoReloadType == ReloadTypeTimer && interval < time.Second {
		interval = interval * time.Second
	}
	ld.TimerInterval = interval
}

// Load will unmarshal configurations to struct from files that you provide
func (ld *Loader) Handle(handle func([]byte), files ...string) (err error) {

	err = load(handle, files...)
	if err != nil {
		return err
	}
	if ld.AutoReloadType != "" && ld.AutoReloadType != ReloadTypeNone {
		if ld.AutoReloadType == ReloadTypeTimer {
			go ld.watchTimer(handle, files...)
		} else {
			go ld.watchNotify(handle, files...)
		}
	}

	return
}

func (ld *Loader) watchNotify(handle func([]byte), files ...string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
	}
	defer watcher.Close()
	for _, file := range files {
		err = watcher.Add(file)
		if err != nil {
			log.Error(err)
		}
	}

	interval := make(map[string]time.Time)

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			now := time.Now()
			if now.Sub(interval[event.Name]) < time.Second {
				continue
			}
			interval[event.Name] = now
			//log.Info("event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Info("modified file:", event.Name)
				if err := load(handle, event.Name); err != nil {
					log.Error("failed to reload data from %v, got error %v\n", files, err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Error("error:", err)
		}
	}
}

func (ld *Loader) watchTimer(handle func([]byte), files ...string) {
	var fileModTimes map[string]time.Time
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]

		// check configuration
		if fileInfo, err := os.Stat(file); err == nil && fileInfo.Mode().IsRegular() {
			fileModTimes[file] = fileInfo.ModTime()
		}
	}
	timer := time.NewTicker(ld.TimerInterval)
	for range timer.C {
		for i := len(files) - 1; i >= 0; i-- {
			file := files[i]

			// check configuration
			if fileInfo, err := os.Stat(file); err == nil && fileInfo.Mode().IsRegular() {
				if fileInfo.ModTime().After(fileModTimes[file]) {
					fileModTimes[file] = fileInfo.ModTime()
					if err := load(handle, file); err != nil {
						log.Error("failed to reload data from %v, got error %v\n", files, err)
					}
				}
			}
		}
	}
}

func load(handle func([]byte), files ...string) (err error) {
	for _, file := range files {
		log.Debugf("load data from: '%v'", file)
		data, err := os.ReadFile(file)
		if err != nil {
			return err
		}
		handle(data)
	}
	return err
}
