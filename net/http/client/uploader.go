/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var DefaultUploader = NewUploader()

type Uploader = Client

func NewUploader() *Uploader {
	return &Uploader{
		typ:           ClientTypeUpload,
		httpClient:    DefaultDownloadHttpClient,
		retryTimes:    3,
		retryInterval: time.Second,
		logger:        nil,
		logLevel:      LogLevelSilent,
	}
}

func (u *Uploader) UploadReq(url string) *UploadReq {
	return NewUploadReq(url).Uploader(u)
}

const (
	chunkSize = 1024 * 1024 // 每个分块的大小，这里是1MB
)

// uploadChunk 上传单个文件分块
func uploadChunk(url, paramName, filePath string, chunkNum int, chunkTotal int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// 跳到当前分块的起始位置
	_, err = file.Seek(int64(chunkNum)*chunkSize, io.SeekStart)
	if err != nil {
		return err
	}

	// 读取分块数据
	buffer := make([]byte, chunkSize)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return err
	}
	buffer = buffer[:n]

	// 创建HTTP请求
	body := bufPool.Get().(*bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(paramName, filepath.Base(filePath)+"."+strconv.Itoa(chunkNum))
	if err != nil {
		return err
	}
	_, err = part.Write(buffer)
	if err != nil {
		return err
	}

	// 添加分块信息
	_ = writer.WriteField("chunkNumber", strconv.Itoa(chunkNum))
	_ = writer.WriteField("chunkTotal", strconv.Itoa(chunkTotal))

	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status code: %d", resp.StatusCode)
	}

	return nil
}

// uploadChunk 上传单个文件分块
func uploadChunk2(ctx context.Context, url, filePath string, startByte, endByte int64) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(startByte, io.SeekStart)
	if err != nil {
		return err
	}

	requestBody := &bytes.Buffer{}
	reader := io.LimitReader(file, endByte-startByte+1)
	_, err = io.Copy(requestBody, reader)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, requestBody)
	if err != nil {
		return err
	}

	// 设置Range头部
	req.Header.Set("Content-Range", fmt.Sprintf("bytes %d-%d/*", startByte, endByte))
	req.Header.Set("Content-Length", strconv.FormatInt(endByte-startByte+1, 10))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusPartialContent && resp.StatusCode != http.StatusOK {
		return fmt.Errorf("upload failed with status code: %d", resp.StatusCode)
	}

	return nil
}

func uploadStream(url string, filePath string) error {
	// 打开文件以读取
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// 创建一个HTTP请求，使用文件的Reader作为Body
	req, err := http.NewRequestWithContext(context.Background(), "POST", url, file)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Add("Content-Type", "application/octet-stream") // 根据实际情况设置Content-Type
	req.ContentLength = -1                                     // 如果已知文件大小，可以设置准确的Content-Length

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to upload: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server responded with status %s: %s", resp.Status, body)
	}

	fmt.Println("File uploaded successfully.")
	return nil
}
