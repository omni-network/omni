package app

const (
	port  = 8080
	dbURL = "postgres://omni:password@db:5432/omni_db"
)

type ExplorerGralQLConfig struct {
	Port  int
	DBUrl string
}

func DefaultExplorerAPIConfig() ExplorerGralQLConfig {
	return ExplorerGralQLConfig{
		Port:  port,
		DBUrl: dbURL,
	}
}
