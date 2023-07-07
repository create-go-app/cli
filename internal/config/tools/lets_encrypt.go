package tools

// LetsEncrypt ...
type LetsEncrypt struct {
	Wildcard bool   `koanf:"wildcard"`
	Domain   string `koanf:"domain"`
	ACME     *ACME  `koanf:"acme"`
}

// ACME ...
type ACME struct {
	Staging bool   `koanf:"staging"`
	Email   string `koanf:"email"`
}
