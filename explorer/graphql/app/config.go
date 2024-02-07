package app

const (
	defaultListenAddr = ":8080"
	defaultDBURL      = "postgres://omni:password@db:5432/omni_db"
)

type Config struct {
	ListenAddress string
	DBUrl         string
}

func DefaultConfig() Config {
	return Config{
		ListenAddress: defaultListenAddr,
		DBUrl:         defaultDBURL,
	}
}
