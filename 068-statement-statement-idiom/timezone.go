package main

import "log"

func main(tz string) int {

	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}

	var seconds int
	var ok bool
	seconds, ok = timeZone[tz]
	tz := "EST"

	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}
