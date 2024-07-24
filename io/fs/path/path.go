package path

import (
	stringsi "github.com/hopeio/utils/strings"
	sdpath "path"
	"path/filepath"
	"strings"
)

// windows需要,由于linux的文件也要放到windows看,统一处理
func FileRewrite(filename string) string {
	filename = stringsi.ReplaceRunesEmpty(filename, '/', '\\', '*', '|')
	filename = strings.ReplaceAll(filename, "<", "《")
	filename = strings.ReplaceAll(filename, ">", "》")
	filename = strings.ReplaceAll(filename, "?", "？")
	filename = strings.ReplaceAll(filename, ":", "：")
	return filename
}

// 仅仅针对文件名,Removed unsupported characters
func FileCleanse(filename string) string {

	filename = strings.Trim(filename, ".")
	filename = strings.TrimPrefix(filename, "-")
	filename = strings.TrimPrefix(filename, "+")
	// windows
	//filename = stringsi.ReplaceRunesEmpty(filename, '/', '\\', ':', '*', '?', '"', '<', '>', '|')
	// linux
	//filename = stringsi.ReplaceRunesEmpty(filename, '\'', '*','?', '@', '#', '$', '&', '(', ')', '|', ';',  '/', '%', '^', ' ', '\t', '\n')

	filename = stringsi.ReplaceRunesEmpty(filename, '/', '\\', ':', '*', '?', '"', '<', '>', '|', ';', '/', '%', '^', ' ', '\t', '\n', '$', '&')
	// 中文符号
	//filename = stringsi.ReplaceRunesEmpty(filename, '：', '，', '。', '！', '？', '、', '“', '”', '、')
	return filename
}

// 仅仅针对目录名,Removed unsupported characters
func DirCleanse(dir string) string { // will be used when save the dir or the part
	// remove special symbol
	// :unix允许存在，windows需要
	// windows path
	if len(dir) > 2 && dir[1] == ':' && ((dir[0] >= 'A' && dir[0] <= 'Z') || (dir[0] >= 'a' && dir[0] <= 'z')) && (dir[2] == '/' || dir[2] == '\\') {
		return dir[:3] + stringsi.ReplaceRunesEmpty(dir[3:], ':', '*', '?', '"', '<', '>', '|', ',', ' ', '\t', '\n')
	}
	return stringsi.ReplaceRunesEmpty(dir, ':', '*', '?', '"', '<', '>', '|', ',', ' ', '\t', '\n')
}

// 针对带目录的完整文件名,Removed unsupported characters
func Cleanse(path string) string { // will be used when save the dir or the part
	dir, file := filepath.Split(path)
	if dir == "" {
		return DirCleanse(dir)
	}
	if file == "" {
		return FileCleanse(file)
	}
	// remove special symbol
	return DirCleanse(dir) + string(path[len(dir)-1-len(file)]) + FileCleanse(file)
}

// 获取文件名除去扩展名
func FileNoExt(filepath string) string {
	base := sdpath.Base(filepath)
	return base[:len(base)-len(sdpath.Ext(base))]
}
