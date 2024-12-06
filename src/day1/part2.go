package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {
	dat, err := os.ReadFile("./src/day1/2.in")
	check(err)

	var map1, map2 = make(map[int]int), make(map[int]int)
	s := strings.Split(string(dat), "\n")
	for _, v := range s {
		v_split := strings.Fields(v)
		a, err := strconv.Atoi(v_split[0])
		check(err)
		b, err := strconv.Atoi(v_split[1])
		check(err)
		map1[a]++
		map2[b]++
	}

	ans := 0
	for k, v := range map1 {
		if val, ok := map2[k]; ok {
			ans += k * v * val
		}
	}
	fmt.Println(ans)
}
