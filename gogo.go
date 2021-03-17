package main

import "fmt"

func outer() {
	var s4 string = "outer"
	fmt.Println(s4)
}

func main() {

	// 変数
	var i int = 100
	fmt.Println(i)

	var s string = "Hello GO"
	fmt.Println(s)

	var (
		i2 int    = 200
		s2 string = "Golang"
	)
	fmt.Println(i2, s2)

	var i3 int    // 空
	var s3 string // 空
	fmt.Println(i3, s3)

	i3 = 300
	s3 = "Go"
	fmt.Println(i3, s3)

	i = 150
	fmt.Println(i)

	i4 := 400
	fmt.Println(i4)

	i4 = 450
	fmt.Println(i4)

	// 明示的な定義は、関数外で定義することができる
	// 暗黙的な定義は、関数外で定義することができない

	outer()

	var i5 int = 100
	fmt.Println(i5 + 50)

	// 配列
	var arr [3]int
	fmt.Println(arr)

	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	arr3 := [3]string{"A", "B", "C"}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"} // [...]勝手に要素数を数えてくれる
	fmt.Println(arr4)

	fmt.Println(arr3[0])
	fmt.Println(arr3[1])

	arr3[1] = "G"
	fmt.Println(arr3)

	// 要素数を数える
	fmt.Println(len(arr3))

	// interfaceは、全ての型との互換性を持つ
	// interfaceは、初期値にnillを持つ
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = "D"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)

	// intからfloatへ
	var g int = 1
	fl64 := float64(g)
	fmt.Println(fl64)

	fl := 10.5
	i10 := int(fl)
	fmt.Printf("i10 = %T\n", i10)
	fmt.Println(i10)

}
