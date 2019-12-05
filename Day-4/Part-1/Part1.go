package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func checkDouble(num int) bool {
	str := strconv.Itoa(num)
	chars := strings.Split(str, "")
	digits := []int{}
	for _, c := range chars {
		d, _ := strconv.Atoi(c)
		digits = append(digits, d)
	}
	flag := false
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] == digits[i+1] {
			flag = true
		}
	}
	return flag
}

func main() {
	start := time.Now()
	count := 0
	for f := 1; f <= 5; f++ {
		for e := f; e <= 9; e++ {
			for d := e; d <= 9; d++ {
				for c := d; c <= 9; c++ {
					for b := c; b <= 9; b++ {
						for a := b; a <= 9; a++ {
							num := a + (b * 10) + (c * 100) + (d * 1000) + (e * 10000) + (f * 100000)
							if num >= 145888 {
								if checkDouble(num) {
									count++
								}
							}
						}
					}
				}
			}
		}
	}
	fmt.Println(count)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
