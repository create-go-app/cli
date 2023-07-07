package containers

// Redis ...
type Redis struct {
	Name        string            `koanf:"name"`
	Version     string            `koanf:"version"`
	Credentials *redisCredentials `koanf:"credentials"`
	Options     *redisOptions     `koanf:"options"`
}

// redisCredentials ...
type redisCredentials struct {
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Database string `koanf:"database"`
}

// redisOptions ...
type redisOptions struct {
	Port int `koanf:"port"`
}
