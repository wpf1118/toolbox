package chs

import (
	"fmt"
	"sync"

	"github.com/wpf1118/toolbox/tools/logging"
)

type Chs struct {
	*sync.RWMutex
	*sync.WaitGroup
	taskName  string `comment:"任务名称"`
	workerNum int    `comment:"工人数量"`
	fn        func(interface{}) (interface{}, error)
	results   []interface{}
	chs       []chan interface{}
	total     int
}

func NewChs(taskName string, workerNum int, fn func(interface{}) (interface{}, error)) *Chs {
	var RWMutex sync.RWMutex
	var waitGroup sync.WaitGroup
	return &Chs{RWMutex: &RWMutex, WaitGroup: &waitGroup, fn: fn, taskName: taskName, workerNum: workerNum}
}

// worker 处理一个通道里的数据
func (c *Chs) worker(workerNo string, ch <-chan interface{}) {
	defer c.Done()
	var n int
	var result []interface{}
	for msg := range ch {
		n++
		logging.DebugF("%s Worker: %s  正在处理第: %d ", c.taskName, workerNo, n)

		res, err := c.fn(msg)
		if err != nil {
			logging.ErrorF("%s Worker: %s Error: %v", c.taskName, workerNo, err)
			continue
		}

		if res != nil {
			result = append(result, res)
		}
	}

	c.Lock()
	c.results = append(c.results, result...)
	c.total += n
	c.Unlock()
}

// createWorkers 创建消费者
func (c *Chs) createWorkers() {
	num := c.workerNum

	for i := 0; i < num; i++ {
		ch := make(chan interface{})
		c.chs = append(c.chs, ch)
		c.Add(1)

		workerNo := fmt.Sprintf("worker-%d", i)
		go c.worker(workerNo, ch)
	}

	return
}

// Work 消费者
func (c *Chs) Work(tasks []interface{}) {
	num := c.workerNum
	// 创建消费者
	c.createWorkers()

	for k, task := range tasks {
		i := k % num

		// 阻塞，等待消费者处理
		c.chs[i] <- task
	}

	// 关闭通道，通道数据仍然可读
	for i := 0; i < num; i++ {
		close(c.chs[i])
	}

	// 等待子协程退出
	c.Wait()
}

// GetResults 返回结果
func (c *Chs) GetResults() []interface{} {
	return c.results
}

// GetTotal 返回结果总数
func (c *Chs) GetTotal() int {
	return c.total
}
