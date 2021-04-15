package factory

import "testing"

func TestCreate(t *testing.T) {
	apple := Apple{}
	apple.Create()
	//
	bear := Bear{}
	bear.Create()
}

