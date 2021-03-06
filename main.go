package main

import (
	"fmt"

	"github.com/kazu697/learningGolang/app"
)

type Person struct {
	name   string
	age    int
	height int
}

func main() {
	// 変数の宣言方法
	i := 9
	// ポインタの宣言方法
	pointer := &i
	fmt.Println("Hello World!")
	fmt.Printf("変数の表示： %d\n", i)
	fmt.Printf("アドレスの表示： %p\n", pointer)
	fmt.Printf("ポインタが指している値： %v\n", *pointer)

	// 構造体に値を入力
	p := Person{
		name:   "bob",
		age:    24,
		height: 169,
	}

	fmt.Printf("構造体の表示%#v\n", p)

	// 関数実行
	fmt.Printf("関数の返り値表示：%v\n", addOne(1))

	printf("method形式")

	array()

	forLoop(10)

	ifsentence("aaa")

	p.height = p.changeHeight(10)

	fmt.Printf("高さ変更： %#v\n", p)

	s := []int{1, 2, 3, 4, 5}
	fmt.Println(app.Average(s))

	// packageからApp Structを呼び出す
	app1 := app.App{Num: 1}
	fmt.Println(app1)

	fmt.Println(app.Public)
	// 先頭を小文字にすると呼び出せなくなる
	// fmt.Println(app.private)
}

func addOne(i int) int {
	return i + 1
}

func printf(word string) {
	fmt.Printf("%s\n", word)
}

func array() {
	// 配列の宣言方法
	var arr [2]string
	arr2 := [...]string{"array", "start"}
	arr[0] = "Golang"
	arr[1] = "learning"
	fmt.Printf("配列1: %v\n", arr)
	fmt.Printf("配列2: %v\n", arr2)

	// スライスの生成
	// スライスは後で追加ができる
	var slice []string = []string{"Golang", "slice"}
	fmt.Printf("slice: %v\n", slice)
	newSlice := append(slice, "append")
	fmt.Printf("newSlice: %v\n", newSlice)

	slicePart := newSlice[:2]
	fmt.Printf("slicePart: %v\n", slicePart)
	fmt.Printf("slicePart length: %v\n", len(slicePart))
	fmt.Printf("slicePart capacity: %v\n", cap(slicePart))

	var sliceNil []int
	if sliceNil == nil {
		fmt.Println(nil)
	}

}

func forLoop(loopNum int) {
	sum := 0
	for i := 0; i < loopNum; i++ {
		sum += 1
	}
	fmt.Println(sum)

	arrayWord := []string{"Golang", "forLoop", "Test"}
	for a, b := range arrayWord {
		fmt.Println(a, b)
	}
}

func ifsentence(arg string) {
	if v := "Go"; v == arg {
		fmt.Println(v)
	} else {
		fmt.Println("not Go")
	}
}

// 構造体の中に関数を定義する
func (p *Person) changeHeight(h int) int {
	return p.height + h
}
