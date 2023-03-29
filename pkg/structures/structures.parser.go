package structures

import (
	"strconv"
	"strings"
)

func Parser(str string) (info Structure) {
	info = Structure{}
	var splitCount int = 2
	var HasFields bool = false

	str = strings.Trim(str, " \n\t\r")

	if strings.HasSuffix(str, "};") {
		HasFields = true
		splitCount = 3
	} else if strings.Contains(str, "}") {
		HasFields = true
		splitCount = 3

		indx := strings.Index(str, "}")
		str = str[:indx+1]
	}

	res := strings.SplitN(str, " ", splitCount)
	switch {
	case HasFields:
		info.Name = res[1]
		info.Fields = fields(res[2])
	case !HasFields:
		info.Name = res[1]
	}
	info.Source = str

	return
}

func fields(str string) []Field {

	var flds []Field

	lines := breaker(str)
	for _, line := range lines {
		flds = append(flds, field(line))
	}

	return flds
}

func field(str string) Field {
	var fld Field

	fld.Source = str

	if strings.HasSuffix(str, ")") {
		fld.IsField = false
	} else {
		fld.IsField = true
	}

	if strings.Contains(str, "[") {
		indx := strings.Index(str, "[")
		alen, _ := strconv.Atoi(strings.Trim(str[indx:], " ;[]"))

		str = str[:indx]
		fld.IsArray = true
		fld.Length = alen
	} else {
		fld.IsArray = false
	}

	if strings.Contains(str, "*") {
		str = strings.Replace(str, "*", "", 1)
		fld.IsPointer = true
	} else {
		fld.IsPointer = false
	}

	if strings.Contains(str, ":") {
		indx := strings.Index(str, ":")
		blen, _ := strconv.Atoi(strings.Trim(str[indx:], " ;:"))

		str = str[:indx]
		fld.IsBitField = true
		fld.BitLength = blen
	} else {
		fld.IsBitField = false
	}

	res := strings.Split(str, " ")
	switch {

	case fld.IsField:
		fld.Name = res[len(res)-1]
		fld.DataType = strings.Join(res[:len(res)-1], " ")
		fld.CType = ctype(fld.DataType, fld.IsArray)
		fld.GoType = gotype(fld.DataType, fld.IsArray, fld.Length)
	}

	return fld
}

func breaker(str string) []string {

	res := []string{}
	strs := strings.Split(str, ";")
	replacer := strings.NewReplacer("\r", "", "\n", "", "{", "", "}", "", "\t", "")
	for _, ln := range strs {
		s := replacer.Replace(ln)
		if len(s) != 0 {
			res = append(res, s)
		}
	}

	return res
}

func ctype(str string, isArray bool) string {
	vtyp := "__"

	if strings.Contains(str, "unsigned") {
		vtyp += "u"
	} else if (strings.Contains(str, "char") && isArray) || strings.Contains(str, "bool") {
		vtyp += "u"
	} else {
		vtyp += "s"
	}

	if strings.Contains(str, "char") || strings.Contains(str, "bool") {
		vtyp += "8"
	} else if strings.Contains(str, "short") {
		vtyp += "16"
	} else if strings.Contains(str, "long") {
		vtyp += "64"
	} else {
		vtyp += "32"
	}

	return vtyp
}

func gotype(str string, isArray bool, arrLen int) string {
	vtyp := ""

	if strings.Contains(str, "unsigned") {
		vtyp += "u"
	} else if (strings.Contains(str, "char") && isArray) || strings.Contains(str, "bool") {
		vtyp += "u"
	}

	if strings.Contains(str, "char") || strings.Contains(str, "bool") {
		vtyp += "int8"
	} else if strings.Contains(str, "short") {
		vtyp += "int16"
	} else if strings.Contains(str, "long") {
		vtyp += "int64"
	} else {
		vtyp += "int32"
	}

	if isArray {
		vtyp = "[" + strconv.Itoa(arrLen) + "]" + vtyp
	}

	return vtyp
}
