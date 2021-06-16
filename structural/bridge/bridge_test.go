package bridge

import (
	"fmt"
	"testing"
)

func TestDraw(t *testing.T) {
	rectangle := &Rectangle{}
	rectangle.SetColor(&Green{})
	rectangle.Draw()

	triangle := &Triangle{}
	triangle.SetColor(&Red{})
	triangle.Draw()
	round := Round{}
	round.SetColor(&Blue{})
	round.Draw()

	square := &Square{}
	square.Draw()
	fmt.Print("给平方形加个颜色:")
	square.SetColor(&Red{})
	square.Draw()
}
