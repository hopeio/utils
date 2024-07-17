package local

import (
	"github.com/hopeio/utils/log"

	"os"
	"time"
)

func (configor *Configor) getConfigurationFiles(files ...string) ([]string, map[string]time.Time) {
	var resultKeys []string
	var results = map[string]time.Time{}

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]

		// check configuration
		if fileInfo, err := os.Stat(file); err == nil && fileInfo.Mode().IsRegular() {
			resultKeys = append(resultKeys, file)
			results[file] = fileInfo.ModTime()
		}
	}
	return resultKeys, results
}

func (configor *Configor) handle(handle func([]byte), watchMode bool, files ...string) (err error, changed bool) {
	defer func() {
		if err != nil {
			log.Errorf("Failed to load configuration from %v, got %v\n", files, err)
		}
	}()

	configFiles, configModTimeMap := configor.getConfigurationFiles(files...)

	if watchMode {
		if len(configModTimeMap) == len(configor.configModTimes) {
			var changed bool
			for f, t := range configModTimeMap {
				if v, ok := configor.configModTimes[f]; !ok || t.After(v) {
					changed = true
				}
			}

			if !changed {
				return nil, false
			}
		}
	}

	for _, file := range configFiles {
		log.Infof("load config from: '%v'", file)
		data, err := os.ReadFile(file)
		if err != nil {
			return err, true
		}
		handle(data)
	}
	configor.configModTimes = configModTimeMap

	return err, true
}
