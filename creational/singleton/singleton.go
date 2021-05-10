package singleton

import (
	"sync"
)

type singleton struct{
	age int
}

var (
	instance *singleton
	mu       sync.Mutex
	ins      = &singleton{}
	once     sync.Once
)

// lazy loading
func getInstanceLazy() *singleton {
	if instance == nil {
		// 模拟多线程
		//time.Sleep(100000)
		instance = &singleton{}
		return instance
	}
	return instance
}

// hungry loading
func getInstanceHungry() *singleton {
	return ins
}

// lazy loading add lock
func getIns() *singleton {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		instance = &singleton{}
		return instance
	}
	return instance
}

// double lock
func getInstanceDoubleLock() *singleton {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()
		if instance == nil {
			instance = &singleton{}
		}
		return instance
	}
	return instance
}

// sync.Once (recommend)
func getInstanceOnce() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}

func (p *singleton) IncrementAge() {
	mu.Lock()
	p.age++
	mu.Unlock()
}