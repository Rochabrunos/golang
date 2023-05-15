package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("now:", now)

	//get elapsed time since Unix epoch
	fmt.Println("sec:",now.Unix())
	fmt.Println("mili:", now.UnixMilli())
	fmt.Println("nano:", now.UnixNano())

	//we can also convert integer seconds or nanoseconds since epoch into the corresponding time
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.Unix()))

}