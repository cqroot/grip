package target

import (
	"gopkg.in/yaml.v3"
)

type Target struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func ReadConfig(config []byte) ([]Target, error) {
	var targets []Target
	err := yaml.Unmarshal(config, &targets)
	if err != nil {
		return nil, err
	}

	return targets, nil
}
