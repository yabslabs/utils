package tracing

import "fmt"

const (
	spanNameFormat = "%v/%v.%v"
)

func CreateSpanName(gitProjectPath, pkg, method string) string {
	return fmt.Sprintf(spanNameFormat, gitProjectPath, pkg, method)
}
