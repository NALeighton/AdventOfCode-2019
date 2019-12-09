package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var input int = 1
var ops []int

func add(a int, b int) int {
	return a + b
}

func multiply(a int, b int) int {
	return a * b
}

func getVal(mode int, input int) int {
	if mode == 1 {
		return input
	}
	return ops[input]
}

func getInput(pos int) {
	ops[pos] = input
}

func handleOutput(val int) {
	fmt.Println(val)
}

func main() {
	start := time.Now()
	dat, _ := ioutil.ReadFile("res/input")
	nums := strings.Split(string(dat), ",")
	for _, o := range nums {
		i, _ := strconv.Atoi(o)
		ops = append(ops, i)
	}
	cont := true
	for i := 0; i < (len(ops)); {
		switch ops[i] {
		// ADD
		case 1:
			ops[ops[i+3]] = add(getVal(0, ops[i+1]), getVal(0, ops[i+2]))
			i += 4
			break
		case 101:
			ops[ops[i+3]] = add(getVal(1, ops[i+1]), getVal(0, ops[i+2]))
			i += 4
			break
		case 1001:
			ops[ops[i+3]] = add(getVal(0, ops[i+1]), getVal(1, ops[i+2]))
			i += 4
			break
		case 1101:
			ops[ops[i+3]] = add(getVal(1, ops[i+1]), getVal(1, ops[i+2]))
			i += 4
			break

		// MULTIPLY
		case 2:
			ops[ops[i+3]] = multiply(getVal(0, ops[i+1]), getVal(0, ops[i+2]))
			i += 4
			break
		case 102:
			ops[ops[i+3]] = multiply(getVal(1, ops[i+1]), getVal(0, ops[i+2]))
			i += 4
			break
		case 1002:
			ops[ops[i+3]] = multiply(getVal(0, ops[i+1]), getVal(1, ops[i+2]))
			i += 4
			break
		case 1102:
			ops[ops[i+3]] = multiply(getVal(1, ops[i+1]), getVal(1, ops[i+2]))
			i += 4
			break

		// INPUT
		case 3:
			getInput(ops[i+1])
			i += 2
			break

		// OUTPUT
		case 4:
			handleOutput(getVal(0, ops[i+1]))
			i += 2
			break
		case 104:
			handleOutput(getVal(1, ops[i+1]))
			i += 2
			break
		case 99:
			cont = false
			break
		default:
			cont = false
			fmt.Printf("BAD OP-CODE: %d\n", ops[i])
		}
		if cont == false {
			break
		}
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}
