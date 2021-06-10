package bridge

import "testing"

func TestBlue_GetColor(t *testing.T) {
	rectangle := &Rectangle{}
	rectangle.SetColor(&Green{})
	rectangle.Draw()

	triangle := &Triangle{}
	triangle.SetColor(&Red{})
	triangle.Draw()

	round := Round{}
	round.SetColor(&Blue{})
	round.Draw()
}
