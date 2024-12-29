package config

import (
	"os"

	"sigs.k8s.io/yaml"
)

// New return a new config based on the provided file
func New(p string) (*AppConfig, error) {
	cnf := AppConfig{}
	f, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(f, &cnf); err != nil {
		return nil, err
	}

	return &cnf, nil
}
