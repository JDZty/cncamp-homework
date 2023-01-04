// 多生产者，多消费者练习
package main

import (
	"fmt"
	"sync"
	"time"
)

var pArr []Producer

var cArr []Consumer

var q = Queue{
	queue: []string{},
	cond:  sync.NewCond(&sync.Mutex{}),
}

type Queue struct {
	queue []string
	cond  *sync.Cond
}

type Producer struct {
	id   int
	name string
}

type Consumer struct {
	id   int
	name string
}

func main() {
	initData(2, 4)

	// 启动生产者
	for i, p := range pArr {
		go func(i int, p Producer) {
			for {
				msg := fmt.Sprintf("Hello, I am %s", p.name)
				p.SendMsg(msg)
				sleep(i + 1)
			}
		}(i, p)
	}

	// 启动消费者, 轮询消费消息
	go func() {
		count := 0
		clen := len(cArr)
		for {
			if count == clen {
				count = 0
			}
			consumer := cArr[count]
			consumer.ReceiveMsg()
			count++
			sleep(consumer.id + 1)
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("========= main is over =========")
}

// 生产者发送消息
func (p *Producer) SendMsg(msg string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	q.queue = append(q.queue, msg)
	fmt.Println("<<<<<<<<", p.name, "already send msg:", msg)
	q.cond.Broadcast()
}

// 消费者接受消息
// Q1: 如果消费者是多线程的，都在q.cond.Wait()阻塞，当有一个消息广播，
// 只有一个消费者可以消费，其他消费者会数组越界
func (c *Consumer) ReceiveMsg() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		q.cond.Wait()
	}

	msg := q.queue[0]
	q.queue = q.queue[1:]
	fmt.Println(c.name, "alraedy receive msg:", msg)
}

func sleep(n int) {
	var tmp time.Duration
	if n > 3 {
		tmp = 3 * time.Second
	} else {
		tmp = time.Duration(n) * time.Second
	}

	time.Sleep(tmp)
}

func initData(pNum int, cNum int) {
	// 初始化生产者列表
	for i := 0; i < pNum; i++ {
		p := Producer{}
		p.id = i
		p.name = fmt.Sprintf("Producer-%d", i)
		pArr = append(pArr, p)
	}

	// 初始化消费者者列表
	for i := 0; i < cNum; i++ {
		c := Consumer{}
		c.id = i
		c.name = fmt.Sprintf("Consumer-%d", i)
		cArr = append(cArr, c)
	}
}
