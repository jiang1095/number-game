package helper

import (
	"math"
	"strconv"
)

func Compare(base, guess string) (int, int) {
	if len(base) != 4 || len(guess) != 4 {
		return -1, -1
	}
	var count []int = make([]int, 10)
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

func MaxMin(num_set []string) (string, map[string][]string) {
	var result_num = ""
	var score = math.MaxInt32
	var result_set map[string][]string = make(map[string][]string)
	for _, base := range num_set {
		var set map[string][]string = make(map[string][]string)
		var num_score = 0
		for _, v := range num_set {
			a, b := Compare(base, v)
			key := strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
			set[key] = append(set[key], v)
		}
		for _, v := range set {
			if len(v) > num_score {
				num_score = len(v)
			}
		}
		if num_score < score {
			score = num_score
			result_num = base
			result_set = set
		}
	}
	return result_num, result_set
}
