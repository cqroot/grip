package target

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	LocalTargetValidateError error = errors.New("local target validate error")
	LocalTargetExecuteError  error = errors.New("local target execute error")
)

type LocalTarget struct {
	name string
}

func (t LocalTarget) Validate() error {
	var fi os.FileInfo

	// check target
	fi, err := os.Stat(t.name)
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetValidateError, err.Error())
	}
	if !fi.IsDir() {
		return fmt.Errorf("%w: the provided target is not a directory", LocalTargetValidateError)
	}

	// check config
	configName := filepath.Join(t.name, "grip.yaml")
	fi, err = os.Stat(configName)
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetValidateError, err.Error())
	}
	if fi.IsDir() {
		return fmt.Errorf("%w: the provided target config is not a file", LocalTargetValidateError)
	}

	return nil
}

func (t LocalTarget) Execute() error {
	config, err := ioutil.ReadFile(filepath.Join(t.name, "grip.yaml"))
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetExecuteError, err.Error())
	}

	stages, err := ReadConfig(config)
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetExecuteError, err.Error())
	}

	for _, stage := range stages {
		fmt.Printf("%+v", stage)
	}
	return nil
}
