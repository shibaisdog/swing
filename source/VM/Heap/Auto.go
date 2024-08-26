package Heap

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Auto(str string) interface{} {
	if IsEmpty(str) {
		return nil
	} else {
		if regexp.MustCompile(`^-?\d+$`).MatchString(str) {
			i, err := strconv.Atoi(str)
			if err != nil {
				fmt.Println("Integer conversion error:", err)
				os.Exit(1)
			} else {
				return i
			}
		} else if regexp.MustCompile(`^-?\d+(\.\d+)?$`).MatchString(str) {
			f, err := strconv.ParseFloat(str, 64)
			if err != nil {
				fmt.Println("Real number conversion error:", err)
			} else {
				return f
			}
		} else if str == "true" {
			return true
		} else if str == "false" {
			return false
		} else {
			return str
		}
	}
	return str
}
