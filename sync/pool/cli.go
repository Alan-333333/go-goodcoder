package main

import (
	"fmt"
	"sync"
	"time"
)

// 对象池
type BufferPool struct {
	pool  sync.Pool
	queue chan []byte
}

var bufferSize = 1024
var numBuffers = 10

func NewBufferPool() *BufferPool {

	pool := &BufferPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 0, bufferSize)
			},
		},
		queue: make(chan []byte, numBuffers),
	}

	for i := 0; i < numBuffers; i++ {
		pool.queue <- make([]byte, bufferSize)
	}

	return pool
}

func (p *BufferPool) Get() []byte {
	select {
	case buf := <-p.queue:
		return buf
	default:
		return p.pool.Get().([]byte)
	}
}

// 尝试获取,立即返回
func (p *BufferPool) TryGet() []byte {
	if buf := p.pool.Get(); buf != nil {
		return buf.([]byte)
	} else {
		return nil
	}
}

func (p *BufferPool) Put(buf []byte) {
	select {
	case p.queue <- buf:
		// 加入队列
	default:
		// 队列已满,归还给池
		p.pool.Put(buf)
	}
}

// 使用示例
func main() {
	pool := NewBufferPool()

	// 模拟多个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			buf := pool.Get()
			fmt.Printf("%+v", buf)
			for i := 0; i < 10; i++ {
				buf = append(buf, '1')
			}
			// 使用buf
			pool.Put(buf)
		}()
	}

	// 等待结束
	time.Sleep(time.Second)
}

/// 解决方案：细化多个对象池
// var bufPool = sync.Pool{}

// var imgPool = sync.Pool{}

// func getBuf() []byte {
// 	return bufPool.Get().([]byte)
// }

// func getImg() image.Image {
// 	return imgPool.Get().(image.Image)
// }
