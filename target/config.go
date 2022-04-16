package target

import (
	"gopkg.in/yaml.v3"
)

type Stage struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}

func ReadConfig(config []byte) ([]Stage, error) {
	var stages []Stage
	err := yaml.Unmarshal(config, &stages)
	if err != nil {
		return nil, err
	}

	return stages, nil
}
