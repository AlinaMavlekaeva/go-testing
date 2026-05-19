package validate

import "testing"

func TestValidateNameEmpty(t *testing.T) {
	err := ValidateName("")
	if err == nil {
		t.Fatal("Ожидалась ошибка")
	} else if err != ErrEmptyName {
		t.Fatalf("Неожиданная ошибка: %v", err)
	}
}
func TestValidateName(t *testing.T) {
	err := ValidateName("Name")
	if err != nil {
		t.Fatalf("Не ожидалась ошибка: %v", err)
	}
}
