package helpers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/hcl"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/knadh/koanf/v2"
)

// ParseFileWithEnvToStruct parses the given file from path to struct *T using
// "knadh/koanf" package with an (optional) environment variables for a secret
// data.
func ParseFileWithEnvToStruct[T any](path, envPrefix string, model *T) (*T, error) {
	// Check, if path is not empty.
	if path == "" {
		return nil, errors.New("error: given path of the structured file is empty")
	}

	// Check, if environment variables prefix was given.
	if envPrefix == "" {
		return nil, errors.New("error: given environment variables prefix is empty")
	}

	// Create a new koanf instance and parse the given path.
	k, err := newKoanfByPath(path)
	if err != nil {
		return nil, err
	}

	// Load environment variables.
	if err = k.Load(env.Provider(envPrefix, ".", func(s string) string {
		// Return cleared value of the environment variables.
		return strings.ReplaceAll(
			strings.ToLower(strings.TrimPrefix(s, fmt.Sprintf("%s_", envPrefix))),
			"_", ".",
		)
	}), nil); err != nil {
		return nil, fmt.Errorf("error parsing environment variables, %w", err)
	}

	// Merge environment variables into the structured file data.
	if err = k.Merge(k); err != nil {
		return nil, fmt.Errorf("error merging environment variables into the structured file data, %w", err)
	}

	// Unmarshal structured data to the given struct.
	if err = k.Unmarshal("", &model); err != nil {
		return nil, fmt.Errorf("error unmarshalling data from structured file to struct, %w", err)
	}

	return model, nil
}

// newKoanfByPath helps to parse the given path for ParseFileToStruct and
// ParseFileWithEnvToStruct functions.
func newKoanfByPath(path string) (*koanf.Koanf, error) {
	// Create a new koanf instance.
	k := koanf.New(".")

	// Create a new variable with structured file extension.
	parserFormat := filepath.Ext(path)

	// Check the format of the structured file.
	switch parserFormat {
	case ".json", ".yaml", ".yml", ".toml", ".tf":
		// Create a new variable for the koanf parser.
		var parser koanf.Parser

		// Check the format of the structured file for get right koanf parser.
		switch parserFormat {
		case ".json":
			parser = json.Parser() // JSON format parser
		case ".yaml", ".yml":
			parser = yaml.Parser() // YAML format parser
		case ".toml":
			parser = toml.Parser() // TOML format parser
		case ".tf":
			parser = hcl.Parser(true) // HCL (Terraform) format parser
		}

		// Parse path of the structured file as URL.
		u, _ := url.Parse(path)

		// Check the schema of the given URL.
		switch u.Scheme {
		case "", "file":
			// Get the structured file from system path.
			fileInfo, err := os.Stat(path)

			// Check, if file exists.
			if err == nil || !os.IsNotExist(err) {
				// Check, if file is not dir.
				if fileInfo.IsDir() {
					return nil, fmt.Errorf("error: path of the structured file (%s) is dir", path)
				}

				// Load structured file from path (with parser of the file format).
				if err = k.Load(file.Provider(path), parser); err != nil {
					return nil, fmt.Errorf(
						"error: not valid structure of the %s file from the given path (%s)",
						strings.ToUpper(strings.TrimPrefix(parserFormat, ".")), path,
					)
				}
			} else {
				return nil, fmt.Errorf("error: structured file is not found in the given path (%s)", path)
			}
		case "http", "https":
			// Get the given file from URL.
			resp, err := http.Get(path)
			if err != nil {
				return nil, fmt.Errorf("error: structured file is not found in the given URL (%s)", path)
			}
			defer resp.Body.Close()

			// Read the structured file from URL.
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, errors.New("error: raw body from the URL is not valid")
			}

			// Load structured file from URL (with parser of the file format).
			if err = k.Load(rawbytes.Provider(body), parser); err != nil {
				return nil, fmt.Errorf(
					"error: not valid structure of the %s file from the given URL (%s)",
					strings.ToUpper(strings.TrimPrefix(parserFormat, ".")), path,
				)
			}
		default:
			// If the path's schema is unknown, default action is error.
			return nil, errors.New("error: unknown path of structured file, use system path or http(s) URL")
		}
	default:
		// If the format of the structured file is unknown, default action is error.
		return nil, errors.New("error: unknown format of structured file, see: https://github.com/knadh/koanf")
	}

	return k, nil
}
