package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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

func main() {
	var game_type int
	fmt.Println("这是一个猜数字小游戏，你可以选择猜电脑生成的数字，或者让电脑猜你给出的数字：")
	fmt.Println("\t1. 你来猜")
	fmt.Println("\t2. 让我猜")
	reader := bufio.NewReader(os.Stdin)
	for {
		for {
			fmt.Print("请选择游戏模式(1/2): ")
			data, _, _ := reader.ReadLine()
			game_type, _ = strconv.Atoi(string(data))
			if game_type == 1 {
				fmt.Println("你来猜我的数字")
				numberGame()
				break
			} else if game_type == 2 {
				fmt.Println("我来猜你的数字")
				guessNumber()
				break
			}
		}
		fmt.Print("再来一次？(y/n): ")
		data, _, _ := reader.ReadLine()
		switch string(data) {
		case "y":
		case "n":
			os.Exit(0)
		default:
			os.Exit(1)
		}
	}
}

func numberGame() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	number := number_set[rand.Intn(5040)]
	for i := 0; i < 10; i++ {
		fmt.Printf("请输入你的猜测(还剩%d次机会): ", 10-i)
		data, _, _ := reader.ReadLine()
		if match, _ := regexp.Match("^[0-9]*$", data); len(data) != 4 || !match {
			fmt.Println("输入非法，请重新输入。")
			i--
			continue
		}
		guess := string(data)
		a, b := compare(number, guess)
		if a == 4 {
			fmt.Printf("恭喜你在第%d次猜出了这个数字！\n", i+1)
			break
		} else {
			fmt.Printf("本次猜测结果为: %dA%dB\n", a, b)
			if i == 9 {
				fmt.Printf("很抱歉，你没能在十次以内猜到数字。正确答案是: %s\n", number)
			}
		}
	}
}

func guessNumber() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	guess_number_set := number_set
	times := 0
	for {
		var (
			set         map[string][]string = make(map[string][]string)
			guess_cache []string
			state_cache []string
		)
		guess := guess_number_set[rand.Intn(len(guess_number_set))]
		guess_cache = append(guess_cache, guess)
		for _, v := range guess_number_set {
			a, b := compare(guess, v)
			key := strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
			set[key] = append(set[key], v)
		}
		times += 1
		fmt.Println("我猜你的数字是:", guess)
		fmt.Printf("第%d次猜测结果为: ", times)
		data, _, _ := reader.ReadLine()
		state := strings.ToUpper(string(data))
		state_cache = append(state_cache, state)
		if state == "4A0B" {
			fmt.Printf("尽管你的数字很难猜，我还是在第%d次把它猜出来了！\n", times)
			return
		} else {
			for {
				if set[state] == nil {
					fmt.Println("你输入的状态造成了我的困惑，确认没有输错吗？")
					fmt.Print("重新输入本次猜测结果？(y/n)")
					data, _, _ := reader.ReadLine()
					switch string(data) {
					case "y":
						fmt.Printf("第%d次猜测结果为: ", times)
						data, _, _ := reader.ReadLine()
						state = strings.ToUpper(string(data))
					case "n":
						fmt.Println("好吧，你赢了，我没办法猜出你的数字……")
						fmt.Print("告诉我正确答案，让我看看呗：")
						data, _, _ := reader.ReadLine()
						answer := string(data)
						checkAnswer(guess_cache, state_cache, answer)
						return
					}
				} else {
					guess_number_set = set[state]
					break
				}
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

func checkAnswer(guesses, states []string, answer string) {
	if len(guesses) != len(states) {
		fmt.Println("猜测次数和给定的状态数目不匹配，无法检测！")
		return
	}
	for i, v := range guesses {
		a, b := compare(answer, v)
		state := fmt.Sprintf("%dA%dB", a, b)
		if state != states[i] {
			fmt.Printf("在第%d次猜测中\n你给出的结果是:%s\n但我认为应该是:%s\n这可能是我没猜出来的原因\n", i+1, states[i], state)
			return
		}
	}
	fmt.Println("好吧，我没法找出自己失败的原因，你实在是太厉害了！！！")
}
