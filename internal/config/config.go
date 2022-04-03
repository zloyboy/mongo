package config

type Config struct {
	BindAddr string
	MongoURL string
	MongoDB  string
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		MongoURL: "mongodb://localhost:2717",
		MongoDB:  "events",
	}
}
