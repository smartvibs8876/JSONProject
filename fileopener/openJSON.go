package fileopener

import (
	"fmt"
	"os"
)

func OpenJSON(filename string) *os.File {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Successfully Opened file " + filename)
	}
	return jsonFile
}
