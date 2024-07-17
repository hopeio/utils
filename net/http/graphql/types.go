package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	stringsi "github.com/hopeio/utils/strings"
	"google.golang.org/protobuf/types/known/anypb"
	"log"

	"io"
	"strconv"
)

type DummyResolver struct{}

func (r *DummyResolver) Dummy(ctx context.Context) (*bool, error) { return nil, nil }

func MarshalBytes(b []byte) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = fmt.Fprintf(w, "%q", string(b))
	})
}

func UnmarshalBytes(v interface{}) ([]byte, error) {
	switch v := v.(type) {
	case string:
		return []byte(v), nil
	case *string:
		return []byte(*v), nil
	case []byte:
		return v, nil
	case json.RawMessage:
		return v, nil
	default:
		return nil, fmt.Errorf("%T is not []byte", v)
	}
}

func MarshalAny(any anypb.Any) Marshaler {
	return WriterFunc(func(w io.Writer) {
		d, err := any.UnmarshalNew()
		if err != nil {
			log.Println("unable to unmarshal any: ", err)
			return
		}

		if err := json.NewEncoder(w).Encode(d); err != nil {
			log.Println("unable to encode json: ", err)
		}
	})
}

func UnmarshalAny(v interface{}) (anypb.Any, error) {
	switch v := v.(type) {
	case []byte:
		return anypb.Any{}, nil //TODO add an unmarshal mechanism
	case json.RawMessage:
		return anypb.Any{}, nil
	default:
		return anypb.Any{}, fmt.Errorf("%T is not json.RawMessage", v)
	}
}

func MarshalInt(i int) Marshaler {
	return WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Itoa(i))
	})
}

func UnmarshalInt(v interface{}) (int, error) {
	switch v := v.(type) {
	case string:
		return strconv.Atoi(v)
	case int:
		return v, nil
	case int64:
		return int(v), nil
	case json.Number:
		return strconv.Atoi(string(v))
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

func MarshalInt8(any int8) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write(stringsi.ToBytes(strconv.Itoa(int(any))))
	})
}

func UnmarshalInt8(v interface{}) (int8, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseInt(v, 10, 8)
		return int8(iv), err
	case int:
		return int8(v), nil
	case int64:
		return int8(v), nil
	case json.Number:
		i, err := v.Int64()
		return int8(i), err
	default:
		return 0, fmt.Errorf("%T is not int8", v)
	}
}

func MarshalInt16(any int16) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write(stringsi.ToBytes(strconv.Itoa(int(any))))
	})
}

func UnmarshalInt16(v interface{}) (int16, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseInt(v, 10, 16)
		return int16(iv), err
	case int:
		return int16(v), nil
	case int64:
		return int16(v), nil
	case json.Number:
		i, err := v.Int64()
		return int16(i), err
	default:
		return 0, fmt.Errorf("%T is not int16", v)
	}
}

func MarshalInt32(i int32) Marshaler {
	return WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(int64(i), 10))
	})
}

func UnmarshalInt32(v interface{}) (int32, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(iv), nil
	case int:
		return int32(v), nil
	case int64:
		return int32(v), nil
	case json.Number:
		iv, err := strconv.ParseInt(string(v), 10, 32)
		if err != nil {
			return 0, err
		}
		return int32(iv), nil
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

func MarshalInt64(i int64) Marshaler {
	return WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.FormatInt(i, 10))
	})
}

func UnmarshalInt64(v interface{}) (int64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseInt(v, 10, 64)
	case int:
		return int64(v), nil
	case int64:
		return v, nil
	case json.Number:
		return strconv.ParseInt(string(v), 10, 64)
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

func MarshalUint(any uint) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write(stringsi.ToBytes(strconv.Itoa(int(any))))
	})
}

func UnmarshalUint(v interface{}) (uint, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseUint(v, 10, 64)
		return uint(iv), err
	case int:
		return uint(v), nil
	case int64:
		return uint(v), nil //TODO add an unmarshal mechanism
	case json.Number:
		i, err := v.Int64()
		return uint(i), err
	default:
		return 0, fmt.Errorf("%T is not uint", v)
	}
}

