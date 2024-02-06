package asyncgo

import (
	"fmt"
	"time"
)

func AsyncGoo(num int) {

	ch := make(chan int)

	go someAsyncSheeit(ch, num)

	result := <-ch

	fmt.Println("after: ", result)
}

func someAsyncSheeit(ch chan<- int, num int) {
	time.Sleep(time.Second)
	ch <- num
}
