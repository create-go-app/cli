package tools

// SSL ...
type SSL struct {
	Wildcard bool   `koanf:"wildcard"`
	Staging  bool   `koanf:"staging"`
	Email    string `koanf:"email"`
}
