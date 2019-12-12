package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var input int
var first bool = true
var output int
var ops []int
var master []int
var highest int = -999999999
var hI int = 01234

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
	if first {
		ops[pos] = input
		first = false
	} else {
		ops[pos] = output
		first = true
	}
}

func handleOutput(val int) {
	output = val
}

func less(a int, b int, c int) {
	if a < b {
		ops[c] = 1
	} else {
		ops[c] = 0
	}
}

func eq(a int, b int, c int) {
	if a == b {
		ops[c] = 1
	} else {
		ops[c] = 0
	}
}

func run() {
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

		// Jump if true
		case 5:
			if getVal(0, ops[i+1]) != 0 {
				i = getVal(0, ops[i+2])
			} else {
				i += 3
			}
			break
		case 105:
			if getVal(1, ops[i+1]) != 0 {
				i = getVal(0, ops[i+2])
			} else {
				i += 3
			}
			break
		case 1005:
			if getVal(0, ops[i+1]) != 0 {
				i = getVal(1, ops[i+2])
			} else {
				i += 3
			}
			break
		case 1105:
			if getVal(1, ops[i+1]) != 0 {
				i = getVal(1, ops[i+2])
			} else {
				i += 3
			}
			break

		// Jump if false
		case 6:
			if getVal(0, ops[i+1]) == 0 {
				i = getVal(0, ops[i+2])
			} else {
				i += 3
			}
			break
		case 106:
			if getVal(1, ops[i+1]) == 0 {
				i = getVal(0, ops[i+2])
			} else {
				i += 3
			}
			break
		case 1006:
			if getVal(0, ops[i+1]) == 0 {
				i = getVal(1, ops[i+2])
			} else {
				i += 3
			}
			break
		case 1106:
			if getVal(1, ops[i+1]) == 0 {
				i = getVal(1, ops[i+2])
			} else {
				i += 3
			}
			break

		// LESS THAN
		case 7:
			less(getVal(0, ops[i+1]), getVal(0, ops[i+2]), ops[i+3])
			i += 4
			break
		case 107:
			less(getVal(1, ops[i+1]), getVal(0, ops[i+2]), ops[i+3])
			i += 4
			break
		case 1007:
			less(getVal(0, ops[i+1]), getVal(1, ops[i+2]), ops[i+3])
			i += 4
			break
		case 1107:
			less(getVal(1, ops[i+1]), getVal(1, ops[i+2]), ops[i+3])
			i += 4
			break

		// EQUAL
		case 8:
			eq(getVal(0, ops[i+1]), getVal(0, ops[i+2]), ops[i+3])
			i += 4
			break
		case 108:
			eq(getVal(1, ops[i+1]), getVal(0, ops[i+2]), ops[i+3])
			i += 4
			break
		case 1008:
			eq(getVal(0, ops[i+1]), getVal(1, ops[i+2]), ops[i+3])
			i += 4
			break
		case 1108:
			eq(getVal(1, ops[i+1]), getVal(1, ops[i+2]), ops[i+3])
			i += 4
			break

		// HALT
		case 99:
			cont = false
			break

		// ERROR
		default:
			cont = false
			fmt.Printf("BAD OP-CODE: %d\n", ops[i])
			break
		}
		if cont == false {
			break
		}
	}
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

func main() {
	start := time.Now()
	dat, _ := ioutil.ReadFile("res/input")
	nums := strings.Split(string(dat), ",")
	for _, o := range nums {
		i, _ := strconv.Atoi(o)
		master = append(master, i)
	}
	orig := []int{0, 1, 2, 3, 4}
	for p := make([]int, len(orig)); p[0] < len(p); nextPerm(p) {
		ops = master
		perm := getPerm(orig, p)
		output = 0
		for i := 0; i < 5; i++ {
			input = perm[i]
			run()
		}
		if output > highest {
			highest = output
			hI = input
		}
	}
	fmt.Println(highest)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
