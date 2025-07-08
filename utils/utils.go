package utils

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CurrentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
}

func CurrentSecondsTimestamp() int64 {
	return int64(time.Millisecond) * time.Now().UnixMilli() / int64(time.Second)
}

func StringToFloat(num string) float64 {
	newNum, _ := strconv.ParseFloat(strings.Replace(num, ",", ".", -1), 64)
	return newNum
}

func StringToInt64(num string) int64 {
	newNum, _ := strconv.ParseInt(num, 10, 64)
	return newNum
}

func StringToInt(num string) int {
	newNum, _ := strconv.ParseInt(num, 10, 0)
	return int(newNum)
}

func Int64ToString(num int64) string {
	return fmt.Sprintf("%d", num)
}

func IntToString(num int) string {
	return fmt.Sprintf("%d", num)
}

func TimestampMilliToDateFormat(stamp int64, format string) string {
	tHisory := time.UnixMilli(stamp).UTC()
	hisNow := time.Date(tHisory.Year(), tHisory.Month(), tHisory.Day(), tHisory.Hour(), tHisory.Minute(), tHisory.Second(), 0, tHisory.Location())
	return hisNow.Format(format)
}

func FloatToStringAll(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}

func GetPrecisionFromStr(s string) (pr string) {
	pr = "0"
	sp := strings.Split(s, ".")
	if len(sp) > 1 {
		pr = fmt.Sprintf("%d", len(sp[1]))
	}
	return pr
}
