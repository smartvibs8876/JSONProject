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
	for key, _ := range newJSON {
		if fmt.Sprintf("%T", newJSON[key]) == "map[string]interface {}" && fmt.Sprintf("%T", oldJSON[key]) == "map[string]interface {}" {
			checkFieldValues(oldJSON[key].(map[string]interface{}), newJSON[key].(map[string]interface{}))
		}
		if fmt.Sprintf("%T", newJSON[key]) == "[]interface {}" && fmt.Sprintf("%T", oldJSON[key]) == "[]interface {}" {
			objectExists := false
			oldJsonArray := oldJSON[key].([]interface{})
			newJsonArray := newJSON[key].([]interface{})
			for i := range newJsonArray {
				if fmt.Sprintf("%T", newJsonArray[i]) == "map[string]interface {}" {
					if i < len(oldJsonArray) && fmt.Sprintf("%T", oldJsonArray[i]) == "map[string]interface {}" {
						objectExists = true
						checkFieldValues(oldJsonArray[i].(map[string]interface{}), newJsonArray[i].(map[string]interface{}))
					}
				}
			}
			if objectExists == false {
				if len(newJsonArray) != 0 {
					oldJSON[key] = newJSON[key]
				}
			}
		}
		if oldJSON[key] == nil {
			oldJSON[key] = newJSON[key]
		}
	}
	oldJSON = sortByJSONKeys(oldJSON, newJSON)
	fmt.Println(oldJSON)
}

func sortByJSONKeys(oldJSON map[string]interface{}, newJSON map[string]interface{}) map[string]interface{} {
	sortedOldJSON := make(map[string]interface{})
	for key, _ := range newJSON {
		if oldJSON[key] != nil {
			sortedOldJSON[key] = oldJSON[key]
		}
	}
	return sortedOldJSON
}
