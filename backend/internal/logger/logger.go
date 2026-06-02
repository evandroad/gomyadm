package logger

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

const (
	reset  = "\033[0m"
	green  = "\033[32m"
	yellow = "\033[33m"
	red    = "\033[31m"
)

func Info(format string, args ...any) {
	print(green, "INFO", format, args...)
}

func Warn(format string, args ...any) {
	print(yellow, "WARN", format, args...)
}

func Error(format string, args ...any) {
	print(red, "ERROR", format, args...)
}

func print(color, level, format string, args ...any) {
	pc, file, line, ok := runtime.Caller(2)

	context := "unknown"

	if ok {
		fileName := filepath.Base(file)
		function := getFunctionName(pc)

		pkg := "unknown"
		method := function
		fileName = strings.TrimSuffix(fileName, ".go")

		if idx := strings.LastIndex(function, "."); idx != -1 {
			pkg = function[:idx]
			method = function[idx+1:]
		}

		context = fmt.Sprintf("%s/%s-%s:%d", pkg, fileName, method, line)
	}

	message := fmt.Sprintf(format, args...)

	log.Printf("%s[%s]%s [%s] %s", color, level, reset, context, message)
}

func getFunctionName(pc uintptr) string {
	if fn := runtime.FuncForPC(pc); fn != nil {
		full := fn.Name()

		lastSlash := strings.LastIndex(full, "/")
		if lastSlash != -1 {
			full = full[lastSlash+1:]
		}

		if idx := strings.Index(full, ".(*"); idx != -1 {
			pkg := full[:idx]

			if end := strings.Index(full[idx:], ")."); end != -1 {
				method := full[idx+end+2:]
				return pkg + "." + method
			}
		}

		return full
	}

	return "unknown"
}