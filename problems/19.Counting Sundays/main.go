package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	solve()
}

func solve() {
	layout := "2 Jan 2006"
	fromStr := "1 Jan 1901"
	toStr := "31 Dec 2000"

	from, err := time.Parse(layout, fromStr)
	if err != nil {
		log.Fatal(err)
	}
	to, err := time.Parse(layout, toStr)
	if err != nil {
		log.Fatal(err)
	}

	cnt := 0
	for to.After(from) {
		if from.Weekday() == 0 {
			cnt++
		}
		from = from.AddDate(0, 1, 0)
	}
	fmt.Println(cnt)
}
