package math

import (
	"math"
	"time"
)

const magicNumber = 0xf1234fff

//一个数异或同一个数两次还是这个数...

func SecondKey() int64 {
	return time.Now().Unix() ^ magicNumber
}

func ValidateSecondKey(key int64) float64 {
	return math.Abs(float64(key ^ magicNumber - time.Now().Unix()))
}

func GenKey(key int64) int64 {
	return key ^ magicNumber
}

func ValidateKey(key, secretKey int64) float64 {
	return math.Abs(float64(secretKey ^ magicNumber - key))
}
