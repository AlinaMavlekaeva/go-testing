package greet

import "testing"

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got: %v, want: %s", got, want)
	}
}
func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if (got != nil) != want {
		if want {
			t.Fatal("Expected error, got nil")
		}
		t.Fatalf("Unexpected error: %v", got)
	}
}
func TestHello(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		got, err := Hello("Go")
		assertError(t, err, false)
		assertString(t, got, "Hello, Go")
	})

	t.Run("empty", func(t *testing.T) {
		_, err := Hello("")
		assertError(t, err, true)
	})

	t.Run("unicode", func(t *testing.T) {
		got, err := Hello("Гофер")
		assertError(t, err, false)
		assertString(t, got, "Hello, Гофер")
	})

	t.Run("spaces", func(t *testing.T) {
		got, err := Hello(" Go  ")
		assertError(t, err, false)
		assertString(t, got, "Hello,  Go  ")
	})
}
