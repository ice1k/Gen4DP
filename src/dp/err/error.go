package err

import (
	"container/list"
	"fmt"
	"strings"
)

var warnings = list.New()

var fatal = list.New()

func getAll(ls list.List) string {
	ret := make([]string, ls.Len())
	index := 0
	for e := ls.Front(); e != nil; e = e.Next() {
		ret[index] = fmt.Sprint(e.Value)
		index++
	}
	return strings.Join(ret, "\n")
}

func Raise(str string) {
	fmt.Errorf("%s\n", str)
	fatal.PushBack(str)
}

func RaiseFormat(format string, args ...interface{}) {
	res := fmt.Sprintf(format, args...)
	fmt.Errorf(res)
	fatal.PushBack(res)
}

func Warn(str string) {
	fmt.Errorf("%s\n", str)
	warnings.PushBack(str)
}

func WarnFormat(format string, args ...interface{}) {
	res := fmt.Sprintf(format, args...)
	fmt.Errorf(res)
	warnings.PushBack(res)
}

func GetErrors() string {
	return getAll(*fatal)
}

func GetWarnings() string {
	return getAll(*warnings)
}

func HasError() bool {
	return fatal.Len() > 0
}

func Clear() {
	fatal.Init()
	warnings.Init()
}
