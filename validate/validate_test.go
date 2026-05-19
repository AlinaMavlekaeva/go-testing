package validate

import "testing"

func TestValidateName(t *testing.T) {
	err := ValidateName("Name")
	if err != nil {
		t.Errorf("ValidateName('Name'): Не ожидалась ошибка: %w", err)
	}
}
func TestValidateNameEmpty(t *testing.T) {
	err := ValidateName("")
	if err == nil {
		t.Error("Ожидалась ошибка")
	}
}
