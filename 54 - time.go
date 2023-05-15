package main

import (
	"fmt"
	"time"
)
/*
-> Go offers extensive support for times and durations
*/
func main() {
	p := fmt.Println

	//getting current time
	now := time.Now()
	p(now)

	//build a time struct by providing y, m, d, etc
	//times are associated with time zones
	then := time.Date(2009, 11, 17, 20,34,58,651387237, time.UTC)
	p(then)

	//extrac the various components of the time value
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	//the monday-sunday is also available
	p(then.Weekday())

	//it also supports time comparison
	p(then.Before(now))
	p(then.Equal(now))
	p(then.After(now))

	//the sub method  returns a durantion interval between two times
	diff := now.Sub(then)
	p(diff)

	//it is possible to compute the length of duration in a several granularity
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	//advance a time by a given duration
	p(then.Add(diff))
	p(then.Add(-diff))
}