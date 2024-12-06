package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	dat, err := os.ReadFile("./src/day2/1.in")
	check(err)

	s := strings.Split(string(dat), "\n")
	ans := 0
	for _, v := range s {
		v_splits := strings.Fields(v)
		var v_ints []int
		for _, txt := range v_splits {
			num, err := strconv.Atoi(txt)
			check(err)
			v_ints = append(v_ints, num)
		}
		succ := true
		inc := v_ints[0] < v_ints[1]
		for i := 0; i < len(v_ints)-1; i++ {
			if !succ {
				continue
			}
			diff := v_ints[i+1] - v_ints[i]
			if diff == 0 || abs(diff) > 3 {
				succ = false
			} else if (inc && diff < 0) || (!inc && diff > 0) {
				succ = false
			}
		}
		if succ {
			ans++
		}
	}

	fmt.Println(ans)
}
