package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1() {
	dat, err := os.ReadFile("./src/day3/1.in")
	check(err)

	re, err := regexp.Compile("mul\\(\\d+,\\d+\\)")
	check(err)
	found := re.FindAllStringSubmatch(string(dat), -1)
	ans := 0
	for i := 0; i < len(found); i++ {
		s := found[i][0]
		s = s[4 : len(s)-1]
		nums := strings.Split(s, ",")
		a, err := strconv.Atoi(nums[0])
		check(err)
		b, err := strconv.Atoi(nums[1])
		check(err)
		ans += a * b
	}

	fmt.Println(ans)
}
