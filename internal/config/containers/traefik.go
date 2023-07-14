package containers

// Traefik ...
type Traefik struct {
	Name        string              `koanf:"name"`
	Version     string              `koanf:"version"`
	Credentials *traefikCredentials `koanf:"credentials"`
	Options     *traefikOptions     `koanf:"options"`
}

// traefikCredentials ...
type traefikCredentials struct {
	User     string `koanf:"user"`
	Password string `koanf:"password"`
}

// traefikOptions ...
type traefikOptions struct {
	Wildcard     bool                `koanf:"wildcard"`
	Staging      bool                `koanf:"staging"`
	Port         int                 `koanf:"port"`
	DashboardURL string              `koanf:"dashboard_url"`
	Log          *traefikOptionsLog  `koanf:"log"`
	ACME         *traefikOptionsACME `koanf:"acme"`
}

// traefikOptionsLog ...
type traefikOptionsLog struct {
	Level  string `koanf:"level"`
	Format string `koanf:"format"`
}

// traefikOptionsACME ...
type traefikOptionsACME struct {
	Provider string `koanf:"provider"`
	Token    string `koanf:"token"`
}
