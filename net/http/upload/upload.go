package upload

import (
	"fmt"
	httpi "github.com/hopeio/utils/net/http"
	"go.uber.org/atomic"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	fileMap     = sync.Map{}
	fileMapSize = atomic.NewInt64(0)
)

type fileInfo struct {
	*os.File
	size           int64
	lastAccessTime time.Time
}

func Upload(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		var file *fileInfo
		var err error
		if fileValue, exists := fileMap.Load(filePath); !exists {
			file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
			if err != nil {
				http.Error(w, "Failed to open file", http.StatusInternalServerError)
				return
			}
			info, err := file.Stat()
			if err != nil {
				http.Error(w, "Failed to get file info", http.StatusInternalServerError)
				return
			}
			now := time.Now()
			fileMap.Store(filePath, &fileInfo{File: file, size: info.Size(), lastAccessTime: now})
			fileMapSize.Add(1)
			if fileMapSize.Load() > 100 {
				fileMap.Range(func(key, value any) bool {
					if value.(*fileInfo).lastAccessTime.Add(time.Minute * 5).Before(now) {
						fileMap.Delete(key)
					}
					return true
				})
			}
		} else {
			file = fileValue.(*fileInfo)
		}
		if r.Method != http.MethodPost && r.Method != http.MethodPut {
			// 如果一切顺利，发送成功的响应
			w.Header().Set(httpi.HeaderContentRange, fmt.Sprintf("bytes 0-0/%d", file.size))
			w.WriteHeader(http.StatusPartialContent)
			w.Write([]byte("success"))
			return
		}

		// 解析Range头部
		rangeHeader := r.Header.Get(httpi.HeaderContentRange)
		if rangeHeader == "" {
			http.Error(w, "Missing Content-Range header", http.StatusBadRequest)
			return
		}

		// 提取Range值，格式为"bytes unit-unit/*"
		parts := strings.Split(rangeHeader, " ")
		if len(parts) != 2 || parts[0] != "bytes" {
			http.Error(w, "Invalid Content-Range format", http.StatusBadRequest)
			return
		}
		rangeSpec := parts[1]
		info := strings.Split(rangeSpec, "/")
		bounds := strings.Split(info[0], "-")
		start, err := strconv.ParseInt(bounds[0], 10, 64)
		if err != nil {
			http.Error(w, "Invalid range start", http.StatusBadRequest)
			return
		}
		var end int64
		if len(bounds) > 1 {
			end, err = strconv.ParseInt(bounds[1], 10, 64)
			if err != nil {
				http.Error(w, "Invalid range end", http.StatusBadRequest)
				return
			}
		} else {
			// 如果只有开始位置，结束位置默认为文件末尾
			end = -1
		}

		// 打开文件准备写入，使用O_RDWR | O_CREATE | O_APPEND以追加模式打开

		// 移动文件指针到开始位置
		_, err = file.Seek(start, io.SeekStart)
		if err != nil {
			http.Error(w, "Failed to seek file", http.StatusInternalServerError)
			return
		}

		// 读取请求体并写入文件
		_, err = io.CopyN(file, r.Body, end-start+1)
		if err != nil && err != io.EOF {
			http.Error(w, "Failed to write to file", http.StatusInternalServerError)
			return
		}

		file.size = file.size + end - start + 1
		file.lastAccessTime = time.Now()

		if err == io.EOF {
			file.Close()
		} else {
			if len(info) == 2 && info[1] != "*" {
				var total int64
				total, err = strconv.ParseInt(info[1], 10, 64)
				if err != nil {
					http.Error(w, "Invalid range size", http.StatusBadRequest)
					return
				}
				if file.size == total {
					http.Error(w, "Invalid range size", http.StatusBadRequest)
					return
				}
			}
		}

		// 如果一切顺利，发送成功的响应
		w.Header().Set(httpi.HeaderContentRange, fmt.Sprintf("bytes %d-%d/%d", start, end, file.size))
		w.WriteHeader(http.StatusPartialContent)
		fmt.Fprintf(w, "Uploaded chunk from byte %d to %d", start, end)
	}

}
