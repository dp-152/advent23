package path

import (
	"path/filepath"
	"runtime"
)

func OwnPath() string {
	_, callerFile, _, _ := runtime.Caller(1)
	return filepath.Dir(callerFile)
}
