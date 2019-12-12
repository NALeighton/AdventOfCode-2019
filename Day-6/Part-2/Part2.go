package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/yourbasic/graph"
)

var bodies map[string]int

func main() {
	start := time.Now()
	dat, _ := ioutil.ReadFile("res/input")
	orbits := []string{}
	bodies = make(map[string]int)
	for i := 0; i < len(dat); {
		tmp := ""
		for p := i; p < len(dat); p++ {
			if string(dat[p]) == "\n" {
				i = p + 1
				orbits = append(orbits, tmp)
				obs := strings.Split(orbits[len(orbits)-1], ")")
				if _, ok := bodies[obs[0]]; !ok {
					bodies[obs[0]] = len(bodies)
				}
				if _, ok := bodies[obs[1]]; !ok {
					bodies[obs[1]] = len(bodies)
				}
				break
			} else if p == len(dat)-1 {
				tmp += string(dat[p])
				i = p + 1
				orbits = append(orbits, tmp)
				obs := strings.Split(orbits[len(orbits)-1], ")")
				if _, ok := bodies[obs[0]]; !ok {
					bodies[obs[0]] = len(bodies)
				}
				if _, ok := bodies[obs[1]]; !ok {
					bodies[obs[1]] = len(bodies)
				}
				break
			} else {
				tmp += string(dat[p])
			}
		}
	}
	oMap := graph.New(len(bodies))
	for _, o := range orbits {
		tmp := strings.Split(o, ")")
		oMap.AddBothCost(bodies[tmp[0]], bodies[tmp[1]], 1)
	}
	_, dist := graph.ShortestPath(oMap, bodies["YOU"], bodies["SAN"])
	fmt.Println(dist - 2)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
