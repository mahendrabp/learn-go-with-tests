package main

import "testing"

//func TestHello(t *testing.T) {
//	got := Hello("Bima")
//	want := "Hello, Bima"
//
//	if got != want {
//		t.Errorf("got %q want %q", got, want)
//	}
//}
//func TestHello(t *testing.T) {
//	t.Run("saying hello to people", func(t *testing.T) {
//		got := Hello("Bima")
//		want := "Hello, Bima"
//
//		if got != want {
//			t.Errorf("got %q want %q", got, want)
//		}
//
//		t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
//			got := Hello("")
//			want := "Hello, World"
//
//			if got != want {
//				t.Errorf("got %q want %q", got, want)
//			}
//		})
//	})
//}

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got string, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bima", "")
		want := "Hello, Bima"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
