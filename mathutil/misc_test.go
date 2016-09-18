package mathutil

import(
	"testing"
)


func TestReverseInt(t *testing.T) {
	if ReverseInt(123) != 321 {
		t.Error("Reverse failed 123")
	}
	if ReverseInt(100) != 1 {
		t.Error("Reverse failed 100")
	}
	if ReverseInt(3131) != 1313 {
		t.Error("Reverse failed 3131")
	}
}

