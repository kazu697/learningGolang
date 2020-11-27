package main

import "fmt"

func main() {
	// 変数の宣言方法
	i := 9
	// ポインタの宣言方法
	pointer := &i
	fmt.Println("Hello World!")
	fmt.Printf("変数の表示： %d\n", i)
	fmt.Printf("アドレスの表示： %p\n", pointer)
	fmt.Printf("ポインタが指している値： %v\n", *pointer)
}