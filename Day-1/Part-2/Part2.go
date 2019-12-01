package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	dat, _ := ioutil.ReadFile("res/input")
	nums := strings.Split(string(dat), "\n")
	total := 0
	for i := 0; i < len(nums); i++ {
		num, _ := strconv.Atoi(nums[i])
		subTotal := 0
		ans := 1
		for ans > 0 {
			ans = int(math.Floor(float64(num)/3)) - 2
			num = ans
			if ans > 0 {
				subTotal += ans
			}
		}
		total += subTotal
	}
	fmt.Println(total)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
