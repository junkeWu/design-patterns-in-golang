package Prototype

type Prototype interface {
	Name() string
	Clone() Prototype
}

type ConcretePrototype struct {
	name string
}

func (cp *ConcretePrototype) Name() string{
	return cp.name
}

func (cp *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{name: cp.name+"复制"}
}