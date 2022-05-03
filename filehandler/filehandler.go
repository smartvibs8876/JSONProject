package filehandler

import (
	"fmt"
	"os"
)

func ReadFromJSONFile(fileName string) *os.File {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return jsonFile
}
