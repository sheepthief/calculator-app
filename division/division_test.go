package division

import (
	"testing"
)

func TestDivision(t *testing.T) {

	result := Div(4, 3)

	if result != 4/3 {
		t.Error("Wrong answer")
		t.Error(result)
	}
}
