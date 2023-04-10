package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := 3
	pipe := make(chan int, len(nums))
	res := 0
	ctl := make(chan int, m)
	for i := 0; i < len(nums); i++ {
		go compute(nums[i], pipe, ctl)
	}
	for len(pipe) != 0 {
		tmp := <-pipe
		res += tmp
		fmt.Printf("当前计算值%v，已完成结果%v", tmp, res)
	}
}

func compute(num int, pipe chan int, ctl chan int) {
	ctl <- 1
	doTask(num)
	<-ctl
	pipe <- doTask(num)
	// return res
}

func doTask(num int) int {
	return num
}
