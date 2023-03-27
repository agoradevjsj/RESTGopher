package timeexec

import "time"

var start_time time.Time
var end_time time.Time

func StartTimer() time.Time {
	start_time = time.Now()
	return start_time
}

func StopTime() time.Time {
	end_time = time.Now()
	return end_time
}

func GetTime() time.Duration {
	return end_time.Sub(start_time)
}

func GetTimeParams(startTime time.Time, endTime time.Time) time.Duration {
	return startTime.Sub(endTime)
}
