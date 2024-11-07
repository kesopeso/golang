package generics_test

import (
	"generics"
	"testing"
)

func TestBalanceOf(t *testing.T) {
	transactions := []generics.Transaction{
		{"Adel", "Chris", 100},
		{"Chris", "Adel", 25},
		{"Chris", "Mohamed", 25},
	}

	got := generics.BalanceOf(transactions, "Adel")
	want := -75
	if got != want {
		t.Errorf("balance not matching for Adel, got %d, want %d", got, want)
	}

	got = generics.BalanceOf(transactions, "Chris")
	want = 50
	if got != want {
		t.Errorf("balance not matching for Chris, got %d, want %d", got, want)
	}
}
