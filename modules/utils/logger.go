package utils

import (
	"fmt"
	"runtime"
)

// Output format: [function caller]: fileName - line number - your input
func Log(arg interface{}) {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok {
		fmt.Printf("[function %s]: %s:%d \nlog: %s --\n", details.Name(), file, no, arg)
	}
}
