package containers

// Nginx ...
type Nginx struct {
	Name        string            `koanf:"name"`
	Version     string            `koanf:"version"`
	Credentials *nginxCredentials `koanf:"credentials"`
	Options     *nginxOptions     `koanf:"options"`
}

// nginxCredentials ...
type nginxCredentials struct{}

// nginxOptions ...
type nginxOptions struct {
	Port int `koanf:"port"`
}
