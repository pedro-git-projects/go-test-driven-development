package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Pedro", "")
		want := "Hello, Pedro"
		assertCorrectMessage(t, got, want)

	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in portuguese", func(t *testing.T) {
		got := Hello("Pedro", "Portuguese")
		want := "Olá, Pedro"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello in japanese", func(t *testing.T) {
		got := Hello("Pedro", "Japanese")
		want := "こんにちは, Pedro"
		assertCorrectMessage(t, got, want)
	})
}
