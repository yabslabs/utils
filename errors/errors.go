package errors

import (
	"strings"

	"github.com/yabslabs/utils/logging"
)

func Contains(err error, needle string) bool {
	return err != nil && strings.Contains(err.Error(), needle)
}

func PanicErrID(id string, err error, msg ...string) {
	if err != nil {
		logging.Log(id).OnError(err).Panic(msg)
	}
}
