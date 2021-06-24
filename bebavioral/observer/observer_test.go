package observer

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	shirtItem := NewItem("Nike Shirt")

	observerFirst := &Customer{id: "abc@gmail.com"}
	observerSecond := &Customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
	for i, observer := range shirtItem.observerList {
		fmt.Printf(fmt.Sprintf("%v %v", i, observer))
	}

}
