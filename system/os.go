package system

import (
	"runtime"
)

func GetOSName() string {
	return runtime.GOOS
}
