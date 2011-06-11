package main

import (
	"strconv"
	"time"
	"web"
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
		tsec = (now.Seconds() - t.Seconds()) * 1000
	}
	ms := int64(86400 * 1000)
	hour := int64((float64(tsec) / float64(ms)) * 10)
	tsec -= int64((ms / 10) * hour)
	minute := int64((float64(tsec) / (float64(ms) / 10)) * 100)
	tsec -= int64((ms / 1000) * minute)
	second := int64((float64(tsec) / (float64(ms) / 100)) * 10000)
	return strconv.Itoa64(hour) + ":" + strconv.Itoa64(minute) + ":" + strconv.Itoa64(second)
}

func get(ctx *web.Context, val string) string {
	retval := "Date: " + gdate()
	retval += "<br>Time: " + gtime()
	return retval
}

func main() {
	web.Get("/(.*)", get)
	web.Run("0.0.0.0:7001")
}
