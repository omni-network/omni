package app

const (
	port = 8000
)

type ExplorerAPIConfig struct {
	Port int
}

func DefaultExplorerAPIConfig() ExplorerAPIConfig {
	return ExplorerAPIConfig{
		Port: port,
	}
}
