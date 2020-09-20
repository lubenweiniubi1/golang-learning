package day1_goroutine

import (
	"fmt"
	"time"
)

//百万并发
func DoTask(no int) {
	for i := 1; i < 10; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(no)
	}
}

func Go2() {
	start := time.Now()

	//能并发多少要看你cpu有几个核，百万差不多了 ，操
	for i := 1; i < 1000000; i++ {
		go DoTask(i)
	}

	time.Sleep(6 * time.Second)

	end := time.Now()
	//必须阻塞，不然主程序结束，协程就没了
	fmt.Printf("main over,startAt:%v  ,endAt:%v", start, end)
}
