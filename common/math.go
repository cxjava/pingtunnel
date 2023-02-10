package common

import (
	"crypto/rand"
	"encoding/base64"
	"hash/fnv"
	"io"
	"math"
	mrand "math/rand"
	"time"
)

func init() {
	mrand.Seed(time.Now().UnixNano())
}

func MinOfInt(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOfInt(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func MinOfInt64(vars ...int64) int64 {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOfInt64(vars ...int64) int64 {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func AbsInt(v int) int {
	if v > 0 {
		return v
	}
	return -v
}

func AbsInt32(v int32) int32 {
	if v > 0 {
		return v
	}
	return -v
}

func AbsInt64(v int64) int64 {
	if v > 0 {
		return v
	}
	return -v
}

func HashString(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

func RandInt31n(n int) int32 {
	ret := mrand.Int31n((int32)(n))
	return int32(ret)
}

func RandInt() int32 {
	ret := mrand.Int()
	return int32(ret)
}

func Shuffle(n int, swap func(i, j int)) {
	mrand.Shuffle(n, swap)
}

func MAKEINT64(high int32, low int32) int64 {
	return (int64)(((int64)(low)) | ((int64)((int32)(high)))<<32)
}
func HIINT32(I int64) int32 {
	return (int32)(((int64)(I) >> 32) & 0xFFFFFFFF)
}
func LOINT32(l int64) int32 {
	return (int32)(l)
}

func MAKEINT32(high int16, low int16) int32 {
	return (int32)(((int32)(low)) | ((int32)((int16)(high)))<<16)
}
func HIINT16(I int32) int16 {
	return (int16)(((int32)(I) >> 16) & 0xFFFF)
}
func LOINT16(l int32) int16 {
	return (int16)(l)
}

func IsInt(r float64) bool {
	return (r - math.Floor(r)) == 0
}

func ArrayContainInt(a []int, f int) bool {

	for _, i := range a {
		if f == i {
			return true
		}
	}

	return false
}

func ArrayContainString(a []string, f string) bool {

	for _, i := range a {
		if f == i {
			return true
		}
	}

	return false
}

func SafeDivide(a int64, b int64) int64 {
	if b == 0 {
		return 0
	}
	return a / b
}
