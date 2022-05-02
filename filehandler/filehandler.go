package filehandler

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadJSON(fileName string) *os.File {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return jsonFile
}

func WriteJSON(fileName string, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), 0644)
	if err != nil {
		fmt.Println(err.Error())
	}
}
