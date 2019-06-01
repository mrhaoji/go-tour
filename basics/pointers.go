package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i // 指向 i
	fmt.Println(*p) // 通过指针读取 i
	*p = 21 // 通过指针设置 i
	fmt.Println(i) // 查看 i

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
