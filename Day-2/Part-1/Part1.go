package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func print(arr []int) {
	out := ""
	for i := 0; i < len(arr); i++ {
		out += strconv.Itoa(arr[i])
		if i != (len(arr) - 1) {
			out += ","
		}
	}
	fmt.Println(out)
}

func main() {
	start := time.Now()
	dat, _ := ioutil.ReadFile("res/input")
	nums := strings.Split(string(dat), ",")
	ops := []int{}
	for _, o := range nums {
		i, _ := strconv.Atoi(o)
		ops = append(ops, i)
	}
	for i := 0; i < (len(ops) - 4); i += 4 {
		cont := true
		switch ops[i] {
		case 1:
			ops[ops[i+3]] = ops[ops[i+1]] + ops[ops[i+2]]
			break

		case 2:
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
	print(ops)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
