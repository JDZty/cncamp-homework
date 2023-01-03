// 基于 Channel 编写一个简单的单线程生产者消费者模型：
// 队列：
//
//	队列长度 10，队列元素类型为 int
//
// 生产者：
//
//	每 1 秒往队列中放入一个类型为 int 的元素，队列满时生产者可以阻塞
//
// 消费者：
//
//	每一秒从队列中获取一个元素并打印，队列为空时消费者阻塞
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var consumeStopFlag = false

func main() {
	wg.Add(2)
	// 创建通道
	message := make(chan int, 10)
	// 创建通道，用于20S后关闭生产者
	done := make(chan bool)
	go func() {
		time.Sleep(20 * time.Second)
		close(done)
	}()
	// 消费者逻辑
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			consume(message)
			// 消费完通道中所有信息再停止
			if consumeStopFlag {
				break
			}
		}
		wg.Done()
	}()

	// 生产者逻辑
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		count := 0
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("=========== Timeout, Producer stop. ===========")
				close(message)
				wg.Done()
				return
			default:
				produce(message, count)
				count++
			}

		}
	}()

	wg.Wait()
	fmt.Println("Main is over.")
}

func produce(message chan<- int, i int) {
	message <- i
	fmt.Println("Producer send message:", i)
}

func consume(message <-chan int) {
	m, isOpen := <-message
	if !isOpen {
		consumeStopFlag = true
		return
	}
	fmt.Println("Consumer receive message:", m)
}
