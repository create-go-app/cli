//go:generate go run gen.go

package box

type configsBox struct {
	storage map[string][]byte
}

func newConfigsBox() *configsBox {
	return &configsBox{storage: make(map[string][]byte)}
}

// Find for a file
func (r *configsBox) Has(file string) bool {
	if _, ok := r.storage[file]; ok {
		return true
	}
	return false
}

// Get file's content
// Always use / for looking up
// For example: /init/README.md is actually resources/init/README.md
func (r *configsBox) Get(file string) []byte {
	if f, ok := r.storage[file]; ok {
		return f
	}
	return nil
}

// Add a file to box
func (r *configsBox) Add(file string, content []byte) {
	r.storage[file] = content
}

// Configs expose
var configs = newConfigsBox()

// Get a file from box
func Get(file string) []byte {
	return configs.Get(file)
}

// Add a file content to box
func Add(file string, content []byte) {
	configs.Add(file, content)
}

// Has a file in box
func Has(file string) bool {
	return configs.Has(file)
}
