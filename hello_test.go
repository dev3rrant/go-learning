package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Saying hello to the  good people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello to the world", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in spainish", func(t *testing.T) {
		got := Hello("Chris", "Spainish")
		want := "Hola, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in french", func(t *testing.T) {
		got := Hello("Chris", "French")
		want := "Bonjour, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying hello in esperanto", func(t *testing.T) {
		got := Hello("Chris", "Esperanto")
		want := "Saluton, Chris"

		assertCorrectMessage(t, got, want)
	})

}
