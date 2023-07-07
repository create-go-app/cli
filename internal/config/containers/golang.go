package containers

// Golang ...
type Golang struct {
	Name        string             `koanf:"name"`
	Version     string             `koanf:"version"`
	Credentials *golangCredentials `koanf:"credentials"`
	Options     *golangOptions     `koanf:"options"`
}

// golangCredentials ...
type golangCredentials struct{}

// golangOptions ...
type golangOptions struct {
	Port int `koanf:"port"`
}
