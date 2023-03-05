package common

import (
	"path"
	"path/filepath"
	"runtime"
)

var WORKSPACE = filepath.Dir(filepath.Dir(getCurrentAbsPathByCaller()))

func getCurrentAbsPathByCaller() string {
	if _, filename, _, ok := runtime.Caller(0); ok {
		return path.Dir(filename)
	}

	panic("")
}
