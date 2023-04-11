package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := 3
	pipe := make(chan int, len(nums))
	res := 0
	// 在go里头，管道可以带有缓冲区，当缓冲区满时。向管道写的线程会被阻塞，等待读取。
	// 当读取一个了一个数据后，此时管道的容量被释放，被阻塞的线程会继续执行。
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
}

func doTask(num int) int {
	return num
}
