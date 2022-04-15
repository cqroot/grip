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

func ValidateLocalTarget(targetName string) error {
	var fi os.FileInfo

	// check target
	fi, err := os.Stat(targetName)
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetValidateError, err.Error())
	}
	if !fi.IsDir() {
		return fmt.Errorf("%w: the provided target is not a directory", LocalTargetValidateError)
	}

	// check config
	configName := filepath.Join(targetName, "grip.yaml")
	fi, err = os.Stat(configName)
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetValidateError, err.Error())
	}
	if fi.IsDir() {
		return fmt.Errorf("%w: the provided target config is not a file", LocalTargetValidateError)
	}

	return nil
}

func ExecuteLocalTarget(targetName string) error {
	config, err := ioutil.ReadFile(filepath.Join(targetName, "grip.yaml"))
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetExecuteError, err.Error())
	}

	targets, err := ReadConfig(config)
	if err != nil {
		return fmt.Errorf("%w: %s", LocalTargetExecuteError, err.Error())
	}

	for _, target := range targets {
		fmt.Printf("%+v", target)
	}
	return nil
}
