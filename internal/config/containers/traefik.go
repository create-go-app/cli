package containers

// Traefik ...
type Traefik struct {
	Name        string              `koanf:"name"`
	Version     string              `koanf:"version"`
	Credentials *traefikCredentials `koanf:"credentials"`
	Options     *traefikOptions     `koanf:"options"`
}

// traefikCredentials ...
type traefikCredentials struct{}

// traefikOptions ...
type traefikOptions struct {
	Port int `koanf:"port"`
}
