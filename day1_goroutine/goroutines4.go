package day1_goroutine

import (
	"fmt"
	"runtime"
	"time"
)

//查看可用内核数

/*
  可用内核越多，并发质量越高
*/

func Go4() {
	cpu_num := runtime.NumCPU() //居然是8个！查看可用cpu核心数
	fmt.Println(cpu_num)        //8

	//把最大可使用逻辑核心数设置为1，返回先前的设置
	preMaxprocs := runtime.GOMAXPROCS(1)
	fmt.Println(preMaxprocs) //8

}

//子协程自杀
func Go5() {
	go func() {
		for i := 0; i < 10; i++ {
			if i == 5 {
				//干掉当前协程
				runtime.Goexit()
			}
			fmt.Println("goroutine", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println("main", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("main over")
}

//主协程自杀以后，子协程不会受到牵制，正常结束子协程会陪死，但是自杀不会
func Go6() {
	go func() {
		for i := 0; i < 10; i++ {

			fmt.Println("goroutine", i)
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for i := 0; i < 10; i++ {
		if i == 5 {
			//干掉当前协程
			fmt.Printf("main 被干掉了")
			runtime.Goexit()
		} //这里会报dead lock，
		fmt.Println("main", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("main over")
}
