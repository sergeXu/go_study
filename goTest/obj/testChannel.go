package main

import (
	"fmt"
	"time"
)

func testChannel() {
	//testBasicChannel()
	//testBufferChannel()
	testChannelClose()
	testChannelSelect()

}

func testBasicChannel() {
	c := make(chan int)
	go func() {
		fmt.Println("goroutine 1 execute")
		defer fmt.Println("goroutine 1 end")
		//time.Sleep(3 * time.Second)
		c <- 888
	}()
	time.Sleep(3 * time.Second)
	// channel作为同步点，会两边相互等待  无缓存模式
	num := <-c
	fmt.Println("num = ", num)
	time.Sleep(2 * time.Second)
	fmt.Println("main goroutine end")
}

func testBufferChannel() {
	//带缓存的channal
	c := make(chan int, 3)
	fmt.Println("len(c)= ", len(c), " cap(c) = ", cap(c))
	go func() {
		defer fmt.Println("子 goroutine  end")
		for i := 0; i < 5; i++ {
			//超过容量的数据发送会被阻塞
			c <- i
			fmt.Println("子Go程正在运行，发送元素 ", i, "len(c)= ", len(c), " cap(c) = ", cap(c))
		}
	}()
	time.Sleep(2 * time.Second)
	for i := 0; i < 5; i++ {
		//取不到值的时候也会阻塞
		num := <-c
		fmt.Println("num = ", num)
	}
	fmt.Println()
	fmt.Println("testBufferChannel end")
}

func testChannelClose() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		//控制channel状态，表示已关闭
		close(c)
	}()
	//永久循环
	//for {
	//	//ok为true表示channel没有关闭
	//	if data, ok := <-c; ok {
	//		fmt.Println("data = ", data)
	//	} else {
	//		fmt.Println("channel closed")
	//		break
	//	}
	//}
	//可以使用range来遍历操作
	for data := range c {
		fmt.Println("range data = ", data)
	}

	fmt.Println("testChannelClose end")
}
func fibonacii(c chan int, quit chan int) {
	x, y := 1, 1
	for {
		//多路channel结果监控
		select {
		//如果c可写，则可以执行
		case c <- x:
			x, y = y, x+y
		// quit队列可读
		case <-quit:
			fmt.Println("quit")
			//len为0，无缓存channel写入需要读写做握手
			fmt.Println("len(c) = ", len(c))
			return
		}
	}
}

func testChannelSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 15; i++ {
			fmt.Println(<-c)
		}
		time.Sleep(2 * time.Second)
		quit <- 0
	}()
	fibonacii(c, quit)
	fmt.Println("testChannelSelect end")
}
