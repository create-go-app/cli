package containers

// Postgres ...
type Postgres struct {
	Name        string               `koanf:"name"`
	Version     string               `koanf:"version"`
	Credentials *postgresCredentials `koanf:"credentials"`
	Options     *postgresOptions     `koanf:"options"`
}

// postgresCredentials ...
type postgresCredentials struct {
	User     string `koanf:"user"`
	Password string `koanf:"password"`
	Database string `koanf:"database"`
}

// postgresOptions ...
type postgresOptions struct {
	Port    int    `koanf:"port"`
	SSLMode string `koanf:"ssl_mode"`
}
