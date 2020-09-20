package day2_pipeline

import (
	"fmt"
	"time"
)

/*
 · 3个协程分别数数
 · 要求主协程恰好在所有协程结束时结束，目前都是永久sleep
*/

func Count(n int, grName string, chanQuit chan string) {
	for i := 0; i < n; i++ {
		fmt.Println(grName, i)
		time.Sleep(200 * time.Millisecond) //一秒读五个数
	}

	fmt.Println(grName, ":工作完毕！")

	//2，向【任务完毕通知管道】写入数据
	chanQuit <- grName + ":mission completed!"
}

func Go12() {

	//1，创建一个3缓存的【任务完毕通知管道】
	chanQuit := make(chan string, 3)

	go Count(10, "son", chanQuit)
	go Count(70, "daughter", chanQuit)
	Count(5, "main", chanQuit) //数不完， 需要正正好好数完了再over

	//3，阻塞等待从【任务完毕通知管道】读出所有协程的完毕消息
	for i := 0; i < 3; i++ {
		x := <-chanQuit //读到的东西是什么 其实不重要了，重要的是要读够三次
		fmt.Println(x)
	}
	fmt.Println("over")
}
func Go13() {
	var a *int
	// i := 42
	// a = &i
	fmt.Println(a)  //这个指针为nil ，不能寻址
	fmt.Println(&a) //0xc000006028
	// *a = 10

	// fmt.Println(*a) //panic: runtime error: invalid memory address or nil pointer dereference
}
