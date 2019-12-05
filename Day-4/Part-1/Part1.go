package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func checkInc(num int) bool {
	str := strconv.Itoa(num)
	chars := strings.Split(str, "")
	digits := []int{}
	for _, c := range chars {
		d, _ := strconv.Atoi(c)
		digits = append(digits, d)
	}
	flag := true
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			flag = false
		}
	}
	return flag
}

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
	bottom := 145888 // Actual input is 145852, but that includes impossible numbers
	top := 599999    // Actual input is 616942, but that includes impossible numbers
	count := 0
	for i := bottom; i <= top; i++ {
		if checkInc(i) && checkDouble(i) {
			count++
		}
	}
	fmt.Println(count)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
