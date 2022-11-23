package base

import (
	"fmt"
	"runtime"
)

func GetFunc() string {
	pc, _, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s:%d", runtime.FuncForPC(pc).Name(), line)
}
