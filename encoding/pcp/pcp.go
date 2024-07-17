package pcp

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type PCP struct {
	Header Header
	Data   []PointXYZGray // [][][]int
}

type Header struct {
	Type             string
	Size             int32
	Width            int32
	Height           int32
	Channel          int8
	IntervalValid    int8
	Minx, Miny, Minz int32
	Maxx, Maxy, Maxz int32
	Dx, Dy           int32
	Expose           int
	Timestamp        int
	DeviceMessage    string
	decimal          uint8
}

type PointXYZGray struct {
	X, Y, Z int32
	Gray    int32
}

func Parse(filePath string) (*PCP, error) {
	var pcp PCP
	m := toMap(&pcp.Header)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("open file error:%w", err)
	}
	var isHeader bool
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "header" {
			isHeader = true
			continue
		}
		if str == "data" {
			isHeader = false
			pcp.Data = make([]PointXYZGray, 0, pcp.Header.Size)
			continue
		}
		if isHeader {
			kv := strings.Split(str, ":")
			if len(kv) != 2 {
				return nil, fmt.Errorf("parse header error:%s", str)
			}
			if kv[1] == "(null)" {
				continue
			}
			if v, ok := m[strings.ToUpper(strings.ReplaceAll(kv[0], "_", ""))]; ok {
				switch v.Kind() {
				case reflect.Int32, reflect.Int, reflect.Int8:
					num := kv[1]
					if strings.Contains(kv[1], ".") {
						num = trimZero(strings.ReplaceAll(kv[1], ".", ""))
						if pcp.Header.decimal == 0 {
							pcp.Header.decimal = uint8(len(kv[1][strings.Index(kv[1], "."):]))
						}
					}

					istr, err := strconv.Atoi(num)
					if err != nil {
						return nil, err
					}
					v.SetInt(int64(istr))

				case reflect.String:
					v.SetString(kv[1])
				}
			}
		} else {
			values := strings.Split(str, " ")
			if len(values) != 4 {
				return nil, fmt.Errorf("parse data error:%s", str)
			}

			x, err := strconv.Atoi(trimZero(strings.ReplaceAll(values[0], ".", "")))
			if err != nil {
				return nil, err
			}
			y, err := strconv.Atoi(trimZero(strings.ReplaceAll(values[1], ".", "")))
			if err != nil {
				return nil, err
			}
			z, err := strconv.Atoi(trimZero(strings.ReplaceAll(values[2], ".", "")))
			if err != nil {
				return nil, err
			}
			gary, err := strconv.Atoi(trimZero(strings.ReplaceAll(values[3], ".", "")))
			if err != nil {
				return nil, err
			}

			pcp.Data = append(pcp.Data, PointXYZGray{X: int32(x), Y: int32(y), Z: int32(z), Gray: int32(gary)})

		}
	}
	return &pcp, nil
}

func toMap(v any) map[string]reflect.Value {
	m := make(map[string]reflect.Value)
	value := reflect.ValueOf(v).Elem()
	for i := 0; i < value.NumField(); i++ {
		m[strings.ToUpper(value.Type().Field(i).Name)] = value.Field(i)
	}
	return m
}

func trimZero(str string) string {
	for i := 0; i < len(str); i++ {
		if str[i] != '0' {
			return str[i:]
		}
	}
	return "0"
}
