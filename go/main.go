package main  // 包声明
// go 是一个天生支持多核计算的云开发时代 C语言
// node + go   服务器有关
// require('fs') fileSystem

import  (
	"fmt"  // 向命令行输出hello world    format
)  // es6
// node 脚本语言  实时编译
// go c 二进制文件
// 编译原理器  func => function

func minEatingSpeed(piles []int, H int) int {
	low := 1 //let low = 1
	hi := maxPiles(piles)
	// 函数是组织代码的最小单元 { 块级作用域 }
	// maxPiles(piles [])

	for {
		if low > hi {
			break;
		}
		if canEatAllBanas(piles, H, low) {
			return low
		} else {
			low++
		}
	}

	return low
}

func canEatAllBanas(piles []int, H int, low int) bool {

	return true
}

func maxPiles(piles []int) int {
	h := 0
	for i := 0; i < len(piles); i++ {
		if pile[i] > h {
			h = pile[i]
		}
	}
	return h
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


func main()  {  // func 函数关键字
	// 从main开始
	// go 是一个静态， 类型约束
	fmt.Printf("%d小时内吃完香蕉的最慢速度是%d只/小时", 8, minEatingSpeed([]int{3, 6, 7, 11}, 8))
}