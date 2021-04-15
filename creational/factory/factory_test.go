package factory

import "testing"

func TestFactory(t *testing.T) {
	af := AppleFactory{}
	af.create().eat()

	bf := BearFactory{}
	bf.create().eat()

}

