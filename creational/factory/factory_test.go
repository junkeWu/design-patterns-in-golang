package factory

import "testing"

func TestFactory(t *testing.T) {
	af := AppleFactory{}
	af.create().eat()

	fp := BearFactory{}
	fp.create().eat()

}

