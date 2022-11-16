package main

import (
	"fmt"
	"strings"
)

var numberMap = map[int]string{
	1:    "one",
	2:    "two",
	3:    "three",
	4:    "four",
	5:    "five",
	6:    "six",
	7:    "seven",
	8:    "eight",
	9:    "nine",
	10:   "ten",
	11:   "eleven",
	12:   "twelve",
	13:   "thirteen",
	14:   "fourteen",
	15:   "fifteen",
	16:   "sixteen",
	17:   "seventeen",
	18:   "eighteen",
	19:   "nineteen",
	20:   "twenty",
	30:   "thirty",
	40:   "forty",
	50:   "fifty",
	60:   "sixty",
	70:   "seventy",
	80:   "eighty",
	90:   "ninety",
	100:  "hundred",
	1000: "onethousand",
}

func main() {
	solve()
}

func solve() {
	str := make([]string, 0, 1000)
	for i := 1; i <= 1000; i++ {
		str = append(str, parseNum(i))
	}
	fmt.Println(len(strings.Join(str, "")))
}

func parseNum(i int) string {
	var str []string
	if i <= 20 || i == 1000 {
		str = append(str, numberMap[i])
	} else if i < 100 {
		s1 := numberMap[i%10]
		s2 := numberMap[i-i%10]
		str = append(str, s2, s1)
	} else if i < 1000 && i%100 <= 20 {
		s1 := numberMap[i%100]
		s3 := numberMap[i/100] + "hundred"
		if s1 != "" {
			s3 += "and"
		}
		str = append(str, s3, s1)
	} else {
		s1 := numberMap[i%10]
		s2 := numberMap[i%100-i%10]
		s3 := numberMap[i/100] + "hundred"
		if s1 != "" || s2 != "" {
			s3 += "and"
		}
		str = append(str, s3, s2, s1)
	}
	return strings.Join(str, "")
}
