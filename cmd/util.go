package cmd

import (
	"fmt"
)

func CommandBuilder(entity, op string) string {
	return fmt.Sprintf("%s-%s", entity, op)
}
