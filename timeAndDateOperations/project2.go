package timeanddateoperations

import (
	"fmt"
	"time"
)

func Demo2() {

	x := fmt.Println
	var now time.Time = time.Now()
	var xnow time.Time = time.Date(2004, time.September, 6, 10, 10, 10, 10, time.UTC)
	x("-----")
	x(now.Year())
	x(now.Month())
	x(now.Day())
	x(now.Hour())
	x(now.Minute())
	x(now.Second())
	x(now.Nanosecond())
	x(now.Location())
	x(xnow.Before(now))
	x(now.After(xnow))
	x(now.Equal(xnow))
	diff := xnow.Sub(now)
	fmt.Println(diff)
}
