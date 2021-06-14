package time

import (
	"time"
)

func GetTimeFromString(stringTime string) time.Time {
	mydate, _ := time.Parse("2006-01-02", stringTime)

	return mydate
}
