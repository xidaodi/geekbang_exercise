package main

import (
	"fmt"
	"sync"
)

func main() {

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(i)
	// }
	// fullstring := "hello world"
	// fmt.Println(fullstring)
	// for i, c := range fullstring {

	// 	fmt.Println(i, string(c))

	// }

	// var a, b int = 1, 2
	// c := 3
	// fmt.Println(a, b, c)

	// var i int = 42
	// var f float32 = float32(i)
	// u := uint(f)
	// fmt.Print(i, f, u)

	// myarry := [5]int{1, 2, 3, 4, 5}
	// myslice := myarry[:4]

	// fmt.Print("myslice%+v\n", myslice)

	// rmvslice := deleteitem(myslice, 2)
	// fmt.Print("myslice%+v\n", rmvslice)

	// slice := []int{1, 2, 3, 4, 5}
	// for k, _ := range slice {
	// 	slice[k] *= 2
	// }
	// fmt.Print(slice)

	// maptest := make(map[string]string, 10)
	// maptest["a"] = "aa"
	// maptest["b"] = "bb"
	// fmt.Print(maptest, "\n")
	// for k, v := range maptest {
	// 	fmt.Print(k + ":" + v)
	// }

	// myarray := []string{"I", "am", "stupid", "and", "weak"}
	// fmt.Print(myarray, "\n")
	// for k, _ := range myarray {
	// 	if myarray[k] == "stupid" {
	// 		myarray[k] = "smart"
	// 	} else if myarray[k] == "weak" {
	// 		myarray[k] = "strong"
	// 	}
	// }
	// fmt.Print(myarray)

	// mybook := Books{"go lauguage", "Hary", "programming", 1004}
	// fmt.Print(mybook)
	// printBooks(&mybook)

	// for i := 0; i < 10; i++ {
	// 	go fmt.Println(i)
	// }
	// time.Sleep(time.Second)

	// ch := make(chan int)
	// go func() {
	// 	fmt.Print("hello from goroutine")
	// 	ch <- 0
	// }()
	// <-ch

	// ch := make(chan int, 10)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		rand.Seed(time.Now().UnixNano())
	// 		n := rand.Intn(10)
	// 		fmt.Println("putting:", n)
	// 		ch <- n
	// 	}
	// 	close(ch)

	// }()
	// fmt.Println("hello from main")
	// for v := range ch {
	// 	fmt.Println("receiving:", v)

	// }

	//

	// ch := make(chan int, 10)

	// go func() {
	// 	t := time.NewTicker(time.Second)
	// 	for {
	// 		ch <- (<-t.C).Second()
	// 	}

	// }()
	// defer close(ch)

	// t := time.NewTicker(time.Second)
	// for {
	// 	<-t.C
	// 	fmt.Println(<-ch)
	// }

	//生产者消费者练习
	// messages := make(chan int, 10)
	// done := make(chan bool)
	// defer close(messages)

	// //consumer
	// go func() {
	// 	ticker := time.NewTicker(1 * time.Second)
	// 	for range ticker.C {
	// 		select {
	// 		case <-done:
	// 			fmt.Println("be interuppted by child process", <-done)
	// 			return
	// 		default:
	// 			fmt.Println("receive message:", <-messages)
	// 		}
	// 	}
	// }()

	// //producer
	// for i := 0; i < 10; i++ {
	// 	messages <- i
	// }

	// //5秒后主线程关闭done通道
	// time.Sleep(20 * time.Second)
	// close(done)
	//生产者消费者练习结束

	//锁练习
	// fmt.Println("lock practice start")
	// go rLock()
	// go wLock()
	// go lock()
	// time.Sleep(5 * time.Second)
	// fmt.Println("lock practice end")

	//waitgroup练习
	fmt.Println("waitgroup practice start")
	wg := sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("waitgroup practice end")

}

func lock() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("lock:", i)

	}
}

func rLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println("rlock", i)
	}

}

func wLock() {
	lock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println("wlock", i)

	}
}
