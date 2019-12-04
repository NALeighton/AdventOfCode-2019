package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x     int
	y     int
	wires []int
}

func (p point) contains(id int) bool {
	flag := false
	for _, i := range p.wires {
		if i == id {
			flag = true
		}
	}
	return flag
}

func getValues(m map[string]int) []int {
	out := []int{}
	for _, v := range m {
		out = append(out, v)
	}
	return out
}

func getLength(w []string, p point) int {
	x := 0
	y := 0
	length := 0
	for _, s := range w {
		dir := string(s[0])
		dists := s[1:]
		dist, _ := strconv.Atoi(dists)
		for i := 0; i < dist; i++ {
			if x == p.x && y == p.y {
				break
			}
			length++
			switch dir {
			case "U":
				y++
				break
			case "D":
				y--
				break
			case "R":
				x++
				break
			case "L":
				x--
				break
			}
		}
	}
	return length
}

func main() {
	start := time.Now()
	dat, _ := ioutil.ReadFile("res/input")
	input := strings.Split(string(dat), "\n")
	wires := [][]string{}
	for _, w := range input {
		wires = append(wires, strings.Split(w, ","))
	}
	panel := map[string]point{}
	intersections := map[string]int{}
	wireID := 0
	for _, w := range wires {
		x := 0
		y := 0
		for _, p := range w {
			dir := string(p[0])
			dists := p[1:]
			dist, _ := strconv.Atoi(dists)
			for i := 0; i < dist; i++ {
				loc := fmt.Sprintf("%d,%d", x, y)
				if x != 0 && y != 0 {
					v, ok := panel[loc]
					if ok {
						if v.contains(wireID) == false {
							v.wires = append(v.wires, wireID)
							panel[loc] = v
							intersections[loc] = getLength(wires[v.wires[0]], v) + getLength(wires[v.wires[1]], v)
						}
					} else {
						panel[loc] = point{x, y, []int{wireID}}
					}
				}
				switch dir {
				case "U":
					y++
					break
				case "D":
					y--
					break
				case "R":
					x++
					break
				case "L":
					x--
					break
				}
			}
		}
		wireID++
	}
	dists := getValues(intersections)
	sort.Ints(dists)
	fmt.Println(dists[0])
	end := time.Now()
	fmt.Println(end.Sub(start))
}
