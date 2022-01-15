package verification

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/rand"
)

const layout = "20060102"

/* 关于 time 包的用法 */

func ParseSec() {
	t := time.Unix(1626333388, 0).Format("200601")
	fmt.Println(t)
}

func RandFromBound(min, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func Timestamp2DayKey(timestamp int64) int {
	ds := time.Unix(timestamp, 0).UTC().Format(layout)
	date, _ := strconv.Atoi(ds)
	return date
}

const (
	pushTaskDateFormat = "%4d-%02d-%02d"
	pushTaskTimeFormat = "2006-01-02 15:04:05"
)

func parseExpectedTime(t string) time.Time {
	now := time.Now()
	timeStr := strings.Join([]string{fmt.Sprintf(pushTaskDateFormat, now.Year(), now.Month(), now.Day()), t}, " ")
	parsedTime, _ := time.ParseInLocation(pushTaskTimeFormat, timeStr, time.UTC)
	return parsedTime
}

func generateQuick() {
	t := time.Date(2021, 11, 03, 10, 50, 31, 1, time.Local)
	source := rand.NewSource(uint64(t.UnixNano()))
	entropy := rand.New(source)

	fmt.Println(ulid.MustNew(ulid.Timestamp(t), entropy))
	fmt.Println(ulid.Now())
}

func t(a int) {
	println(a)
}