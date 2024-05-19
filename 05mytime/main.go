package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Welcome to Time Study of GoLang")
	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)
	
	fmt.Println(currentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.October, 13, 23, 23, 0, 0, time.UTC)

	fmt.Println("Created Date:", createdDate)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
}