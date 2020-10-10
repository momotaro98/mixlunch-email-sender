package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseBeginEndTime(beginDateTimeStr, endDateTimeStr string) (time.Time, time.Time, error) {
	// Parse begin and end DateTime string to RFC3339 spec
	beginDateTime, err := time.Parse(time.RFC3339, beginDateTimeStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	endDateTime, err := time.Parse(time.RFC3339, endDateTimeStr)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	return beginDateTime, endDateTime, nil
}

func ParseDatetimeString(dt string) (t time.Time, err error) {
	t, err = time.Parse(time.RFC3339, dt)
	if err == nil {
		return t, nil
	}

	t, err = time.Parse("2006-01-02", dt)
	if err == nil {
		return t, nil
	}

	return t, fmt.Errorf("date string pare error. dt: %s", dt)
}

func validateDateStringFormat(dateString string) bool {
	slice := strings.Split(dateString, "-")
	if len(slice) != 3 {
		return false
	}
	year, err := strconv.Atoi(slice[0])
	if err != nil {
		return false
	}
	if year < 2000 {
		return false
	}
	month, err := strconv.Atoi(slice[1])
	if err != nil {
		return false
	}
	if month < 1 || month > 12 {
		return false
	}
	day, err := strconv.Atoi(slice[2])
	if err != nil {
		return false
	}
	if day < 1 || day > 31 {
		return false
	}
	return true
}

func SquashDateString(dateString string) (string, error) {
	if !validateDateStringFormat(dateString) {
		return "", fmt.Errorf("not date string: %s", dateString)
	}
	return strings.Join(strings.Split(dateString, "-"), ""), nil
}

func ShiftTimeStringUtfToJst(timeString string) (string, error) {
	s := strings.Split(timeString, ":")
	if len(s) < 2 {
		return "", fmt.Errorf("not tiem string: %s", timeString)
	}
	utfHour, err := strconv.Atoi(s[0])
	if err != nil {
		return "", err
	}
	jstHourString := fmt.Sprintf("%02d", utfHour+9)
	return jstHourString + ":" + s[1], nil
}
