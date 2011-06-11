package main

import (
	"fmt"
	"strconv"
	"time"
)

func gdate() string {
	now := time.LocalTime()
	var ny time.Time
	ny.Year = now.Year
	ny.Month = 1
	ny.Day = 1
	ny.Zone = "CST"
	return strconv.Itoa64((now.Seconds()-ny.Seconds())/86400) + ", " + strconv.Itoa64(now.Year)
}

func gtime() string {
	now := time.LocalTime()
	var tsec int64
	{
		var t time.Time = *now
		t.Hour = 0
		t.Minute = 0
		t.Second = 0
		tsec = now.Seconds() - t.Seconds()
	}
	hour := int((float64(tsec) / 86400) * 10)
	tsec -= int64(8640 * hour)
	minute := int((float64(tsec) / 8640) * 100)
	tsec -= int64((864 / 100) * minute)
	second := int((float64(tsec) / 864) * 1000)
	return strconv.Itoa(int(hour)) + ":" + strconv.Itoa(int(minute)) + ":" + strconv.Itoa(int(second))
}

func main() {
	fmt.Println("Date:", gdate())
	fmt.Println("Time:", gtime())
}
