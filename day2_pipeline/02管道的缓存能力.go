package day2_pipeline

import "fmt"

//管道的长度 和 缓存能力

func Go11() {
	mychan := make(chan int, 3)
	mychan <- 123

	fmt.Println("管道的长度为：", len(mychan)) //1
	fmt.Println("管道的容量：", cap(mychan))  //3

	/*已满的管道无法再进行写入*/
	mychan <- 344
	mychan <- 345
	//永远阻塞，死锁
	// mychan <- 346

	/*已经读空的管道无法再读出*/
	x := <-mychan
	x = <-mychan
	x = <-mychan

	//fatal error: all goroutines are asleep - deadlock!
	x = <-mychan
	fmt.Println("读到", x)

}
