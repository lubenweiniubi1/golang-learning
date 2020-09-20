package day2_pipeline

import (
	"fmt"
	"strconv"
	"time"
)

/*
   管道通信
  channel介绍了channel提供一种通信机制，通过它，一个goroutine 可以向另一个goroutine发送消息，
  channel本身还需要关联一个类型，也就是channel可以发送的数据类型。例如：发送int类型消息的channel
  写成 chan int.
  创建 channel管道使用内置的make函数创建，下面生命一个chan int类型的channel
  ch:= make(chan int)
  第二个参数为缓存容量，当一个管道没有设置缓存，那么你如果想要写进去，必须有一个人同时在收，否则一直
  阻塞，同理对接受也一样，屁眼堵住

  和map类似，make创建了一个底层数据结构的引用，当复制或参数传递时，只是拷贝了一个channel引用，指向
  相同的channel对象，和其他引用类型一样，channel的空值为nil，使用 ==可以对类型相同的channel进行比较
  只有指向相同对象或同为nil时，才返回true
*/

/*管道的创建与读写*/
func Go7() {
	var mychan chan int
	fmt.Println(mychan) //nil

	//创建缓存能力为0的管道
	mychan = make(chan int)
	fmt.Println(mychan) //0xc0000180c0

	//因为没有缓存能力
	// mychan <- 123 //fatal error: all goroutines are asleep - deadlock!

	mychan = make(chan int, 1)
	mychan <- 123 //能正常运行了

	//管道已满，写不进去，主协程永远阻塞-----deadlock
	mychan <- 124

	x := <-mychan
	fmt.Println(x) //123

	fmt.Println("main over")
}

/*管道读写*/
func Go8() {
	myChan := make(chan string, 3)

	go func() {
		for i := 0; i < 5; i++ {
			//已经读空，没人写，就读出阻塞
			x := <-myChan
			fmt.Println("读出", x)
		}
	}()

	for i := 0; i < 3; i++ {
		//缓存已满，没人读，写入阻塞
		myChan <- strconv.Itoa(i)
		fmt.Println(i, "已写如")
		time.Sleep(time.Second) //就算是不加入，也不会报死锁
	}
	for {
		time.Sleep(1 * time.Second)
	}
}

//管道关闭 ,golang 提供了内置的close函数对channel进行关闭操作
// ch:= make (chan int)
// close(ch)

/*
 有关channel关闭，你需要注意一下事项：

 1.关闭一个未初始化（nil）的channel会产生panic
 2.重复关闭同一个channel会产生panic
 3.向一个已关闭的channel中发送消息会产生panic
 4，从已关闭的channel读取消息不会产生panic，且能读出channel中还未被读取的信息，
   若消息均已读出，则会读到类型的零值。
 5，从一个已经关闭的channel中读取消息永远不会阻塞，并且会返回一个为false的ok标识，可以用它判断是否关闭
 6，关闭channel会产生广播机制，所有向channel读取消息的goroutine都会收到消息
*/

func Go9() {
	bools := make(chan bool, 4)

	go func() {
		//等候管道关闭
		time.Sleep(3 * time.Second)

		/*照样读数据*/
		x := <-bools
		fmt.Println(x)

		//note 4
		x = <-bools
		fmt.Println(x)

		//note 5
		x, ok := <-bools
		fmt.Println(x, ok) //true true
		if ok {
			fmt.Println("读到正确的值", x)
		} else {
			fmt.Println("管道无法争取读出")
		}

		x, ok = <-bools    //如果缓存没被填完，当被读完之后继续读则返回 的ok 为false
		fmt.Println(x, ok) //false false
		if ok {
			fmt.Println("读到正确的值", x)
		} else {
			fmt.Println("管道无法争取读出")
		}

	}()

	bools <- true
	bools <- true
	bools <- true
	close(bools)
	//注意3
	// close(bools) //产生panic
	// bools <- true

	for {
		time.Sleep(time.Second)
	}

}

//先看 管道的遍历 ，note6
func Go10() {
	myChan := make(chan int, 5)

	go func() {
		//遍历管道
		//如果管道没有关闭，会永远尝试进行下一次读取
		//而一旦管道关闭，则阻塞读取会自动结束
		for x := range myChan {
			fmt.Println("开始读取")
			fmt.Println("读到", x)
		}
		//读取完后没有结束,如果在外层close了这里的range就结束遍历了
		fmt.Println("goroutine over")
	}()

	myChan <- 123
	myChan <- 456
	myChan <- 789

	//关闭管道，不再写入，通知读取的协程（不要再阻塞遍历了） note6
	close(myChan)

	time.Sleep(3 * time.Second)
}
