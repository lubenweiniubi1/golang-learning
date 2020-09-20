package day1_goroutine

import (
	"fmt"
	"time"
)

/*
coorperate 合作
routine 日程，事务
co-routine 协程（可以相互协作的日程）
goroutine go语言写成，go程
*/

//创建协程
func newTask() {
	for i := 0; i < 10; i++ {
		fmt.Println("老子是子协程")
		// time.Sleep(time.Second)
	}
}

func Go1() {
	//开辟一条协程 ，与主协程并发地执行newTask()
	go newTask()

	//主协程赖着不死，主协程如果死了，子协程也得陪死
	for i := 0; i < 10; i++ {
		fmt.Println("this is a main goroutine")
		time.Sleep(time.Second) //阻塞代码，如果不阻塞，协程是不会执行的,只要不阻塞，就会执行完
	}

}
