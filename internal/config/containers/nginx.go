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
	UseOnlyHTTPS     bool `koanf:"use_only_https"`
	RedirectToNonWWW bool `koanf:"redirect_to_non_www"`
	Port             int  `koanf:"port"`
}
