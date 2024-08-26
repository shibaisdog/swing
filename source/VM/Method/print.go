package Method

import (
	"strings"
	"swing/source/VM/Stack_Frame"
)

func Ck_print(FS *Stack_Frame.Stack, Line string) {
	if strings.HasPrefix(Line, "print") {
		trimmed := strings.TrimPrefix(Line, "print")
		openParenIndex := strings.Index(trimmed, "(")
		closeParenIndex := strings.Index(trimmed, ")")
		if openParenIndex == -1 || closeParenIndex == -1 || openParenIndex > closeParenIndex {
			return
		}
		content := trimmed[openParenIndex+1 : closeParenIndex]
		//fmt.Println(strings.TrimSpace(content))
		FS.GET_FS(FS.RUN_FILE).GET_STACK().Info_Print(strings.TrimSpace(content))
	}
}
