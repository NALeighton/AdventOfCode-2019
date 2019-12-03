package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	dat, _ := ioutil.ReadFile("res/input")
	nums := strings.Split(string(dat), ",")
	base := []int{}
	for _, o := range nums {
		i, _ := strconv.Atoi(o)
		base = append(base, i)
	}
	ans := 0
	a := -1
	b := 0
	for ans != 19690720 {
		a++
		if a > 99 {
			a = 0
			b++
		}
		if b > 99 {
			break
		}
		ops := [149]int{} // I am fully aware of how bad this is, I could not find a better solution that wasn't simultaneously significantly slower
		copy(ops[:], base)
		ops[1] = a
		ops[2] = b
		for i := 0; i < (len(ops) - 4); i += 4 {
			cont := true
			switch ops[i] {
			case 1:
				if ops[i+1] >= len(ops) || ops[i+2] >= len(ops) || ops[i+3] >= len(ops) {
					break
				}
				ops[ops[i+3]] = ops[ops[i+1]] + ops[ops[i+2]]
				break

			case 2:
				if ops[i+1] > len(ops) || ops[i+2] > len(ops) || ops[i+3] > len(ops) {
					break
				}
				ops[ops[i+3]] = ops[ops[i+1]] * ops[ops[i+2]]
				break
			case 99:
				cont = false
				break
			}
			if cont == false {
				break
			}
		}
		ans = ops[0]
	}
	fmt.Println(100*a + b)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
