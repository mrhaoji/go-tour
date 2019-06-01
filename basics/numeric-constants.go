package main

import "fmt"

const (
	Big = 1 << 100 // 1 左移 100 位创建一个非常大的数字
	Small = Big >> 99 // 再右移 99 位
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt((Small)))
	fmt.Println(needFloat((Small)))
	fmt.Println(needFloat((Big)))
}
