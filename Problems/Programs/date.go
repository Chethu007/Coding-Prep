package main

import (
	"fmt"
	"time"
)

func main() {
	format := "2006-01-02T15:04:05.000"
	postedDate := time.Date(2023, 2, 4, 0, 0, 0, 0, time.UTC)

	fmt.Printf("Posted Date: %v", postedDate.AddDate(0, 0, -1).Format(format))
}
