package time

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
	"unsafe"
)

func TestTimeStamp(t *testing.T) {
	var t1 = Timestamp(time.Now().UnixMilli())
	var t2 = SecondTimestamp(time.Now().Unix())
	var t3 = NanoTimestamp(time.Now().UnixNano())
	var t4 = MicroTimestamp(time.Now().UnixMicro())
	data, _ := t1.MarshalJSON()
	t.Log(string(data))
	data, _ = t2.MarshalJSON()
	t.Log(string(data))
	data, _ = t3.MarshalJSON()
	t.Log(string(data))
	data, _ = t4.MarshalJSON()
	t.Log(string(data))
}

func TestTimeAdd(t *testing.T) {
	log.Println(time.Now().AddDate(0, 0, -16))
}

func TestSizeof(t *testing.T) {
	var ts Timestamp
	assert.Equal(t, 8, unsafe.Sizeof(ts))
}

func TestTimeStampScan(t *testing.T) {
	tm := time.Now()
	t.Log(tm)
	ts := tm.Unix()
	t.Log(ts)
	t.Log(time.Unix(ts, 0))
	t.Log(time.Unix(ts, 0).Local())
}
