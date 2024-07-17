package local

import (
	"github.com/fsnotify/fsnotify"
	"github.com/hopeio/utils/log"
	"time"
)

type Configor struct {
	*Config
	configModTimes map[string]time.Time
}

type ReloadType string

const (
	ReloadTypeFsNotify = "fsnotify"
	ReloadTypeTimer    = "timer"
)

type Config struct {
	AutoReload         bool
	AutoReloadType     ReloadType `json:"autoReloadType" comment:"fsnotify,timer"` // 本地分为Watch和AutoReload，Watch采用系统调用通知，AutoReload定时器去查文件是否变更
	AutoReloadInterval time.Duration
	AutoReloadCallback func(config interface{})
}

// New initialize a Configor
func New(config *Config) *Configor {
	if config == nil {
		config = &Config{}
	}
	if config.AutoReload && config.AutoReloadType == ReloadTypeTimer && config.AutoReloadInterval < time.Second {
		config.AutoReloadInterval = config.AutoReloadInterval * time.Second
	}
	return &Configor{Config: config}
}

// Load will unmarshal configurations to struct from files that you provide
func (configor *Configor) Handle(handle func([]byte), files ...string) (err error) {

	err, _ = configor.handle(handle, false, files...)
	if configor.AutoReload {
		if configor.AutoReloadType == ReloadTypeTimer {
			go func() {
				timer := time.NewTimer(configor.Config.AutoReloadInterval)
				for range timer.C {
					var changed bool
					if err, changed = configor.handle(handle, true, files...); err == nil && changed {
					} else if err != nil {
						log.Error("Failed to reload configuration from %v, got error %v\n", files, err)
					}
					timer.Reset(configor.Config.AutoReloadInterval)
				}
			}()
		} else {
			go configor.watch(handle)
		}
	}

	return
}

func (cc *Configor) watch(handle func([]byte), files ...string) {
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
				if err, changed := cc.handle(handle, true, event.Name); err == nil && changed {
				} else if err != nil {
					log.Error("Failed to reload configuration from %v, got error %v\n", files, err)
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
