package enumeration

import (
	"strings"
)

func Parser(str string) (info Enum) {
	info = Enum{}
	var splitCount int
	var HasFields bool = false
	var HasName bool = false

	str = strings.Trim(str, " \n\t\r")

	if strings.HasSuffix(str, "};") {
		HasFields = true
	}

	if strings.HasPrefix(str, "enum {") {
		HasName = false
		splitCount = 2
	} else {
		HasName = true
		splitCount = 3
	}

	res := split(str, " ", splitCount)

	switch {
	case HasName && !HasFields:
		info.Name = res[1]
	case !HasName && HasFields:
		info.Fields = fields(res[1])
	case HasName && HasFields:
		info.Name = res[1]
		info.Fields = fields(res[2])
	}

	info.Source = str

	return
}

func fields(str string) []constant {

	var flds []constant

	lines := breaker(str)
	for _, line := range lines {
		flds = append(flds, field(line))
	}

	return flds
}

func field(str string) constant {
	var fld constant

	res := strings.Split(str, "=")

	fld.Name = res[0]
	fld.Value = res[1]
	return fld
}

func split(str string, sep string, n int) []string {
	return strings.SplitN(str, sep, n)
}

func breaker(str string) []string {

	res := strings.Split(str, ",")
	replacer := strings.NewReplacer("\r", "", "\n", "", "{", "", "}", "", "\t", "")
	for idx, ln := range res {
		res[idx] = replacer.Replace(ln)
	}

	return res[:len(res)-1]
}
