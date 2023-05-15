package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()
	//formatting a time according to RFC3339
	p(t.Format(time.RFC3339))

	//parsing time
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	if e != nil {
		panic(e)
	}

	p(t1)
	//expressing other formats of time
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-01:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	
	ansic := "Mon Jan _2 15:04:05 2006"
	if _, e = time.Parse(ansic, "8:41PM"); e != nil {
		panic(e)
	}
}