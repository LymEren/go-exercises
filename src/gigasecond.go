package main

import (
	"fmt"
	"time"
)

var gigaSecond time.Duration = 1e18

func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigaSecond)
}

func main() {
	fmt.Println(AddGigasecond(time.Now()))
	fmt.Println(time.Now())
	return
}
