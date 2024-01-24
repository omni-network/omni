package app

const (
	port = 8080
)

type ExplorerAPIConfig struct {
	Port int
}

func DefaultExplorerAPIConfig() ExplorerAPIConfig {
	return ExplorerAPIConfig{
		Port: port,
	}
}
