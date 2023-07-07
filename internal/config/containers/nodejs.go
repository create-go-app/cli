package containers

// NodeJS ...
type NodeJS struct {
	Name        string             `koanf:"name"`
	Version     string             `koanf:"version"`
	Credentials *nodeJSCredentials `koanf:"credentials"`
	Options     *nodeJSOptions     `koanf:"options"`
}

// nodeJSCredentials ...
type nodeJSCredentials struct{}

// nodeJSOptions ...
type nodeJSOptions struct {
	Port int `koanf:"port"`
}
