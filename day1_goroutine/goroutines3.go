package day1_goroutine

import (
	"fmt"
	"runtime"
	"time"
)

//协程之间的调度，默认随机乱序公平的竞争资源

//出让协程资源 通过 runtime.Gosched()出让资源，让其他协程优先执行 ,这里调低五号协程的优先级
func Go3() {
	for i := 0; i < 10; i++ {
		// i0 := i 都行
		go func(no int) {
			for j := 1; j < 10; j++ {
				if no == 5 {
					runtime.Gosched() //将5号协程挂到 日程里去，自己挂起，如让当前协程资源,优先级由cpu自己决定
				}
				fmt.Printf("协程%d: %d\n", no, j)
				//有报错 ,都是协程2，解决方法在下面
			}
		}(i)
	}
	// effectiveFinalForGo()

	time.Sleep(2 * time.Second)
}

//关于协程引入外部变量 的坑 loop variable i captured by func literal
func effectiveFinal() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			/*
				10
				10
				10
				10
				10
				10
				10
				10
				10*/
		}()
	}
}

//可以看到他的执行结果大家基本都输出10。其实原因也很容易解释：
/*
主协程的循环很快就跑完了，而各个协程才开始跑，此时i的值已经是10了，所以各协程都输出了10。
（输出7的两个协程，在开始输出的时候主协程的i值刚好是7，这个结果每次运行输出都不一样）

出现这个问题最主要的原因是Golang中允许启动的协程中引用外部的变量。
Java对这类问题的解决方式比较合理，它也允许异步任务引用外部变量，
但是要求外部变量必须是final或者是effective final的[1]。
*/
//Go代码也能改成类似的代码使运行出正确的结果
func effectiveFinalForGo() {
	for i := 0; i < 10; i++ {
		i0 := i
		go func() {
			fmt.Println(i0)
		}()
	}
	//这里就没有警告了
}
