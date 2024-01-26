package app

const (
	Port  = 8081
	DBURL = "postgres://omni:password@db:5432/omni_db"
)

type ExplorerIndexerConfig struct {
	Port  int
	DBUrl string
}

func DefaultExplorerAPIConfig() ExplorerIndexerConfig {
	return ExplorerIndexerConfig{
		Port:  Port,
		DBUrl: DBURL,
	}
}
