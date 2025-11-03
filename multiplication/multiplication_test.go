package multiplication

import (
	"testing"
)

func TestMultiplication(t *testing.T) {

	result := Mul(4, 3)

	if result != 12 {
		t.Error("Wrong answer")
		t.Error(result)
	}
}
