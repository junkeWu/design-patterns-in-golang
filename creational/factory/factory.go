package factory

import (
	"fmt"
)
type Apple struct {}

type Bear struct {}

type Fruit interface {
	eat()
}

type AppleFactory struct {}

type BearFactory struct {}


type FruitFactory interface {
	create() Fruit
}

func (a *AppleFactory) create() Fruit{
	fmt.Println("appleFactory: create a apple")
	return &Apple{}
}

func (b *BearFactory) create() Fruit{
	fmt.Println("bearFactory: create a bear")
	return &Bear{}
}

func (a *Apple) eat() {
	fmt.Println("Apple: eating apple")
}


func (b *Bear) eat() {
	fmt.Println("Bear: eating bear")
}
