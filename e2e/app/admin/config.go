package admin

type Config struct {
	// Chain is the Name of chain to run on. Leave empty to run on all applicable chains.
	Chain string
}

func DefaultConfig() Config {
	return Config{Chain: ""}
}
