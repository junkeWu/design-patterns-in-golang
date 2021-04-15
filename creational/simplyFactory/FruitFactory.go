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

type bear struct {
	color string
	Size  int
}

func (a *apple) create() {
	fmt.Println("produce a apple")
}

func (a *bear) create() {
	fmt.Println("produce a bear")
}

func (f *Factory) NewFruit(s string) (FruitFactory,error) {
	switch s {
		case "apple":
			return &apple{}, nil
		case "bear":
			return &bear{}, nil
		default:
			return nil, errors.New("fruit: this fruit is not exiting")
	}
}