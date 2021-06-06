package builder

import (
	"fmt"
	"testing"
)

func TestGunBuilder_Build(t *testing.T) {
	gun := NewGun("AK47","s")
	fmt.Println(gun)
}