/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/grafana/loki-client-go/loki"
	"github.com/hopeio/utils/log"
	"github.com/prometheus/common/model"
	"os"
	"time"
)

type LogEntry struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"message"`
	TraceId string `json:"traceId"`
	Caller  string `json:"caller"`
}

func main() {
	// 初始化 Loki 客户端
	config, _ := loki.NewDefaultConfig("http://host/loki/api/v1/push")
	config.BatchWait = time.Second
	client, err := loki.New(config)
	if err != nil {
		log.Fatalf("Failed to create Loki client: %v", err)
	}

	file, err := os.Open("xxx.log")
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var entry LogEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err != nil {
			log.Fatalf("Failed to unmarshal JSON: %v", err)
		}

		t, err := time.Parse(time.DateTime, entry.Time)
		if err != nil {
			log.Fatalf("Failed to parse timestamp: %v", err)
		}

		err = client.Handle(model.LabelSet{
			"level":   model.LabelValue(entry.Level),
			"traceId": model.LabelValue(entry.TraceId),
			"caller":  model.LabelValue(entry.Caller),
		}, t, entry.Message)
		if err != nil {
			log.Error(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading log file: %v", err)
	}

	// 确保所有日志都已发送
	client.Stop()
	fmt.Println("Logs sent to Loki successfully")

}
