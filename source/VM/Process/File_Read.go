package Process

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("file read error : ", err)
	}
	content := string(data)
	content = strings.ReplaceAll(content, "\n", "")
	list := strings.Split(content, ";")
	return list
}
