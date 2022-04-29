package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"JSONProject.com/fileopener"
)

func GetNewJSON() {
	oldJSONFile := fileopener.OpenJSON("oldJSON.json")
	newJSONFile := fileopener.OpenJSON("newJSON.json")
	oldByteValue, _ := ioutil.ReadAll(oldJSONFile)
	newByteValue, _ := ioutil.ReadAll(newJSONFile)
	var oldJSON map[string]interface{}
	var newJSON map[string]interface{}
	json.Unmarshal([]byte(oldByteValue), &oldJSON)
	json.Unmarshal([]byte(newByteValue), &newJSON)
	checkFieldValues(oldJSON, newJSON)
	//fmt.Println(oldJSON)
	jsonStr, _ := json.Marshal(oldJSON)
	ioutil.WriteFile("oldJSON.json", jsonStr, 0644)
	oldJSONFile.Close()
	newJSONFile.Close()
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
