package config

type Config struct {
	Addr string
}

func Load() Config {
	return Config{Addr: ":6380"}
}
