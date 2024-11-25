package gerber

import (
	"os"
	"testing"
)

func TestGerber(t *testing.T) {
	path := `D:\博瀚智能（深圳）有限公司\Kelly Zhang - Apulis File Share\CTO\内部项目\勃朗峰\技术实现\测试文件\客户1样例文件\SB047A5W230643-F7-FG-SB3-MLB-TOP-DVT-A001  G651-08585-04.GBR`
	p := LogProcessor{}

	f, _ := os.Open(path)
	defer f.Close()
	err := NewParser(p).Parse(f)
	if err != nil {
		t.Error(err)
	}
}
