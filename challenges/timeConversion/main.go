package timeConversion

import (
	"strconv"
	"strings"
	"time"
)

//https://www.hackerrank.com/challenges/time-conversion/problem

func Validate(data string) bool {
	//12:00:00AM OR 12:00:00PM
	return len(data) == 10 &&
		(strings.HasSuffix(data, "AM") || strings.HasSuffix(data, "PM")) &&
		strings.Contains(data, ":")
}

//working with strings, not an optimal method
func Run(timeAMPM string) string {
	if !Validate(timeAMPM) {
		return ""
	}
	strLen := len(timeAMPM)

	//get AM or PM
	amOrPm := timeAMPM[strLen-2:]

	//[11 00 00AM]
	timeAMPMArr := strings.Split(timeAMPM, ":")

	hours, err := strconv.Atoi(timeAMPMArr[0])

	if err != nil {
		return ""
	}

	//remove AM/PM from string
	timeAMPMArr[2] = strings.Replace(timeAMPMArr[2], amOrPm, "", -1)

	//11:45:00PM
	if hours < 12 {
		if amOrPm == "PM" {
			timeAMPMArr[0] = strconv.Itoa(12 + hours)
		}

	} else {
		//12:45:00AM
		if amOrPm == "AM" {
			timeAMPMArr[0] = "00"
		}

	}
	return strings.Join(timeAMPMArr, ":")
}

//using standard time functions
func Run2(timeAMPM string) string {
	formatAMPM := "03:04:05PM"
	format24 := "15:04:05"

	t, err := time.Parse(formatAMPM, timeAMPM)

	if err != nil {
		return ""
	}

	return t.Format(format24)
}
