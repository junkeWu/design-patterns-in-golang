package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	ceo := NewEmployee("1", "ceo", 100000000)
	pm := NewEmployee("2-1", "technology", 10000)
	programmer1 := NewEmployee("3-1", "technology", 8000)
	programmer2 := NewEmployee("3-2", "technology", 8000)
	minister := NewEmployee("4-1", "accounting", 5000)
	finance1 := NewEmployee("5-1", "accounting", 3000)
	finance2 := NewEmployee("5-2", "accounting", 2900)

	ceo.Add(*pm)
	ceo.Add(*minister)

	pm.Add(*programmer1)
	pm.Add(*programmer2)

	minister.Add(*finance1)
	minister.Add(*finance2)

	//打印所有职员
	fmt.Println("[ceo:]" + ceo.ToString())
	for i := ceo.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(Employee)
		fmt.Println(em.ToString())
		for j := i.Value.(Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(Employee)
			fmt.Println(em.ToString())
		}
	}

	fmt.Println("[pm:]" + pm.ToString())
	for i := pm.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(Employee)
		fmt.Println(em.ToString())
		for j := i.Value.(Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(Employee)
			fmt.Println(em.ToString())
		}
	}
	fmt.Println("[finance1:]" + finance1.ToString())
	for i := finance1.Subordinates.Front(); i != nil; i = i.Next() {
		em := i.Value.(Employee)
		fmt.Println(em.ToString())
		for j := i.Value.(Employee).Subordinates.Front(); j != nil; j = j.Next() {
			em := j.Value.(Employee)
			fmt.Println(em.ToString())
		}
	}

	front := ceo.GetSubordinate().Front()
	employee := front.Value
	employee2 := ceo.GetSubordinate().Front().Value
	fmt.Println(employee)
	fmt.Println(employee2)
	element := front.Value.(Employee).Subordinates.Front().Value
	fmt.Println(element)
}
