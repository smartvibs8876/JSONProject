/*
	Jsonhandler package to compare json file and a config map
	Makes necessary changes to the config map after comparison
	Makes calls to file handler package for reading the json file
*/
package jsonhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"JSONProject.com/filehandler"
)

//Function to copy contents of json file and paste in config map. Used for first time initialisation
func CreateConfigMapFromJSONFile(fileName string) map[string]string {
	var configMap map[string]string = make(map[string]string)
	byteValue := []byte(ReadFromJSONFile(fileName))
	var JSON map[string]interface{}
	json.Unmarshal(byteValue, &JSON)
	JSONStr, _ := json.Marshal(JSON)
	configMap["configuration"] = string(JSONStr)
	return configMap
}

//Function to read json file and return in string format
func ReadFromJSONFile(fileName string) string {
	JSONFile := filehandler.ReadFromJSONFile(fileName)
	ByteValue, _ := ioutil.ReadAll(JSONFile)
	return string(ByteValue)
}

//Create a map for json and config map and makes a call to checkFieldValues for comparing and updating
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

//Recursive function for comparing and updating field values of config map
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
					objectExists = true
					newObjectFound := true
					objectFromNewArray := newJsonArray[i].(map[string]interface{})
					for j := range oldJsonArray {
						if fmt.Sprintf("%T", oldJsonArray[j]) == "map[string]interface {}" {
							objectFromOldArray := oldJsonArray[j].(map[string]interface{})
							equal := checkSameObjects(objectFromOldArray, objectFromNewArray)
							if equal {
								newObjectFound = false
								checkFieldValues(objectFromOldArray, objectFromNewArray)
								break
							}
						}
					}
					if newObjectFound == true {
						oldJSON[key] = append(oldJSON[key].([]interface{}), objectFromNewArray)
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
}

//Function to check if 2 objects are same in an array
func checkSameObjects(obj1 map[string]interface{}, obj2 map[string]interface{}) bool {
	if len(obj1) != len(obj2) {
		return false
	}
	for key := range obj1 {
		if fmt.Sprintf("%T", obj1[key]) == "[]interface {}" && fmt.Sprintf("%T", obj2[key]) == "[]interface {}" {
			continue
		} else if fmt.Sprintf("%T", obj1[key]) == "map[string]interface {}" && fmt.Sprintf("%T", obj2[key]) == "map[string]interface {}" {
			if obj1[key] != obj2[key] {
				return false
			}
		} else if fmt.Sprintf("%T", obj1[key]) == "string" && fmt.Sprintf("%T", obj2[key]) == "string" {
			if obj1[key] != obj2[key] {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
