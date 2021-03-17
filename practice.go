package main

import (
	"fmt"
)

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "no"
	}
}

func main() {
	result := by2(10)

	if result == "ok" {
		fmt.Println("great")
	}
	fmt.Println(result)

	m := map[string]int{
		"Mike":  30,
		"Nancy": 20,
		"meisa": 50,
	}
	fmt.Printf("%T %v", m, m)

	// if文
	num := 11
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// continueとbreak
	for i1 := 0; i1 < 10; i1++ {
		if i1 == 3 {
			fmt.Println("continue")
			continue
		}
		if i1 > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i1)

	}

	// sum := 1
	// for sum < 10 {
	// 	sum += sum
	// 	fmt.Println(sum)
	// }

	// Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。
	// l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var min int
	for i, num := range l {
		if i == 0 {
			min = num
			continue
		}

		if min >= num {
			min = num
		}
	}
	fmt.Println(min)

	m3 := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	sum := 0
	for _, v := range m3 {
		sum += v
	}
	fmt.Println(m3)

	// 	Q2. 以下の果物の価格の合計を出力するコードを書いてください。
	// m := map[string]int{
	//     "apple":  200,
	//     "banana": 300,
	//     "grapes": 150,
	//     "orange": 80,
	//     "papaya": 500,
	//     "kiwi":   90,

}
