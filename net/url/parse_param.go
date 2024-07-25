package url

import (
	"github.com/hopeio/utils/math"
	stringsi "github.com/hopeio/utils/strings"
	"net/url"
	"path"
	"reflect"
	"strconv"
	"strings"
)

var tag = "json"

func SetTag(t string) {
	tag = t
}

func ResolveURL(u *url.URL, p string) string {
	if strings.HasPrefix(p, "https://") || strings.HasPrefix(p, "http://") {
		return p
	}
	var baseURL string
	if strings.Index(p, "/") == 0 {
		baseURL = u.Scheme + "://" + u.Host
	} else {
		tU := u.String()
		baseURL = tU[0:strings.LastIndex(tU, "/")]
	}
	return baseURL + path.Join("/", p)
}

func QueryParam(param any) string {
	return QueryParamByTag(param, tag)
}

func QueryParamByTag(param any, tag string) string {
	if param == nil {
		return ""
	}
	query := url.Values{}
	parseParamByTag(param, query, tag)
	return query.Encode()
}

func parseParamByTag(param any, query url.Values, tag string) {
	v := reflect.ValueOf(param)
	kind := v.Kind()
	if kind == reflect.Interface || kind == reflect.Ptr {
		if v.IsNil() {
			return
		}
		v = v.Elem()
	}
	if v.Kind() == reflect.Struct {
		t := v.Type()
		for i := 0; i < v.NumField(); i++ {
			filed := v.Field(i)
			fieldKind := filed.Kind()
			if fieldKind == reflect.Interface || fieldKind == reflect.Ptr || fieldKind == reflect.Struct {
				if t.Field(i).Anonymous {
					parseParamByTag(filed.Interface(), query, tag)
				} else {
					//TODO: 处理对象
				}
				continue
			}
			if fieldKind == reflect.Map {
				//TODO: 处理map
				continue
			}
			if fieldKind == reflect.Slice || fieldKind == reflect.Array {
				for i := 0; i < filed.Len(); i++ {
					query.Add(t.Field(i).Tag.Get(tag), getFieldValue(filed.Index(i)))
				}
				continue
			}

			value := getFieldValue(filed)
			if value != "" {
				query.Set(t.Field(i).Tag.Get(tag), getFieldValue(v.Field(i)))
			}
		}
	}
	if kind == reflect.Map {
		if v.IsNil() {
			return
		}
		for _, key := range v.MapKeys() {
			query.Set(key.Interface().(string), getFieldValue(v.MapIndex(key)))
		}
	}
}

func getFieldValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.Itoa(int(v.Int()))
	case reflect.Float32, reflect.Float64:
		return math.FormatFloat(v.Float())
	case reflect.String:
		return v.String()
	case reflect.Interface, reflect.Ptr, reflect.Struct:
		panic("unsupported kind " + v.Kind().String())
	}
	return ""
}

func AppendQueryParamByTag(url string, param interface{}, tag string) string {
	if param == nil {
		return url
	}
	sep := "?"
	if strings.Contains(url, sep) {
		sep = "&"
	}
	switch paramt := param.(type) {
	case string:
		url += sep + paramt
	case []byte:
		url += sep + stringsi.BytesToString(paramt)
	default:
		params := QueryParamByTag(param, tag)
		url += sep + params
	}
	return url
}

func AppendQueryParam(url string, param interface{}) string {
	return AppendQueryParamByTag(url, param, tag)
}
