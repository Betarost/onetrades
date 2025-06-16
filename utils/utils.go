package utils

import (
	"strconv"
	"strings"
	"time"
)

func CurrentTimestamp() int64 {
	return int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
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

func TimestampMilliToDateFormat(stamp int64, format string) string {
	tHisory := time.UnixMilli(stamp).UTC()
	hisNow := time.Date(tHisory.Year(), tHisory.Month(), tHisory.Day(), tHisory.Hour(), tHisory.Minute(), tHisory.Second(), 0, tHisory.Location())
	return hisNow.Format(format)
}

func FloatToStringAll(num float64) string {
	return strconv.FormatFloat(num, 'f', -1, 64)
}
