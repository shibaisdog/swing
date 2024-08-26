package Extract

import (
	"regexp"
	"strings"
	"swing/source/VM/Stack_Frame"
)

func Replace_value(FS *Stack_Frame.Stack, str string) string {
	re := regexp.MustCompile(`\$([a-zA-Z0-9]+(?:\[[a-zA-Z0-9]*\])+|[a-zA-Z0-9]+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	for _, match := range matches {
		_, v, _ := FS.GET_FS(FS.RUN_FILE).GET_STACK().Info(match[1])
		s, ok := v.(string)
		if !ok {
			s = "null"
		}
		str = strings.ReplaceAll(str, "$"+match[1], s)
	}
	return str
}
