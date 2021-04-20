package simplyFactory

import (
	"errors"
	"fmt"
)
type Factory struct {}

type FruitFactory interface {
	create()
}

type apple struct {
	color string
	Size  int
}

type pear struct {
	color string
	Size  int
}

func (a *apple) create() {
	fmt.Println("produce a apple")
}

func (a *pear) create() {
	fmt.Println("produce a pear")
}

func (f *Factory) NewFruit(s string) (FruitFactory,error) {
	switch s {
		case "apple":
			return &apple{}, nil
		case "pear":
			return &pear{}, nil
		default:
			return nil, errors.New("fruit: this fruit is not exiting")
	}

}