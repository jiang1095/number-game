package helper

import (
	"math"
	"strconv"
)

func Compare(base, guess string) (int, int) {
	if len(base) != 4 || len(guess) != 4 {
		return -1, -1
	}
	var count = make([]int, 10)
	baseBytes := []byte(base)
	guessBytes := []byte(guess)
	a, b := 0, 0
	for i, v := range baseBytes {
		if v == guessBytes[i] {
			a++
		} else {
			count[v-'0']++
			if count[v-'0'] <= 0 {
				b++
			}
			count[guessBytes[i]-'0']--
			if count[guessBytes[i]-'0'] >= 0 {
				b++
			}
		}
	}
	return a, b
}

func MaxMin(numSet []string) (string, map[string][]string) {
	var resultNum = ""
	var score = math.MaxInt32
	var resultSet = make(map[string][]string)
	for _, base := range numSet {
		var set = make(map[string][]string)
		var numScore = 0
		for _, v := range numSet {
			a, b := Compare(base, v)
			key := strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
			set[key] = append(set[key], v)
		}
		for _, v := range set {
			if len(v) > numScore {
				numScore = len(v)
			}
		}
		if numScore < score {
			score = numScore
			resultNum = base
			resultSet = set
		}
	}
	return resultNum, resultSet
}
