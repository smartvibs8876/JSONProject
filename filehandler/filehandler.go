/*
	Filehandler package to read the contents of the json file
	Used by jsonhandler package
*/
package filehandler

import (
	"fmt"
	"os"
)

//Function opens the file and returns the file object
func ReadFromJSONFile(fileName string) *os.File {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err.Error())
	}
	return jsonFile
}
