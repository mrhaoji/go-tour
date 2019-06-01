package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8  {
	a := make([][]uint8,dy)  //外层切片
	for x := range a{
		b := make([]uint8,dx)  //里层切片
		for y := range b{
			b[y] = uint8(x*y - 1)  //给里层切片里的每一个元素赋值。其中x*y可以替换成别的函数
		}
		a[x] = b  //给外层切片里的每一个元素赋值
	}
	return a
}

func main() {
	pic.Show(Pic)
}
