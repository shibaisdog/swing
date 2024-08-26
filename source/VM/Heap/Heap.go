package Heap

import (
	"fmt"
	"os"
	"strconv"
)

func (t *Type) SetType(typeName string, value bool) {
	switch typeName {
	case "INT_16":
		t.INT_16 = value
	case "INT_32":
		t.INT_32 = value
	case "INT_64":
		t.INT_64 = value
	case "FLOAT_32":
		t.FLOAT_32 = value
	case "FLOAT_64":
		t.FLOAT_64 = value
	case "STRING":
		t.STRING = value
	case "MAP":
		t.MAP = value
	case "LIST":
		t.LIST = value
	case "BOOL":
		t.BOOL = value
	case "FUNC":
		t.FUNC = value
	case "NULL":
		t.NULL = value
	default:
		fmt.Println("Unknown type:", typeName)
		os.Exit(1)
	}
}

func (t Type) IsType(typeName string) bool {
	switch typeName {
	case "INT_16":
		return t.INT_16
	case "INT_32":
		return t.INT_32
	case "INT_64":
		return t.INT_64
	case "FLOAT_32":
		return t.FLOAT_32
	case "FLOAT_64":
		return t.FLOAT_64
	case "STRING":
		return t.STRING
	case "MAP":
		return t.MAP
	case "LIST":
		return t.LIST
	case "BOOL":
		return t.BOOL
	case "FUNC":
		return t.FUNC
	case "NULL":
		return t.NULL
	default:
		fmt.Println("Unknown type:", typeName)
		os.Exit(1)
		return false
	}
}

func (t Type) GET_Name() string {
	if t.INT_16 {
		return "INT_16"
	} else if t.INT_32 {
		return "INT_32"
	} else if t.INT_64 {
		return "INT_64"
	} else if t.FLOAT_32 {
		return "FLOAT_32"
	} else if t.FLOAT_64 {
		return "FLOAT_64"
	} else if t.STRING {
		return "STRING"
	} else if t.MAP {
		return "MAP"
	} else if t.LIST {
		return "LIST"
	} else if t.BOOL {
		return "BOOL"
	} else if t.FUNC {
		return "FUNC"
	} else if t.NULL {
		return "NULL"
	} else {
		fmt.Println("Unknown type:", t)
		os.Exit(1)
		return "Unknown"
	}
}

func IsEmpty(value interface{}) bool {
	if value == nil {
		return true
	}
	switch v := value.(type) {
	case string:
		return v == ""
	case []interface{}:
		return len(v) == 0
	case map[interface{}]interface{}:
		return len(v) == 0
	case int, float64:
		return v == ""
	case bool:
		return false
	default:
		return v == ""
	}
}

func (v *Variable) INIT_Dynamic(Set_Name string, Set_Value interface{}, Set_Const bool) {
	v.Name = Set_Name
	Var_Type := Type{}
	if IsEmpty(Set_Value) {
		Var_Type.SetType("NULL", true)
	} else {
		switch value := Set_Value.(type) {
		case string:
			Var_Type.SetType("STRING", true)
			v.Value = value
		case int:
			Var_Type.SetType("INT_32", true)
			v.Value = strconv.Itoa(value)
		case float64:
			Var_Type.SetType("FLOAT_64", true)
			v.Value = strconv.FormatFloat(value, 'f', -1, 64)
		case bool:
			Var_Type.SetType("BOOL", true)
			v.Value = strconv.FormatBool(value)
		case nil:
			Var_Type.SetType("NULL", true)
		case []interface{}:
			Var_Type.SetType("LIST", true)
			v.Value = value
		case map[interface{}]interface{}:
			Var_Type.SetType("MAP", true)
			v.Value = value
		default:
			fmt.Println("Unknown type:", value)
			os.Exit(1)
		}
	}
	v.Type = Var_Type
	v.Const = Set_Const
}

func (v Variable) Info() (Type, interface{}, string) {
	return v.Type, v.Value, v.Name
}

func (v *Variable) Change_Value(Set_Name string, Set_Value interface{}) {
	if v.Const {
		fmt.Println("The value of this variable cannot be changed")
		os.Exit(1)
	} else {
		v.INIT_Dynamic(Set_Name, Set_Value, v.Const)
	}
}
