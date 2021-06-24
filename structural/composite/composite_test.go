package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	ceo := NewEmployee("1", "ceo", 100000000)
	pm := NewEmployee("技术", "technology", 10000)
	programmer1 := NewEmployee("程序员1", "technology", 8000)
	programmer2 := NewEmployee("程序员2", "technology", 8000)
	minister := NewEmployee("部长", "accounting", 5000)
	finance1 := NewEmployee("会计1", "accounting", 3000)
	finance2 := NewEmployee("会计2", "accounting", 2900)
	tranee := NewEmployee("实习生", "tranee", 1000)
	tranee2 := NewEmployee("实习生带实习生", "tranee2", 1000)


	ceo.Add(*pm)
	ceo.Add(*minister)

	pm.Add(*programmer1)
	pm.Add(*programmer2)

	minister.Add(*finance1)
	minister.Add(*finance2)

	finance1.Add(*tranee)

	tranee.Add(*tranee2)
	//////打印所有职员
	//fmt.Println("[ceo:]" + ceo.ToString())
	//for i := ceo.Subordinates.Front(); i != nil; i = i.Next() {
	//	em := i.Value.(Employee)
	//	fmt.Println(em.ToString())
	//	for j :=em.Subordinates.Front(); j != nil; j = j.Next() {
	//		em := j.Value.(Employee)
	//		fmt.Println(em.ToString())
	//	}
	//}
	fmt.Println("[ceo:]" + ceo.ToString())
	PrintTree(ceo)
}
