package subtraction

import (
	"testing"
)

func TestSubtraction(t *testing.T) {

	result := Sub(4, 3)

	if result != 1 {
		t.Error("Wrong answer")
		t.Error(result)
	}
}
