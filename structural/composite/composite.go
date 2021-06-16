package composite

import (
	"container/list"
	"reflect"
	"strconv"
)

type Employee struct {
	Name         string
	Dept         string
	Salary       int
	Subordinates *list.List
}

func NewEmployee(name, dept string, salary int) *Employee {
	sub := list.New()
	return &Employee{
		Name:         name,
		Dept:         dept,
		Salary:       salary,
		Subordinates: sub,
	}
}
func (e *Employee) Add(emp Employee) {
	e.Subordinates.PushBack(emp)
}

func (e *Employee) Remove(emp Employee) {
	for i := e.Subordinates.Front(); i != nil; i = i.Next() {
		if reflect.DeepEqual(i.Value, emp) {
			e.Subordinates.Remove(i)
		}
	}
}

func (e *Employee) GetSubordinate() *list.List {
	return e.Subordinates
}

func (e *Employee) ToString() string {
	return "[Name : " + e.Name + ", Dept: " + e.Dept + ", Salary: " + strconv.Itoa(e.Salary) + "]"
}
