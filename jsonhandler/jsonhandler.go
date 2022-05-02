package jsonhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"JSONProject.com/filehandler"
)

func ReadJSONFromFiles() map[string]string {
	oldJSONFile := filehandler.OpenJSON("oldJSON.json")
	newJSONFile := filehandler.OpenJSON("newJSON.json")
	oldByteValue, _ := ioutil.ReadAll(oldJSONFile)
	newByteValue, _ := ioutil.ReadAll(newJSONFile)
	mapStr := make(map[string]string)
	mapStr["oldJSON"] = string(oldByteValue)
	mapStr["newJSON"] = string(newByteValue)
	return mapStr

}

func GenerateJSON(JSONFromCM map[string]string) map[string]string {
	oldByteValue := []byte(JSONFromCM["oldJSON"])
	newByteValue := []byte(JSONFromCM["newJSON"])
	var oldJSON map[string]interface{}
	var newJSON map[string]interface{}
	json.Unmarshal([]byte(oldByteValue), &oldJSON)
	json.Unmarshal([]byte(newByteValue), &newJSON)
	checkFieldValues(oldJSON, newJSON)
	oldJSONStr, _ := json.Marshal(oldJSON)
	newJSONStr, _ := json.Marshal(newJSON)
	mapStr := make(map[string]string)
	mapStr["oldJSON"] = string(oldJSONStr)
	mapStr["newJSON"] = string(newJSONStr)
	return mapStr
}

func checkFieldValues(oldJSON map[string]interface{}, newJSON map[string]interface{}) {
	for key, _ := range oldJSON {
		if fmt.Sprintf("%T", oldJSON[key]) == "map[string]interface {}" {
			a := oldJSON[key]
			b := newJSON[key]
			checkFieldValues(a.(map[string]interface{}), b.(map[string]interface{}))
		}
		if fmt.Sprintf("%T", newJSON[key]) == "[]interface {}" {
			wentDown := false
			oldJsonArray := oldJSON[key].([]interface{})
			newJsonArray := newJSON[key].([]interface{})
			for i := range oldJsonArray {
				if fmt.Sprintf("%T", oldJsonArray[i]) == "map[string]interface {}" {
					wentDown = true
					a := oldJsonArray[i]
					b := newJsonArray[i]
					checkFieldValues(a.(map[string]interface{}), b.(map[string]interface{}))
				}
			}
			if wentDown == false {
				if len(newJsonArray) != 0 {
					oldJSON[key] = newJSON[key]
				}
			}
		}
	}
}
