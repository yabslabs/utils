package errors

import (
	"strings"
)

func Contains(err error, needle string) bool {
	return err != nil && strings.Contains(err.Error(), needle)
}
