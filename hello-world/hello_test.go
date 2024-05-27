package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("Testing Hello World", func(t *testing.T) { // Passing in functions to a function is really nice I feel like I'm writing JS
		got := Hello("Chris", "English")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Testing Hello World with empty string", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)

	})
	t.Run("Testing in Zulu", func(t *testing.T) {
		got := Hello("Sizwe", "Zulu")
		want := "Kodwa uyinja yaz wena, Sizwe"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Testing in Spanish", func(t *testing.T) {
		got := Hello("Sizwe", "Spanish")
		want := "Hola, Sizwe"

		assertCorrectMessage(t, got, want)
	})
	t.Run("Testing in Zulu", func(t *testing.T) {
		got := Hello("Sizwe", "French")
		want := "Bonjour, Sizwe"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) { // This is NICE I can type mutliple variables at once, this is really niceeeee
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
