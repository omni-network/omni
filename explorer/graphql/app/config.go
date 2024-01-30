package app

const (
	port  = 8080
	dbURL = "postgres://omni:password@db:5432/omni_db"
)

type ExplorerGraphQLConfig struct {
	Port  int
	DBUrl string
}

func DefaultExplorerAPIConfig() ExplorerGraphQLConfig {
	return ExplorerGraphQLConfig{
		Port:  port,
		DBUrl: dbURL,
	}
}
