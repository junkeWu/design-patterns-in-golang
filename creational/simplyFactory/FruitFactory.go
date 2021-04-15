package simplyFactory

import (
	"errors"
	"fmt"
)
type Factory struct {}

type FruitFactory interface {
	Create()
}

type Apple struct {
	color string
	Size  int
}

type Bear struct {
	color string
	Size  int
}

func (a *Apple)Create() {
	fmt.Println("produce a apple")
}

func (a *Bear)Create() {
	fmt.Println("produce a bear")
}

func (f *Factory) NewFruit(s string) (FruitFactory,error) {
	switch s {
		case "apple":
			return &Apple{}, nil
		case "bear":
			return &Bear{}, nil
		default:
			return nil, errors.New("fruit: this fruit is not exiting")
	}
}