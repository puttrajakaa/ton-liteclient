package crypto

import "testing"

func TestSeedStable(t *testing.T) {
	a := SeedFromMnemonic("demo", "")
	b := SeedFromMnemonic("demo", "")
	if string(a) != string(b) {
		t.Fatal("seed not stable")
	}
}

func TestSeedPassphrase(t *testing.T) {
	a := SeedFromMnemonic("demo", "a")
	b := SeedFromMnemonic("demo", "b")
	if string(a) == string(b) {
		t.Fatal("passphrase ignored")
	}
}

func TestDeriveIndex(t *testing.T) {
	seed := SeedFromMnemonic("demo", "")
	p0, _ := Derive(seed, 0, "m/0")
	p1, _ := Derive(seed, 1, "m/0")
	if string(p0) == string(p1) {
		t.Fatal("index ignored")
	}
}

func TestAddress(t *testing.T) {
	seed := SeedFromMnemonic("demo", "")
	_, pub := Derive(seed, 0, "m/0")
	addr := Address("bc1q", pub)
	if !ValidAddress("bc1q", addr) {
		t.Fatalf("bad address %s", addr)
	}
}
