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
