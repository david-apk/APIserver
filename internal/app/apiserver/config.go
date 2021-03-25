package apiserver

//config
type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

//ew Config
func NewConfig() *Config {
	return &Config{
		BindAddr: ":3333",
		LogLevel: "debug",
	}
}
