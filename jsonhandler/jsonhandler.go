package jsonhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"JSONProject.com/filehandler"
)

func CreateConfigMapFromJSONFile(fileName string) map[string]string {
	var configMap map[string]string = make(map[string]string)
	configMap["configuration"] = ReadFromJSONFile(fileName)
	return configMap
}
func ReadFromJSONFile(fileName string) string {
	JSONFile := filehandler.ReadFromJSONFile(fileName)
	ByteValue, _ := ioutil.ReadAll(JSONFile)
	return string(ByteValue)
}

func GenerateJSONForConfigMap(JSONFile string, ConfigMap string) map[string]string {
	ConfigMapByteValue := []byte(ConfigMap)
	JSONFileByteValue := []byte(JSONFile)
	var oldJSON map[string]interface{}
	var newJSON map[string]interface{}
	json.Unmarshal(ConfigMapByteValue, &oldJSON)
	json.Unmarshal(JSONFileByteValue, &newJSON)
	checkFieldValues(oldJSON, newJSON)
	oldJSONStr, _ := json.Marshal(oldJSON)
	mapStr := make(map[string]string)
	mapStr["configuration"] = string(oldJSONStr)
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
