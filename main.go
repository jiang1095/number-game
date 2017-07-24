package main

import (
	"fmt"
	"strconv"
)

var number_set []string

func init() {
	for i := 100; i < 10000; i++ {
		a := i / 1000
		b := (i / 100) % 10
		c := (i / 10) % 10
		d := i % 10
		if a == b || a == c || a == d || b == c || b == d || c == d {
			continue
		} else {
			if i < 1000 {
				number_set = append(number_set, "0"+strconv.Itoa(i))
			} else {
				number_set = append(number_set, strconv.Itoa(i))
			}
		}
	}
}

func compare(base, guess string) (int, int) {
	if len(base) != 4 || len(guess) != 4 {
		return -1, -1
	}
	a, b := 0, 0
	for i1, v1 := range []byte(base) {
		for i2, v2 := range []byte(guess) {
			if v1 == v2 {
				if i1 == i2 {
					a++
				} else {
					b++
				}
			}
		}
	}
	return a, b
}

func main() {
	var set map[string][]string = make(map[string][]string)
	for _, v := range number_set {
		a, b := compare("0123", v)
		key := strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
		set[key] = append(set[key], v)
	}
	for k, v := range set {
		fmt.Println(k, v)
	}
}
