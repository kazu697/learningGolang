package main

import "fmt"

type Person struct {
	name string
	age int
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
		name: "bob",
		age: 24,
		height: 169,
	}

	fmt.Printf("構造体の表示%#v\n", p)
}