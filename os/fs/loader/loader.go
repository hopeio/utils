/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package loader

import (
	"github.com/fsnotify/fsnotify"
	"github.com/hopeio/utils/log"
	"io"
	"os"
	"time"
)

type Loader struct {
	// 间隔大于1秒采用timer定时加载，小于1秒用fsnotify
	AutoReloadInterval time.Duration `json:"autoReloadInterval" comment:"none"`
}

type ReloadType int

const (
	ReloadTypeNone ReloadType = iota
	ReloadTypeFsNotify
	ReloadTypeTimer
)

const (
	ReloadTypeNoneName     = "none"
	ReloadTypeFsNotifyName = "fsnotify"
	ReloadTypeTimerName    = "timer"
)

func (t ReloadType) MarshalText() ([]byte, error) {
	return []byte(t.String()), nil
}
func (t *ReloadType) UnmarshalText(data []byte) error {
	switch string(data) {
	case ReloadTypeNoneName:
		*t = ReloadTypeNone
	case ReloadTypeFsNotifyName:
		*t = ReloadTypeFsNotify
	case ReloadTypeTimerName:
		*t = ReloadTypeTimer
	default:
		*t = ReloadTypeNone
	}
	return nil
}

func (t *ReloadType) UnmarshalJSON(data []byte) error {
	data = data[1 : len(data)-1]
	return t.UnmarshalText(data)
}

func (t ReloadType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

func (t ReloadType) String() string {
	switch t {
	case ReloadTypeNone:
		return ReloadTypeNoneName
	case ReloadTypeFsNotify:
		return ReloadTypeFsNotifyName
	case ReloadTypeTimer:
		return ReloadTypeTimerName
	default:
		return ReloadTypeNoneName
	}
}

// New initialize a Loader
func New(interval time.Duration) *Loader {
	return &Loader{AutoReloadInterval: interval}
}

// Load will unmarshal configurations to struct from files that you provide
func (ld *Loader) Handle(handle func(io.Reader), filepaths ...string) (err error) {

	err = load(handle, filepaths...)
	if err != nil {
		return err
	}
	if ld.AutoReloadInterval != 0 {
		if ld.AutoReloadInterval >= time.Second {
			go ld.watchTimer(handle, filepaths...)
		} else {
			go ld.watchNotify(handle, filepaths...)
		}
	}

	return
}

func (ld *Loader) watchNotify(handle func(reader io.Reader), filepaths ...string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error(err)
	}
	defer watcher.Close()
	for _, filepath := range filepaths {
		err = watcher.Add(filepath)
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
					log.Errorf("failed to reload data from %v, got error %v\n", filepaths, err)
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

func (ld *Loader) watchTimer(handle func(reader io.Reader), files ...string) {
	var fileModTimes map[string]time.Time
	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]

		// check configuration
		if fileInfo, err := os.Stat(file); err == nil && fileInfo.Mode().IsRegular() {
			fileModTimes[file] = fileInfo.ModTime()
		}
	}
	timer := time.NewTicker(ld.AutoReloadInterval)
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

func load(handle func(io.Reader), filepaths ...string) (err error) {
	for _, filepath := range filepaths {
		log.Debugf("load data from: '%v'", filepath)
		file, err := os.Open(filepath)
		if err != nil {
			return err
		}
		handle(file)
		file.Close()
	}
	return err
}