func MarshalUint8(any uint8) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write(stringsi.ToBytes(strconv.Itoa(int(any))))
	})
}

func UnmarshalUint8(v interface{}) (uint8, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseUint(v, 10, 8)
		return uint8(iv), err
	case int:
		return uint8(v), nil
	case int64:
		return uint8(v), nil //TODO add an unmarshal mechanism
	case json.Number:
		i, err := v.Int64()
		return uint8(i), err
	default:
		return 0, fmt.Errorf("%T is not uint64", v)
	}
}

func MarshalUint16(any uint16) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write(stringsi.ToBytes(strconv.Itoa(int(any))))
	})
}

func UnmarshalUint16(v interface{}) (uint16, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseUint(v, 10, 16)
		return uint16(iv), err
	case int:
		return uint16(v), nil
	case int64:
		return uint16(v), nil //TODO add an unmarshal mechanism
	case json.Number:
		i, err := v.Int64()
		return uint16(i), err
	default:
		return 0, fmt.Errorf("%T is not uint64", v)
	}
}

func MarshalUint32(any uint32) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write([]byte(strconv.Itoa(int(any))))
	})
}

func UnmarshalUint32(v interface{}) (uint32, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseUint(v, 10, 32)
		return uint32(iv), err
	case int:
		return uint32(v), nil
	case int64:
		return uint32(v), nil
	case json.Number:
		i, err := v.Int64()
		return uint32(i), err
	default:
		return 0, fmt.Errorf("%T is not int32", v)
	}
}

func MarshalUint64(i uint64) Marshaler {
	return WriterFunc(func(w io.Writer) {
		w.Write(stringsi.ToBytes(strconv.FormatUint(i, 10)))
	})
}

func UnmarshalUint64(v interface{}) (uint64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseUint(v, 10, 64)
	case int:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case json.Number:
		return strconv.ParseUint(string(v), 10, 64)
	default:
		return 0, fmt.Errorf("%T is not an int", v)
	}
}

func MarshalFloat32(any float32) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write(stringsi.ToBytes(strconv.Itoa(int(any))))
	})
}

func UnmarshalFloat32(v interface{}) (float32, error) {
	switch v := v.(type) {
	case string:
		iv, err := strconv.ParseFloat(v, 32)
		return float32(iv), err
	case float64:
		return float32(v), nil
	case json.Number:
		f, err := v.Float64()
		return float32(f), err
	default:
		return 0, fmt.Errorf("%T is not float32", v)
	}
}

func MarshalFloat64(any float32) Marshaler {
	return WriterFunc(func(w io.Writer) {
		_, _ = w.Write(stringsi.ToBytes(strconv.FormatFloat(float64(any), 'g', -1, 64)))
	})
}

func UnmarshalFloat64(v interface{}) (float64, error) {
	switch v := v.(type) {
	case string:
		return strconv.ParseFloat(v, 64)
	case float64:
		return v, nil
	case json.Number:
		f, err := v.Float64()
		return f, err
	default:
		return 0, fmt.Errorf("%T is not float32", v)
	}
}

type Marshaler interface {
	MarshalGQL(w io.Writer)
}

type WriterFunc func(writer io.Writer)

func (f WriterFunc) MarshalGQL(w io.Writer) {
	f(w)
}

func MarshalHeader(val map[string]string) Marshaler {
	return WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalHeader(v interface{}) (map[string]string, error) {
	if m, ok := v.(map[string]string); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}

func MarshalMap(val map[string]interface{}) Marshaler {
	return WriterFunc(func(w io.Writer) {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			panic(err)
		}
	})
}

func UnmarshalMap(v interface{}) (map[string]interface{}, error) {
	if m, ok := v.(map[string]interface{}); ok {
		return m, nil
	}

	return nil, fmt.Errorf("%T is not a map", v)
}
