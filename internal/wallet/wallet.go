package wallet

import (
	"crypto/sha256"
	"encoding/hex"

	"github.com/vaultlabs/ton-liteclient/internal/config"
	cry "github.com/vaultlabs/ton-liteclient/internal/crypto"
)

type Account struct {
	Index   int
	Label   string
	Address string
	Balance int
}

type Vault struct {
	ID       string
	Name     string
	Accounts []Account
}

func (v Vault) Total() int {
	n := 0
	for _, a := range v.Accounts {
		n += a.Balance
	}
	return n
}

type Service struct {
	cfg config.Config
}

func New(cfg config.Config) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Create(name, pass string) Vault {
	seed := cry.SeedFromMnemonic(name+" "+pass, s.cfg.Coin)
	sum := sha256.Sum256([]byte(name + s.cfg.Coin))
	_, pub := cry.Derive(seed, 0, s.cfg.Path)
	return Vault{
		ID:   hex.EncodeToString(sum[:8]),
		Name: name,
		Accounts: []Account{{
			Index:   0,
			Label:   "Primary",
			Address: cry.Address(s.cfg.Prefix, pub),
		}},
	}
}

func (s *Service) Add(v Vault, label string) Vault {
	seed := cry.SeedFromMnemonic(v.Name+" demo", s.cfg.Coin)
	_, pub := cry.Derive(seed, len(v.Accounts), s.cfg.Path)
	v.Accounts = append(v.Accounts, Account{
		Index:   len(v.Accounts),
		Label:   label,
		Address: cry.Address(s.cfg.Prefix, pub),
	})
	return v
}
