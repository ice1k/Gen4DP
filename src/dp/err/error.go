package err

import (
	"fmt"
	"container/list"
	"strings"
)

var Messages = list.New()

func Raise(str string) {
	fmt.Errorf("%s\n", str)
	Messages.PushBack(str)
}

func GetMessages() string {
	ret := make([]string, Messages.Len())
	i := 0
	for e := Messages.Front(); e != nil; e = e.Next() {
		ret[i] = fmt.Sprint(e.Value)
		i++
	}
	return strings.Join(ret, "\n")
}
