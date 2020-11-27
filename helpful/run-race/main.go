package main

import (
	"fmt"
	"sync"
	"time"
)

// 10 个人跑步比赛
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			cond.L.Lock()
			fmt.Println(num, "号已经就位")
			cond.Wait() // 先解锁，挂起，等唤醒信号，再锁上，然后向下执行
			fmt.Println(num, "号开始跑...")
			cond.L.Unlock() // 然后再解琐
		}(i)
	}
	time.Sleep(time.Second * 2)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家开始跑")
		cond.Broadcast()
	}()
	wg.Wait()
}
