/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package apollo

import (
	"testing"

	"github.com/hopeio/utils/log"
)

func TestApollo(t *testing.T) {
	s := Client{
		Addr:           "192.168.1.212:8080",
		AppId:          "hoper",
		Cluster:        "default",
		IP:             "",
		Configurations: map[string]*Configuration{},
		Notifications:  map[string]int{"application": -1},
	}
	_, err := s.GetCacheConfig("application")
	if err != nil {
		log.Error(err)
	}
	log.Debug(s.Configurations["application"])
	_, err = s.GetNoCacheConfig("application")
	if err != nil {
		log.Error(err)
	}
	log.Info(s.Configurations["application"])
}

func TestApolloPoll(t *testing.T) {
	s := Client{
		Addr:           "192.168.1.212:8080",
		AppId:          "hoper",
		Cluster:        "default",
		IP:             "",
		Configurations: map[string]*Configuration{},
		Notifications:  map[string]int{"application": -1},
	}
	err := s.UpdateConfig("hoper", "default")
	if err != nil {
		log.Error(err)
	}
}

func TestApolloConfig_Generate(t *testing.T) {
	type Config struct {
		A int    `json:"a"`
		B string `json:"b"`
	}
	type Foo struct {
		Bar Config
	}
	var foo Foo
	var conf = map[string]string{"baR": "A = 1\nB = \"哈哈哈\""}
	apolloConfigEnable(&foo, conf)
	log.Info(foo)
}
