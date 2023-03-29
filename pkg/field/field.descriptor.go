package field

import (
	"fmt"
	"strings"
)

func Describe(src string) (desc Field) {
	desc = Field{}

	_ = parser(src)

	return
}

func parser(str string) (res []string) {
	res = split(str, " ")

	for _, item := range res {
		fmt.Println(item)
	}

	return
}

func split(str string, sep string) (res []string) {
	res = strings.Split(str, sep)

	return res
}

func index(str string, substr string) (idx int) {
	idx = strings.Index(str, substr)

	return
}
