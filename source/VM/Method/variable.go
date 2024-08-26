package Method

import (
	"strings"
	"swing/source/VM/Extract"
	"swing/source/VM/Heap"
	"swing/source/VM/Stack_Frame"
)

func Ck_var(FS *Stack_Frame.Stack, Line string) {
	if strings.HasPrefix(Line, "val") {
		trimmed := strings.TrimPrefix(Line, "val ")
		equalIndex := strings.Index(trimmed, "=")
		if equalIndex == -1 {
			beforeEqual := strings.TrimSpace(trimmed)
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, nil, false)
			return
		}
		beforeEqual := strings.TrimSpace(trimmed[:equalIndex])
		afterEqual := strings.TrimSpace(trimmed[equalIndex+1:])
		if strings.HasPrefix(afterEqual, "[") && strings.HasSuffix(afterEqual, "]") {
			afterEqual = Extract.Replace_value(FS, afterEqual)
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, Heap.Creaft_List(Heap.String_to_List(afterEqual)), false)
		} else if strings.HasPrefix(afterEqual, "{") && strings.HasSuffix(afterEqual, "}") {
			afterEqual = Extract.Replace_value(FS, afterEqual)
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, Heap.Creaft_Map(Heap.String_to_Map(afterEqual), ""), false)
		} else {
			if strings.HasPrefix(afterEqual, "`") && strings.HasSuffix(afterEqual, "`") {
				afterEqual = strings.TrimSpace(afterEqual[1 : len(afterEqual)-1])
				afterEqual = Extract.Replace_value(FS, afterEqual)
				FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, afterEqual, false)
			} else {
				afterEqual = strings.ReplaceAll(afterEqual, "'", "")
				afterEqual = strings.ReplaceAll(afterEqual, `"`, "")
				FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, Heap.Auto(afterEqual), false)
			}
		}
		return
	} else if strings.HasPrefix(Line, "const") {
		trimmed := strings.TrimPrefix(Line, "const ")
		equalIndex := strings.Index(trimmed, "=")
		if equalIndex == -1 {
			beforeEqual := strings.TrimSpace(trimmed)
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, nil, true)
			return
		}
		beforeEqual := strings.TrimSpace(trimmed[:equalIndex])
		afterEqual := strings.TrimSpace(trimmed[equalIndex+1:])
		if strings.HasPrefix(afterEqual, "[") && strings.HasSuffix(afterEqual, "]") {
			afterEqual = Extract.Replace_value(FS, afterEqual)
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, Heap.Creaft_List(Heap.String_to_List(afterEqual)), false)
		} else if strings.HasPrefix(afterEqual, "{") && strings.HasSuffix(afterEqual, "}") {
			afterEqual = Extract.Replace_value(FS, afterEqual)
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, Heap.Creaft_Map(Heap.String_to_Map(afterEqual), ""), false)
		} else {
			if strings.HasPrefix(afterEqual, "`") && strings.HasSuffix(afterEqual, "`") {
				afterEqual = strings.TrimSpace(afterEqual[1 : len(afterEqual)-1])
				afterEqual = Extract.Replace_value(FS, afterEqual)
				FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, afterEqual, false)
			} else {
				afterEqual = strings.ReplaceAll(afterEqual, "'", "")
				afterEqual = strings.ReplaceAll(afterEqual, `"`, "")
				FS.GET_FS(FS.RUN_FILE).GET_STACK().Init_Var(beforeEqual, Heap.Auto(afterEqual), false)
			}
		}
		return
	} else if strings.Contains(Line, "=") {
		index := strings.Index(Line, "=")
		left := strings.TrimSpace(Line[:index])
		right := strings.TrimSpace(Line[index+1:])
		if strings.HasPrefix(right, "[") && strings.HasSuffix(right, "]") {
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Change_Value(left, Heap.Creaft_List(Heap.String_to_List(right)))
		} else if strings.HasPrefix(right, "{") && strings.HasSuffix(right, "}") {
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Change_Value(left, Heap.Creaft_Map(Heap.String_to_Map(right), ""))
		} else {
			FS.GET_FS(FS.RUN_FILE).GET_STACK().Change_Value(left, Heap.Auto(right))
		}
	}
}
