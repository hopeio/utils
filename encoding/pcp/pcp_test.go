package pcp

import "testing"

func TestParse(t *testing.T) {
	pcp, err := Parse("./sdk_save_cloud_csharp.pcp")
	if err != nil {
		t.Error(err)
	}
	if pcp.Data == nil {
		t.Error("pcp.Data is nil")
	}
	if pcp.Header.Type != "POINTXYZGRAY" {
		t.Error("pcp.Header.Type is not PointXYZGray")
	}
	if pcp.Header.IntervalValid != 1 {
		t.Error("pcp.Header.IntervalValid is not 1")
	}
	if int(pcp.Header.Size) != len(pcp.Data) {
		t.Error("pcp.Header.Size is not equal to len(pcp.Data)")
	}
}
