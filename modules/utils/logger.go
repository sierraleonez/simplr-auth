package utils

import (
	"fmt"
	"runtime"
)

func Trace_caller() {
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	if ok {
		fmt.Println("called from", file, no, details.Name())
	}
}
