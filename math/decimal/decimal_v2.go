/*
 * Copyright 2024 hopeio. All rights reserved.
 * Licensed under the MIT License that can be found in the LICENSE file.
 * @Created by jyb
 */

package decimal

import (
	"errors"
	"fmt"
	"github.com/hopeio/utils/log"
	"strconv"
	"strings"
)

// 放弃，省空间，但计算时间浪费，来回不停转字符串
type DecimalV2 struct {
	neg bool
	Int uint64
	//小数部分翻转 0.001 =》 100
	dec uint64
}

//居然没有运算符重载

func New1(Dec string) (dec DecimalV2, err error) {
	nums := strings.Split(Dec, ".")
	dec.Int, err = strconv.ParseUint(nums[0], 10, 64)
	if len(nums) == 1 {
		if err != nil {
			return
		}
		return
	}

	if len(nums[1]) > 19 {
		err = errors.New("小数最多19位")
		log.Error(err)
		return
	}

	err = dec.SetDec(nums[1])
	if err != nil {
		return
	}
	return dec, nil
}

func reverse(s string) string {
	bytes := make([]byte, 0, len(s))
	zeroFlag := true
	for i := len(s) - 1; i > 0; i++ {
		// 去掉末位的零
		if !zeroFlag {
			bytes = append(bytes, s[i])
		} else {
			if s[i] == '0' {
				continue
			} else {
				zeroFlag = false
				bytes = append(bytes, s[i])
			}
		}
	}
	return string(bytes)
}

func (d *DecimalV2) Dec() string {
	dec := strconv.FormatUint(d.dec, 10)
	return reverse(dec)
}

func (d *DecimalV2) DecInt() uint64 {
	dec, _ := strconv.ParseUint(d.Dec(), 10, 64)
	return dec
}

func (d *DecimalV2) SetDec(dec string) error {
	var err error
	dec = reverse(dec)
	d.dec, err = strconv.ParseUint(dec, 10, 64)
	return err
}

func (d *DecimalV2) SetDecInt(dec uint64) error {
	var err error
	decStr := strconv.FormatUint(dec, 10)
	decStr = reverse(decStr)
	d.dec, err = strconv.ParseUint(decStr, 10, 64)
	return err
}

func (d *DecimalV2) Add(v DecimalV2) DecimalV2 {
	var dec DecimalV2
	dec.Int = d.Int + v.Int
	d1 := d.Dec()
	d2 := v.Dec()
	if i := len(d1) - len(d2); i > 0 {
		d2 = d2 + strings.Repeat("0", i)
	} else {
		d1 = d1 + strings.Repeat("0", -i)
	}
	decStr := strconv.FormatUint(d.DecInt()+v.DecInt(), 10)

	if len(decStr)-len(d.Dec()) > 0 {
		dec.SetDec(decStr[1:])
		dec.Int += 1
	} else {
		dec.SetDec(decStr)
	}

	return dec
}

func (d *DecimalV2) Multi(v DecimalV2) {
	/*	i := d.Int * v.Int
		Decimal := d.Decimal * v.Int
		Decimal = Decimal + d.Int*v.Decimal + (d.Decimal*d.Decimal)/(int(exponent(10, uint64(d.accuracy*2))))
		i = i + Decimal/(int(exponent(10, uint64(d.accuracy))))
		d.Int = i
		d.Decimal = Decimal % (int(exponent(10, uint64(d.accuracy))))*/
}

func (d DecimalV2) String() string {
	return fmt.Sprintf("%d.%s", d.Int, d.Dec())
}

func exponent(a, n uint64) uint64 {
	result := uint64(1)
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			result *= a
		}
		a *= a
	}
	return result
}

type Decimal11 struct {
	Int      uint64
	dec      uint64
	accuracy int
}

func New11(Dec string, eff int) (dec *Decimal11, err error) {
	dec = new(Decimal11)
	nums := strings.Split(Dec, ".")
	dec.Int, err = strconv.ParseUint(nums[0], 10, 64)
	if len(nums) == 1 {
		if err != nil {
			return
		}
		return
	}

	if eff > 19 || len(nums[1]) > 19 {
		err = errors.New("小数最多19位")
		log.Error(err)
		return
	}
	dec.accuracy = eff
	if len(nums[1]) >= eff {
		nums[1] = nums[1][0:eff]
	} else {
		nums[1] = nums[1] + strings.Repeat("0", eff-len(nums[1]))
	}
	dec.dec, err = strconv.ParseUint(nums[1], 10, 64)
	return
}

func (x *Decimal11) String() string {
	dec := strconv.FormatUint(x.dec, 10)
	dec = dec + strings.Repeat("0", x.accuracy-len(dec))
	return fmt.Sprintf("%d.%s", x.Int, dec)
}

func (x *Decimal11) Add(v *Decimal11) *Decimal11 {
	var dec = *x
	dec.Int += v.Int
	if x.accuracy > v.accuracy {
		dec.accuracy = v.accuracy
		dec.dec = x.dec / exponent(uint64(x.accuracy-v.accuracy), 10)
	} else if x.accuracy < v.accuracy {
		dec.dec = x.dec / exponent(uint64(v.accuracy-x.accuracy), 10)
	}
	d := dec.dec + v.dec
	dStr := strconv.FormatUint(d, 10)

	if len(dStr) > x.accuracy {
		dec.dec, _ = strconv.ParseUint(dStr[1:], 10, 64)
		dec.Int += 1
	} else {
		dec.dec = d
	}
	return &dec
}
