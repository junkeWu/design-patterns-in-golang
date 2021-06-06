package Prototype

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConcretePrototype_Clone(t *testing.T) {
	name := "去浪"
	proto := ConcretePrototype{
		name: name,
	}
	newProto := proto.Clone()
	actualName := newProto.Name()

	assert.Equal(t, actualName, "去浪复制")
	fmt.Println(name, actualName)
}
