package singleton

import (
	"fmt"
	"sort"
	"sync"
	"testing"
	"time"
)

var mu2 sync.Mutex

func Test_getIns(t *testing.T) {
	ch := make(chan int, 10)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instanceOnce := getInstanceHungry()
			mu2.Lock()
			instanceOnce.IncrementAge()
			age := instanceOnce.age
			// fmt.Printf("address: %p\n", instanceOnce)
			// fmt.Printf("age1: %v\n ", instanceOnce.age)
			mu2.Unlock()
			time.Sleep(time.Millisecond * 1000)
			// 把通道和锁分开，避免死锁
			ch <- age
		}()
		fmt.Printf("")
	}
	wg.Wait()
	close(ch)
	var res []int
	for v := range ch {
		res = append(res, v)
	}

	// var res []int
	// // 取出通道数据
	// for v := range ch {
	//	res = append(res, v)
	// }
	// 顺序打印
	sort.Ints(res)
	for _, v := range res {
		fmt.Println(v)
	}
}

func Test_getInstanceLazy(t *testing.T) {
	ch := make(chan int, 100)
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instanceOnce := getInstanceLazy()
			mu2.Lock()
			instanceOnce.IncrementAge()
			age := instanceOnce.age
			// fmt.Printf("address: %p\n", instanceOnce)
			// fmt.Printf("age1: %v\n ", instanceOnce.age)
			mu2.Unlock()
			time.Sleep(time.Millisecond * 100)
			ch <- age
		}()
		fmt.Printf("")
	}
	wg.Wait()
	close(ch)
	var res []int
	// 取出通道数据
	for v := range ch {
		res = append(res, v)
	}
	// 顺序打印
	sort.Ints(res)
	for _, v := range res {
		fmt.Println(v)
	}
}
