package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(dayOfYear("2019-01-09"))
}

type Month int

const (
	January Month = 1 + iota
	February
	March
	April
	May
	June
	July
	August
	September
	October
	November
	December
)

func dayOfYear(date string) int {
	str := strings.Split(date, "-")
	y, _ := strconv.Atoi(str[0])

	var mStr uint8
	var m int
	if str[1][0] == 0 {
		mStr = str[1][1]
		m = int(mStr)
	}
	m, _ = strconv.Atoi(str[1])
	d, _ := strconv.Atoi(str[2])

	t := time.Date(y, time.Month(m), d, 1,
		00, 00, 0, time.UTC)

	yrday := t.YearDay()

	return yrday
}
