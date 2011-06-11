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
	return strconv.Itoa64((now.Seconds() - ny.Seconds()) / 86400) + ", " + strconv.Itoa64(now.Year)
}

func gtime() string {
	now := time.LocalTime()
	hour := int((float64(now.Hour) / 24) * 10)
	minute := int((float64(now.Minute) / 60) * 100)
	second := int((float64(now.Second) / 60) * 1000)
	return strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(second)
}

func main() {
	fmt.Println("Date:", gdate())
	fmt.Println("Time:", gtime())
}

