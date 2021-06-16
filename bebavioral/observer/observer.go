package observer

import (
	"fmt"
	"sync"
	"time"
)

type Event struct {
	Data int
}

// 观察者接口
type Observer interface {
	NotifyCallback(event Event)
}

// 被观察者接口
type Observable interface {
	AddListener(observer Observer)
	RemoveListener(observer Observer)
	Notify(event Event)
}

type eventObserver struct {
	ID   int
	Time time.Time
}
type eventObservable struct {
	Observers sync.Map
}

func (e eventObserver) NotifyCallback(event Event) {
	fmt.Printf("Recieved: %d after %v\n", event.Data, time.Since(e.Time))
}

func (e *eventObservable) AddListener(obs Observer) {
	e.Observers.Store(obs, struct{}{})
}

func (e *eventObservable) RemoveListener(obs Observer) {
	e.Observers.Delete(obs)
}
func (e *eventObservable) Notify(event Event) {
	e.Observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).NotifyCallback(event)
		return true
	})
}

func Fib(n int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i, j := 0, 1; i < n; i, j = i+j, i {
			out <- i
		}
	}()
	return out
}
