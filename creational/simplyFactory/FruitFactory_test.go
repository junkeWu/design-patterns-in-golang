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
	apple.Create()
	bear, err := new(Factory).NewFruit("bear")
	if err != nil {
		err := fmt.Errorf("Wrap了一个错误%w", err)
		fmt.Println(errors.Unwrap(err))
		return
	}
	bear.Create()

}