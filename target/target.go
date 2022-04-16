package target

import (
	"fmt"
	"strings"
)

type Target interface {
	Validate() error
	Execute() error
}

func NewTarget(targetName string) Target {
	if strings.HasPrefix(targetName, "github.com") {
		fmt.Printf("Read t from github: %s\n", targetName)
		return nil
	} else {
		return LocalTarget{
			name: targetName,
		}
	}
}
