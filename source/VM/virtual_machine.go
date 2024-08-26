package VM

import (
	"strings"
	"swing/source/VM/Method"
	"swing/source/VM/Process"
	"swing/source/VM/Stack_Frame"
)

var F = Stack_Frame.New_Stack()

func Main_run(file_name string) {
	F.ADD_Stack(file_name)
	Line := Process.ReadFile(file_name)
	for _, item := range Line {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		Method.Ck_print(F, item)
		Method.Ck_var(F, item)
	}
}
