package common

import (
	"time"
)

var gnowsecond time.Time

func init() {
	gnowsecond = time.Now()
	go updateNowInSecond()
}

func GetNowUpdateInSecond() time.Time {
	return gnowsecond
}

func updateNowInSecond() {
	defer CrashLog()

	for {
		gnowsecond = time.Now()
		time.Sleep(time.Second)
	}
}

func Elapsed(f func(d time.Duration)) func() {
	start := time.Now()
	return func() {
		f(time.Since(start))
	}
}
