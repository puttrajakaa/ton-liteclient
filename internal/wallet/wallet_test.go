package wallet

import (
	"testing"

	"github.com/vaultlabs/ton-liteclient/internal/config"
)

func TestCreate(t *testing.T) {
	s := New(config.Default())
	v := s.Create("test", "pw")
	if v.Name != "test" || len(v.Accounts) != 1 || v.Accounts[0].Address == "" {
		t.Fatalf("%+v", v)
	}
}

func TestAddUnique(t *testing.T) {
	s := New(config.Default())
	v := s.Create("test", "pw")
	v = s.Add(v, "A")
	v = s.Add(v, "B")
	seen := map[string]bool{}
	for _, a := range v.Accounts {
		if seen[a.Address] {
			t.Fatal("dup address")
		}
		seen[a.Address] = true
	}
}
