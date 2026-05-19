package safe

import "testing"

var xs = []int{1, 2, 3}

func TestMustAt(t *testing.T) {
	got := MustAt(xs, 0)
	want := 1
	if got != want {
		t.Fatalf("MustAt([]int{1, 2, 3}, 0): got = %v, want = %v", got, want)
	}
}
func TestMustAt_BelowZero_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Ожидалась паника")
		} else if r != "index out of range" {
			t.Errorf("Неожиданное сообщение паники: %v", r)
		}
	}()
	_ = MustAt(xs, -1)
}
func TestMustAt_OutOfRange_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("Ожидалась паника")
		} else if r != "index out of range" {
			t.Errorf("Неожиданное сообщение паники: %v", r)
		}
	}()
	_ = MustAt(xs, 3)
}
