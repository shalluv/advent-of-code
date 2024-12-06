package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	dat, err := os.ReadFile("./src/day1/1.in")
	check(err)

	var arr1, arr2 []int
	s := strings.Split(string(dat), "\n")
	for _, v := range s {
		v_split := strings.Fields(v)
		a, err := strconv.Atoi(v_split[0])
		check(err)
		b, err := strconv.Atoi(v_split[1])
		check(err)
		arr1 = append(arr1, a)
		arr2 = append(arr2, b)
	}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	ans := 0
	for i := 0; i < 1000; i++ {
		if arr1[i] < arr2[i] {
			ans += arr2[i] - arr1[i]
		} else {
			ans += arr1[i] - arr2[i]
		}
	}

	fmt.Println(ans)
}
