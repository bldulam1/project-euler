package main

import (
	"fmt"
	"projecteuler/problem19"
	"time"
)

func main() {
	sundays := 0
	for year := 1901; year <= 2000; year++ {
		for month := uint8(1); month <= 12; month++ {
			weekday := problem19.GetDay(year, month, 1)
			if weekday == time.Sunday {
				sundays++
			}
		}
	}
	fmt.Println(sundays)
}
