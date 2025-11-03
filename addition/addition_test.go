package addition

import (
	"testing"
)

func TestAddition(t *testing.T) {

	result := Add(4, 3)

	if result != 7 {
		t.Error("Wrong answer")
		t.Error(result)
	}
}
