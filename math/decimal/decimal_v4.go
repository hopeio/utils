/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package decimal

import (
	"database/sql/driver"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

type DecimalV4 struct {
	mant uint64
	exp  int
	neg  bool
}

func New3(mant uint64, exp int, neg bool) *DecimalV4 {
	if mant == 0 {
		return &DecimalV4{}
	}
	mantStr := strconv.FormatUint(mant, 10)
	for i := len(mantStr) - 1; i >= 0; i-- {
		if mantStr[i] == '0' {
			mant /= 10
			exp += 1
		} else {
			break
		}
	}
	return &DecimalV4{
		mant: mant,
		exp:  exp,
		neg:  neg,
	}
}

func New3FromStr(str string) (DecimalV4, error) {
	var dec DecimalV4
	if str != "" && str[0] == '-' {
		dec.neg = true
		str = str[1:]
	}

	nums := strings.Split(str, ".")
	if len(nums) == 2 {
		for i := len(nums[1]) - 1; i >= 0; i-- {
			if nums[1][i] == '0' {
				nums[1] = nums[1][:i]
			} else {
				break
			}
		}
		dec.exp = -len(nums[1])
		str = nums[0] + nums[1]
	} else {
		str = nums[0]
	}

	dec.mant, _ = strconv.ParseUint(str, 10, 64)

	return dec, nil
}

func (x *DecimalV4) Add(v DecimalV4) *DecimalV4 {
	var dec = *x

	if x.exp > v.exp {
		dec.exp = v.exp
		dec.mant = dec.mant * uint64(math.Pow10(x.exp-v.exp))
	} else if x.exp < v.exp {
		v.mant = v.mant * uint64(math.Pow10(v.exp-x.exp))
		v.exp = x.exp
	}

	if x.neg == v.neg {
		dec.mant += v.mant
	} else {
		if x.mant >= v.mant {
			dec.mant -= v.mant
		} else {
			dec.mant = v.mant - dec.mant
			dec.neg = v.neg
		}
	}

	return &dec
}

func (x *DecimalV4) Sub(v DecimalV4) *DecimalV4 {
	v.neg = !v.neg
	return x.Add(v)
}

func (x *DecimalV4) Mul(v DecimalV4) *DecimalV4 {
	if x.mant == 0 || v.mant == 0 {
		return &DecimalV4{}
	}
	v.mant *= x.mant
	v.exp += x.exp
	if x.neg != v.neg {
		v.neg = true
	}
	return &v
}

type RoundingMode uint8

const (
	UP RoundingMode = iota
	DOWN
	HALFUP
)

func (x *DecimalV4) Div(v DecimalV4, mode RoundingMode) *DecimalV4 {
	if x.mant == 0 {
		return &DecimalV4{}
	}
	if v.mant == 0 {
		panic("除数不能为0")
	}
	if v.exp == 0 {
		return &*x
	}
	d1 := x.mant
	d2 := v.mant

	if v.exp < 0 {
		d1 = x.mant * uint64(math.Pow10(0-v.exp))
	} else {
		v.exp = x.exp - v.exp
	}

	v.mant = d1 / d2
	if x.neg != v.neg {
		v.neg = true
	}
	return &v
}

func (x *DecimalV4) DivInt(v int, mode RoundingMode) *DecimalV4 {
	dec := *x
	dec.mant = dec.mant / uint64(v)
	return &dec
}

func (x *DecimalV4) Float() float64 {
	if x.neg {
		return -float64(x.mant) / math.Pow10(abs(x.exp))
	}
	return float64(x.mant) / math.Pow10(abs(x.exp))
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

/*
	func (x *DecimalV4) Div2(v DecimalV4) *DecimalV4 {
		f1:=big.NewFloat(x.Float())
		f2:=big.NewFloat(v.Float())
		f3,_:=f1.Quo(f1,f2).Float64()
	}
*/
func (x DecimalV4) String() string {
	if x.mant == 0 {
		return "0"
	}
	d := x.mant
	if x.exp > 0 {
		d = x.mant * uint64(math.Pow10(x.exp))
	}

	str := strconv.FormatUint(d, 10)
	var in, de string
	if x.neg == true {
		in = "-"
	}
	if len(str)+x.exp < 0 {
		in += "0"
		de = "." + strings.Repeat("0", -x.exp-len(str)) + str
	} else {
		if x.exp >= 0 {
			in += str
		} else {
			in += str[:len(str)+x.exp]
			de = "." + str[len(str)+x.exp:]
		}

	}

	return fmt.Sprintf("%s%s", in, de)
}

func (x *DecimalV4) Decompose(buf []byte) (form byte, negative bool, mant []byte, exponent int32) {
	return 0, x.neg, (*(*[8]byte)(unsafe.Pointer(&x.mant)))[:], int32(x.exp)
}

func (x *DecimalV4) Compose(form byte, negative bool, mant []byte, exponent int32) error {
	x.neg = negative
	x.exp = int(exponent)
	var array [8]byte
	copy(array[:], mant)
	x.mant = *(*uint64)(unsafe.Pointer(&array))
	return nil
}

func (x *DecimalV4) Scan(raw interface{}) error {
	var err error
	switch v := raw.(type) {
	case []byte:
		*x, err = New3FromStr(string(v))
	case string:
		*x, err = New3FromStr(v)
	default:
		return fmt.Errorf("cannot sql.Scan() strfmt.Base64 from: %#v", v)
	}
	if err != nil {
		return err
	}
	return nil
}

func (x DecimalV4) Value() (driver.Value, error) {
	return driver.Value(x.String()), nil
}
