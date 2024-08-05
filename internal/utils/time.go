package utils

import "time"

func CalculateTimeDifference(currentTime time.Time, futureTime time.Time) time.Duration {
	diff := futureTime.Sub(currentTime)
	return diff
}
