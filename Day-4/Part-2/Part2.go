package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func checkDouble(num int) bool {
	str := strconv.Itoa(num)
	doublere := regexp.MustCompile(`(11|22|33|44|55|66|77|88|99)`)
	triplere := regexp.MustCompile(`(111|222|333|444|555|666|777|888|999)`)
	double := doublere.FindAllString(str, -1)
	triple := triplere.FindAllString(str, -1)
	flag := false
	if len(triple) == 2 {
		flag = false
	} else if len(double) > 1 && len(triple) == 1 {
		flag = (double[0][0] != triple[0][0] || double[1][0] != triple[0][0])
	} else if len(double) == 1 && len(triple) == 1 {
		flag = (double[0][0] != triple[0][0])
	} else if len(double) >= 1 {
		flag = true
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
							if num > 145852 {
								fmt.Println(num, checkDouble(num))
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
