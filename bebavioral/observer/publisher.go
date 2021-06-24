package observer

import "fmt"

// 主题服务接口
type subject interface {
	register(obs Observer)
	deregister(obs Observer)
	notifyAll()
}
// Item 订阅对象
type item struct {
	observerList	[]Observer
	name string
	inStock bool
}

func NewItem(name string) *item {
	return &item{
		name: name,
	}
}

func (i *item) register(obs interface{}) {
	s := obs.(Observer)
	i.observerList = append(i.observerList, s)
}

func (i *item) deregister(obs Observer) {
	i.observerList = RemoveFromSlice(i.observerList, obs)
}

func (i *item) notifyAll() {
	for _, observer := range i.observerList {
		observer.update(i.name)
	}
}
func (i *item) updateAvailability()  {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.notifyAll()
}