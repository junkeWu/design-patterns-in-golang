package simplyFactory

import (
	"errors"
	"fmt"
	"testing"
)

func TestNewFruitFactory(t *testing.T) {
	apple, err := new(Factory).NewFruit("apple")
	if err != nil {
		err := fmt.Errorf("Wrap了一个错误%w", err)
		fmt.Println(errors.Unwrap(err))
		return
	}
	apple.create()
	pear, err := new(Factory).NewFruit("pear")
	if err != nil {
		err := fmt.Errorf("Wrap了一个错误 %w", err)
		fmt.Println(errors.Unwrap(err))
		fmt.Println(err)
		return
	}
	pear.create()

}
