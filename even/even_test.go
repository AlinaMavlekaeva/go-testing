package even

import "testing"

func TestIsEven(t *testing.T) {
	if !IsEven(2) {
		t.Error("Ожидалось true для 2, получено - false")
	}
	if IsEven(3) {
		t.Error("Ожидалось false для 3, получено - true")
	}
	if !IsEven(0) {
		t.Error("Ожидалось true для 0б получено - false")
	}
	if !IsEven(-4) {
		t.Error("Ожидалось true для -4, получено - false")
	}
	if IsEven(-5) {
		t.Error("Ожидалось false для -5, получено - true")
	}
}
