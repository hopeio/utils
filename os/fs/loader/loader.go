/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package loader

import (
	"github.com/fsnotify/fsnotify"
	"github.com/hopeio/utils/log"
	"os"
	"time"
)

type Loader struct {
	AutoReloadType ReloadType    `json:"autoReloadType" comment:"none,fsnotify,timer"` // 本地分为fsnotify和timer，fsnotify采用系统调用通知，timer定时器去查文件是否变更
	TimerInterval  time.Duration `json:"timerInterval" comment:"timer interval"`
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
	if ld.AutoReloadType != ReloadTypeNone {
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
