package main

import (
	"fmt"
	"sync"
	"time"
)

// 用互斥锁实现对计数器的保护
var count int

var mutex sync.Mutex

var rwmutex sync.RWMutex

func increment() {
	mutex.Lock()
	defer mutex.Unlock()
	count++
}

func decrement() {
	mutex.Lock()
	defer mutex.Unlock()
	count--
}

func getCount() int {
	rwmutex.RLock()
	defer rwmutex.RUnlock()
	return count
}

func setCount(num int) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	count = num
}

var m = make(map[string]int)

func read(key string) int {
	mutex.Lock()
	defer mutex.Unlock()
	return m[key]
}

func write(key string, value int) {
	mutex.Lock()
	defer mutex.Unlock()
	m[key] = value
}

func main() {
	// 使用互斥锁保护对计数器的访问
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()
	fmt.Println("count:", count)

	// 重置
	count = 0
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			increment()
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)

	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			count--
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)

	count = 1000
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			decrement()
		}()
	}
	wg.Wait()
	fmt.Println("count:", count)
	// // 使用读写锁分离读和写

	wg.Add(3)
	go func() {
		defer wg.Done()
		count = 20
	}()

	fmt.Println("count:", count)

	go func() {
		defer wg.Done()
		count = 10
	}()

	fmt.Println("count:", count)

	go func() {
		defer wg.Done()
		count = 5
	}()

	fmt.Println("count:", count)

	wg.Wait()

	wg.Add(4)
	go func() {
		defer wg.Done()
		setCount(10)
	}()

	go func() {
		defer wg.Done()
		setCount(20)
	}()

	go func() {
		defer wg.Done()
		result := getCount()
		fmt.Println("count:", result)
	}()

	go func() {
		defer wg.Done()
		result := getCount()
		fmt.Println("count:", result)
	}()

	wg.Wait()

	// // 使用互斥锁保护map
	go write("a", 1)
	go func() {
		fmt.Printf("map:%+v\n", m)
	}()
	go write("b", 2)
	go func() {
		fmt.Printf("map:%+v\n", m)
	}()

	go func() {
		m["a"] = 1
	}()

	go func() {
		m["a"] = 2
	}()

	go func() {
		m["a"] = 3
	}()

	go func() {
		m["a"] = 4
	}()

	go func() {
		fmt.Printf("map:%+v\n", m)
	}()

	time.Sleep(1 * time.Second)
}
