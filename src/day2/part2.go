package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_valid(arr []int, inc bool) bool {
	for i := 0; i < len(arr)-1; i++ {
		diff := arr[i+1] - arr[i]
		if abs(diff) > 3 || diff == 0 {
			return false
		}
		if inc && diff < 0 {
			return false
		}
		if !inc && diff > 0 {
			return false
		}
	}
	return true
}

func part2() {
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

		if is_valid(v_ints, true) || is_valid(v_ints, false) {
			ans++
		} else {
			for removing := 0; removing < len(v_ints); removing++ {
				var new_arr []int
				for i, v := range v_ints {
					if i == removing {
						continue
					}
					new_arr = append(new_arr, v)
				}
				if is_valid(new_arr, true) || is_valid(new_arr, false) {
					ans++
					break
				}
			}
		}
	}

	fmt.Println(ans)
}
