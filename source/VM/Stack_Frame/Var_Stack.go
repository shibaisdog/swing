package Stack_Frame

import (
	"encoding/json"
	"fmt"
	"swing/source/VM/Heap"
)

type Var_Stack struct {
	Heap map[string]Heap.Variable
}

func INIT_Var_Stack() *Var_Stack {
	return &Var_Stack{
		Heap: make(map[string]Heap.Variable),
	}
}

func (h *Var_Stack) Init_Var(name string, value interface{}, isConst bool) {
	var v Heap.Variable
	v.INIT_Dynamic(name, value, isConst)
	h.Heap[name] = v
}

func (h *Var_Stack) Change_Value(name string, value interface{}) {
	v := h.Heap[name]
	v.Change_Value(name, value)
}

func (h *Var_Stack) Info(name string) (Heap.Type, interface{}, string) {
	value := h.Heap[name]
	return value.Info()
}

func (m Var_Stack) Info_Print(name string) {
	typ, value, _ := m.Info(name)
	if typ.LIST {
		m.List_Print(name)
		return
	} else if typ.MAP {
		m.Map_Print(name)
		return
	}
	//fmt.Printf("type: %s value: %v name: %s\n", typ.GET_Name(), value, name)
	fmt.Printf("%v\n", value)
}

func (m Var_Stack) Map_Print(name string) {
	_, value, _ := m.Info(name)
	jsonData, err := json.Marshal(Heap.ConvertToStringMap(value.(map[interface{}]interface{})))
	if err != nil {
		fmt.Printf("Error marshaling map to JSON: %v\n", err)
		return
	}
	//fmt.Printf("type: %s value: %v name: %s\n", typ.GET_Name(), string(jsonData), name)
	fmt.Printf("%v\n", string(jsonData))
}

func (m Var_Stack) List_Print(name string) {
	_, value, _ := m.Info(name)
	jsonData, err := json.Marshal(value)
	if err != nil {
		fmt.Printf("Error marshaling list to JSON: %v\n", err)
		return
	}
	//fmt.Printf("type: %s value: %v name: %s\n", typ.GET_Name(), string(jsonData), name)
	fmt.Printf("%v\n", string(jsonData))
}
