package main

import (
	"bufio"
	"fmt"
	"github.com/jiang1095/number-game/helper"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var noDuplicateNumbersSet []string
var duplicateNumbersSet []string

func init() {
	for i := 0; i < 10000; i++ {
		a := i / 1000
		b := (i / 100) % 10
		c := (i / 10) % 10
		d := i % 10
		duplicateNumbersSet = append(duplicateNumbersSet, strconv.Itoa(a)+strconv.Itoa(b)+strconv.Itoa(c)+strconv.Itoa(d))
		if a == b || a == c || a == d || b == c || b == d || c == d {
			continue
		} else {
			noDuplicateNumbersSet = append(noDuplicateNumbersSet, strconv.Itoa(a)+strconv.Itoa(b)+strconv.Itoa(c)+strconv.Itoa(d))
		}
	}
}

func main() {
	var gameType int
	fmt.Println("这是一个猜数字小游戏，你可以选择猜电脑生成的数字，或者让电脑猜你给出的数字：")
	reader := bufio.NewReader(os.Stdin)
	for {
		for {
			fmt.Println("\t1. 你来猜（数字无重复）")
			fmt.Println("\t2. 让我猜（数字无重复）")
			fmt.Println("\t3. 你来猜（数字有重复）")
			fmt.Println("\t4. 让我猜（数字有重复）")
			fmt.Println("\t5. 退出游戏")
			fmt.Print("请选择游戏模式(1-5): ")
			data, _, _ := reader.ReadLine()
			gameType, _ = strconv.Atoi(string(data))
			if gameType == 1 {
				fmt.Println("你来猜我的数字（数字无重复）!")
				numberGame(noDuplicateNumbersSet, 10)
				break
			} else if gameType == 2 {
				fmt.Println("我来猜你的数字（数字无重复）!")
				guessNumber(noDuplicateNumbersSet)
				break
			} else if gameType == 3 {
				fmt.Println("你来猜我的数字（数字有重复）!")
				numberGame(duplicateNumbersSet, 15)
				break
			} else if gameType == 4 {
				fmt.Println("我来猜你的数字（数字有重复）!")
				guessNumber(duplicateNumbersSet)
				break
			} else if gameType == 5 {
				os.Exit(0)
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

func numberGame(numbersSet []string, maxGuessTimes int) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	number := numbersSet[rand.Intn(len(numbersSet))]
	for i := 0; i < maxGuessTimes; i++ {
		fmt.Printf("请输入你的猜测(还剩%d次机会): ", maxGuessTimes-i)
		data, _, _ := reader.ReadLine()
		if match, _ := regexp.Match("^[0-9]*$", data); len(data) != 4 || !match {
			fmt.Println("输入非法，请重新输入。")
			i--
			continue
		}
		guess := string(data)
		a, b := helper.Compare(number, guess)
		if a == 4 {
			fmt.Printf("恭喜你在第%d次猜出了这个数字！\n", i+1)
			break
		} else {
			fmt.Printf("本次猜测结果为: %dA%dB\n", a, b)
			if i == maxGuessTimes-1 {
				fmt.Printf("很抱歉，你没能在%d次以内猜到数字。正确答案是: %s\n", maxGuessTimes, number)
			}
		}
	}
}

func guessNumber(numbersSet []string) {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	guessNumberSets := numbersSet
	times := 0
	var (
		guessCaches []string
		stateCaches []string
	)
	for {
		var set = make(map[string][]string)
		guess := guessNumberSets[rand.Intn(len(guessNumberSets))]
		guessCaches = append(guessCaches, guess)
		for _, v := range guessNumberSets {
			a, b := helper.Compare(guess, v)
			key := strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
			set[key] = append(set[key], v)
		}
		times++
		fmt.Println("我猜你的数字是:", guess)
		fmt.Printf("第%d次猜测结果为: ", times)
		data, _, _ := reader.ReadLine()
		state := strings.ToUpper(string(data))
		stateCaches = append(stateCaches, state)
		if state == "4A0B" {
			if times <= 5 {
				fmt.Printf("这个数字太简单了，我在第%d次就猜出来了！\n", times)
			} else if times < 10 {
				fmt.Printf("尽管你的数字很难猜，我还是在第%d次把它猜出来了！\n", times)
			} else {
				fmt.Printf("不得不说，你差点难倒我了，很难想象我居然猜了%d次才找到正确结果！\n", times)
			}
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
						checkAnswer(guessCaches, stateCaches, answer)
						return
					}
				} else {
					guessNumberSets = set[state]
					break
				}
			}
		}
	}
}

func checkAnswer(guesses, states []string, answer string) {
	if len(guesses) != len(states) {
		fmt.Println("猜测次数和给定的状态数目不匹配，无法检测！")
		return
	}
	for i, v := range guesses {
		a, b := helper.Compare(answer, v)
		state := fmt.Sprintf("%dA%dB", a, b)
		if state != states[i] {
			fmt.Printf("在第%d次猜测中，我的数字是：%s\n你给出的结果是:%s\n但我认为结果应该是:%s\n这可能是我没猜出来的原因\n", i+1, v, states[i], state)
			return
		}
	}
	fmt.Println("好吧，我没法找出自己失败的原因，你实在是太厉害了！！！")
}
