package factory

import (
	"fmt"
)
type Apple struct {}

type Pear struct {}

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
	fmt.Println("bearFactory: create a pear")
	return &Pear{}
}

func (a *Apple) eat() {
	fmt.Println("Apple: eating apple")
}


func (b *Pear) eat() {
	fmt.Println("Pear: eating pear")
}
