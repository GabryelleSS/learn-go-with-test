package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("should say hello to people", func(t *testing.T) {
		assertCorrectMessage(t, HelloWorld("Chris", ""), "Hello, Chris")
	})
	t.Run("should return 'Hello, World' when it has an empty string", func(t *testing.T) {
		assertCorrectMessage(t, HelloWorld("", ""), "Hello, World")
	})
	t.Run("should return 'Hola, Elodie' if the language is Spanish", func(t *testing.T) {
		got := HelloWorld("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("should return 'Bonjour, Elodie' if the language is French", func(t *testing.T) {
		got := HelloWorld("Anny", "French")
		want := "Bonjour, Anny"
		assertCorrectMessage(t, got, want)
	})

	t.Run("should return 'Olá, Gaby' if the language is Portuguese-BR", func(t *testing.T) {
		got := HelloWorld("Gaby", "Portuguese-BR")
		want := "Olá, Gaby"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}