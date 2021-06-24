package observer

import "fmt"

// 观察者接口

type Observer interface {
	update(string)
	getID() string
}

func RemoveFromSlice(observerList []Observer,obs Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if obs.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

type Customer struct {
	id string
}

func (c *Customer) update(s string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, s)
}

func (c *Customer) getID() string {
	return c.id
}


