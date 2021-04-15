package abstractFactory

import (
	"fmt"
)

type Apple struct {}

type Bear struct {}

type Banana struct {}

type AppleFactory struct {}

type BearFactory struct {}

type BananaFactory struct {}

type Factory struct {}

type Fruit interface {
	eat()
}

type FruitFactory interface {
	create() Fruit
}

func (a *AppleFactory) create() Fruit {
	fmt.Println("appleFactory: create a apple")
	return &Apple{}
}

func (b *BearFactory) create() Fruit {
	fmt.Println("bearFactory: create a bear")
	return &Bear{}
}
func (b *BananaFactory) create() Fruit {
	fmt.Println("bananaFactory: create a banana")
	return &Banana{}
}

func (a *Apple) eat() {
	fmt.Println("Apple: eating apple")
}

func (b *Bear) eat() {
	fmt.Println("Bear: eating bear")
}

func (b *Banana) eat() {
	fmt.Println("Banana: eating banana")
}


func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) CreatAppleFactory() FruitFactory {
	return &AppleFactory{}
}

func (f *Factory) CreatBearFactory() FruitFactory {
	return &BearFactory{}
}

func (f *Factory) CreatBananaFactory() FruitFactory {
	return &BananaFactory{}
}

