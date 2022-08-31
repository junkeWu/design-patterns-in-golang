package abstractFactory

import "testing"

func TestAbstractFactory(t *testing.T) {
	factory := NewFactory(2)

	appleFactory := factory.CreatAppleFactory()
	appleFactory.create().eat()

	factory.CreatBearFactory().create().eat()

	factory.CreatBananaFactory().create().eat()
}
