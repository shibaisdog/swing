package Heap

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func (v *ML_Variable) INIT(Set_Value interface{}) {
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
		}
	}
	v.Type = Var_Type
}

func (v ML_Variable) Info() (Type, interface{}) {
	return v.Type, v.Value
}

func (v ML_Variable) GET_Value() interface{} {
	return v.Value
}

// /////////////////////////////////////////////////////
func Creaft_List(list []interface{}) []interface{} {
	var List = []interface{}{}
	for _, item := range list {
		var Var = ML_Variable{}
		Var.INIT(item)
		List = append(List, Var)
	}
	return List
}

func Creaft_Map(m map[interface{}]interface{}, prefix string) map[interface{}]interface{} {
	result := map[interface{}]interface{}{}
	for key, value := range m {
		switch v := value.(type) {
		case map[interface{}]interface{}:
			nestedMap := Creaft_Map(v, prefix+fmt.Sprintf("%v.", key))
			for nestedKey, nestedValue := range nestedMap {
				result[nestedKey] = nestedValue
			}
		default:
			compoundKey := fmt.Sprintf("%s%v", prefix, key)
			var Var = ML_Variable{}
			Var.INIT(v)
			result[compoundKey] = Var
		}
	}
	return result
}

func ConvertToInterfaceMap(m map[string]interface{}) map[interface{}]interface{} {
	result := make(map[interface{}]interface{})
	for k, v := range m {
		switch v := v.(type) {
		case map[string]interface{}:
			result[k] = ConvertToInterfaceMap(v)
		default:
			result[k] = v
		}
	}
	return result
}

func ConvertToStringMap(m map[interface{}]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for k, v := range m {
		keyStr := fmt.Sprintf("%v", k)
		switch v := v.(type) {
		case map[interface{}]interface{}:
			result[keyStr] = ConvertToStringMap(v)
		default:
			result[keyStr] = v
		}
	}
	return result
}

func String_to_Map(str string) map[interface{}]interface{} {
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		panic(err)
	}
	return ConvertToInterfaceMap(result)
}

func String_to_List(str string) []interface{} {
	var result []interface{}
	if err := json.Unmarshal([]byte(str), &result); err != nil {
		panic(err)
	}
	return result
}
