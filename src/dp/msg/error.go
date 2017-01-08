package msg

import (
	"dp/util/sb"
)

var warnings = sb.NewStringBuffer()

var fatal = sb.NewStringBuffer()

func Raise(str string) {
	fatal.Append(str)
}

func RaiseFormat(format string, args ...interface{}) {
	fatal.AppendFormat(format, args...)
}

func Warn(str string) {
	warnings.Append(str)
}

func WarnFormat(format string, args ...interface{}) {
	warnings.AppendFormat(format, args...)
}

func GetErrors() string {
	return fatal.ToString()
}

func GetWarnings() string {
	return warnings.ToString()
}

func HasError() bool {
	return !fatal.IsEmpty()
}

func Clear() {
	fatal.Clear()
	warnings.Clear()
}
