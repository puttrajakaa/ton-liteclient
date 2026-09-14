package config

type Config struct {
	Coin       string
	Path       string
	Prefix     string
	StorageDir string
}

func Default() Config {
	return Config{
		Coin:       "TON",
		Path:       "m/44'/607'/0'",
		Prefix:     "EQ",
		StorageDir: ".wallets",
	}
}
