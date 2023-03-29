package parser

import "strings"

func Parse(data string) Parser {
	defs := breaker(data)

	return diff(defs)
}

func diff(defs []string) Parser {
	det := Parser{}
	for _, def := range defs {
		switch {
		case strings.HasPrefix(def, "enum"):
			det.Enum = append(det.Enum, def)
		case strings.HasPrefix(def, "typedef"):
			det.Definitions = append(det.Definitions, def)
		case strings.HasPrefix(def, "struct"):
			det.Structures = append(det.Structures, def)
		}
	}

	return det
}

func breaker(data string) []string {
	var defs []string
	temp := ""
	count := 0
	var prev rune = 0

	for idx, ch := range data {
		temp += string(ch)

		if ch == '\n' && prev == '\r' {
			count += 1
		} else if ch != '\n' && ch != '\r' {
			count = 0
		}

		if count == 2 || idx == (len(data)-1) {
			defs = append(defs, temp)
			temp = ""
			count = 0
		}

		prev = ch
	}

	return defs
}
