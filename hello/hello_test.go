package hello

import "testing"

func TestHello(t *testing.T) {
	assertString := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("hello to people", func(t *testing.T) {
		got := Hello("Chris", "English")
		want := "Hello, Chris"
		assertString(t, got, want)
	})

	t.Run("hello default world", func(t *testing.T) {
		got := Hello("", "Spanish")
		want := "Hola, World"
		assertString(t, got, want)
	})

	t.Run("hello to french", func(t *testing.T) {
		got := Hello("", "French")
		want := "Bonjour, World"
		assertString(t, got, want)
	})
}
