package tools

// Docker ...
type Docker struct {
	Network *Network `koanf:"network"`
}

// Network ...
type Network struct {
	Name string `koanf:"name"`
}
