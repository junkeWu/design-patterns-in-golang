package factory

import "fmt"

type Apple struct {}

type Bear struct {}

type AppleFactory interface {
	Create()
}

type BearFactory interface {
	Create()
}

func (a *Apple)Create() {
	fmt.Println("appleFactory: create a apple")
}

func (a *Bear)Create() {
	fmt.Println("bearFactory: create a bear")
}


