package err

import (
	"container/list"
	"fmt"
	"strings"
)

var messages = list.New()

func Raise(str string) {
	fmt.Errorf("%s\n", str)
	messages.PushBack(str)
}

func GetMessages() string {
	ret := make([]string, messages.Len())
	i := 0
	for e := messages.Front(); e != nil; e = e.Next() {
		ret[i] = fmt.Sprint(e.Value)
		i++
	}
	return strings.Join(ret, "\n")
}

func HasGotError() bool {
	return messages.Len() > 0
}
