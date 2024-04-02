package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	startDate := time.Date(1996, 03, 13, 05, 00, 00, 00, time.UTC)
	t := time.Now()
	fmt.Println(t.Format(time.Kitchen))
	fmt.Println(startDate.Format(time.DateTime))
	elapse := time.Since(startDate)
	fmt.Printf("Hours %.0f Minutes: %.0f Seconds: %.0f\n", elapse.Hours(), elapse.Minutes(), elapse.Seconds())

	t = t.AddDate(1, 1, 2)
	fmt.Println(t.Format(time.DateTime))

	bedtime := time.Date(2024, 4, 1, 23, 0, 0, 0, time.Local)
	fmt.Printf("There is %.0f hours until bed time \n", time.Until(bedtime).Hours())

	elapsed := time.Since(start)
	defer fmt.Printf("Elapsed time: %.2f microseconds\n", float64(elapsed.Microseconds()))
}
