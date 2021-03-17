package main

import (
	"fmt"
	"strings"
)

const Pi = 3.14
const (
	Username = "test_user"
	Userpass = "test_pass"
)

func getOsname() string {
	return "windows"
}

// func add(x int, y int) int { // x, y int
// 	return x + y
// }

//  関数addを使用して、3+1, 3-1の結果を表示
func add(x int, y int) (int, int) { // x, y int
	return x + y, x - y
}

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
	r1, r2 := add(3, 1)
	fmt.Println(r1, r2)

	// r := add(1, 2)
	// fmt.Println(r)

	var i int = 1
	fmt.Println(i)

	xi := 1
	fmt.Println(xi)

	fmt.Println("HEllO")
	l := []string{"python", "go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for i, v := range l {
		fmt.Println(i, v)
	}

	os := getOsname()
	switch os {
	case "mac":
		fmt.Println("mac")
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("ddd")
	}

	// t := time.Now()
	// fmt.Println(t.Hour())
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("morning")
	// case t.Hour() < 17:
	// 	fmt.Println("afternoon")
	// }

	x := 0
	fmt.Println(x)
	x++
	x--

	// 実行結果が「3 + 1 = 4」
	fmt.Println("3 + 1 =", 3+1)
	// 実行結果が「3 - 1 = 2」
	fmt.Println("3 - 1 =", 3-1)
	// 実行結果が「3 / 1 = 3」
	fmt.Println("3 / 1 =", 3/1)
	// 30を4で割った時の余り
	fmt.Println(30 % 4)

	// 「Hello Go」を表示(二通り)
	fmt.Println("Hello Go")
	fmt.Println("Hello " + "Go")
	// 「Hello Go」2番目の文字「e」を表示
	fmt.Println(string("Hello Go"[1]))
	// 「Hello Go」を「Hello go」に変更
	s := "Hello Go"
	s = strings.Replace(s, "G", "g", 1)
	fmt.Println(s)
	// =>sに代入されている「Hello Go」の"G"を"g"に置き換える。最初の一個だけ
	// sの文字列に「Hello」が含まれているか？ => true or false
	fmt.Println(strings.Contains(s, "Hello"))

	// trueとfalseの型を表示
	t, f := true, false
	fmt.Printf("%T", t)
	fmt.Printf("%T", f)

	var a [2]int
	a[0] = 500
	a[1] = 1000
	fmt.Println(a)

	var b [2]int = [2]int{500, 1000}
	fmt.Println(b)
	// 配列はサイズの変更不可 => appendができない

	// スライスはサイズの変更可 => appendができる
	var c []int = []int{500, 1000}
	c = append(c, 1500)
	fmt.Println(c)

	// 1~5の数字を配列とスライスの二通りで表示
	num1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(num1)
	// var num [5]int = [5]int{1, 2, 3, 4, 5}
	// fmt.Println(num)

	num3 := []int{1, 2, 3, 4, 5}
	fmt.Println(num3)
	// var num2 []int = []int{1, 2, 3, 4, 5}
	// fmt.Println(num2)

	// 2~4を表示
	fmt.Println(num3[1:4])
	// 3までの数字を表示
	fmt.Println(num3[:3])
	// 3以降の数字を表示
	fmt.Println(num3[2:])

	//num3に100を追加
	num3 = append(num3, 100)
	fmt.Println(num3)
	// num3の1を10に変更
	num3[0] = 10
	fmt.Println(num3)

	// map
	m := map[string]int{"coffee": 350, "frape": 500, "cocoa": 400}
	fmt.Println(m)
	fmt.Println("cocoa")
	m["coffee"] = 370
	fmt.Println(m)
	m["seasonlate"] = 450
	fmt.Println(m)

	// Q1. 以下の1.11をint型に変換して出力してください。f := 1.11
	// f := 1.11
	// i := int(f)
	// fmt.Printf("%T %v", i, i)

	m2 := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Printf("%T %v", m2, m2)

	// 1~10
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("continue")
			continue
		}
		if i == 9 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	// fmt.Println("hello")
	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")
	// fmt.Println("GO")

	var n int = 100
	fmt.Println(n)
	fmt.Println(&n)

	var p *int = &n
	fmt.Println(p)
	fmt.Println(*p)

	s1 := make([]int, 0)
	fmt.Printf("%T\n", s1)
	m1 := make(map[string]int)
	fmt.Printf("%T\n", m1)

	var p1 *int = new(int)
	fmt.Printf("%T\n", p1)

}
