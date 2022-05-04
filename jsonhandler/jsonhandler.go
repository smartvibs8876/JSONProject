package jsonhandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"JSONProject.com/filehandler"
)

func CreateConfigMapFromJSONFile(fileName string) map[string]string {
	var configMap map[string]string = make(map[string]string)
	byteValue := []byte(ReadFromJSONFile(fileName))
	var JSON map[string]interface{}
	json.Unmarshal(byteValue, &JSON)
	JSONStr, _ := json.Marshal(JSON)
	configMap["configuration"] = string(JSONStr)
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
					objectExists = true
					newObjectFound := true
					objectFromNewArray := newJsonArray[i].(map[string]interface{})
					for j := range oldJsonArray {
						if fmt.Sprintf("%T", oldJsonArray[j]) == "map[string]interface {}" {
							objectFromOldArray := oldJsonArray[j].(map[string]interface{})
							//fmt.Println("Old Object", objectFromOldArray)
							//fmt.Println("New Object", objectFromNewArray)
							equal := checkSameObjects(objectFromOldArray, objectFromNewArray)
							//fmt.Println(equal)
							if equal {
								newObjectFound = false
								checkFieldValues(objectFromOldArray, objectFromNewArray)
								break
							}
						}
					}
					if newObjectFound == true {
						// fmt.Println("here", objectFromNewArray)
						oldJSON[key] = append(oldJSON[key].([]interface{}), objectFromNewArray)
					}
					// if i < len(oldJsonArray) && fmt.Sprintf("%T", oldJsonArray[i]) == "map[string]interface {}" {
					// 	objectExists = true
					// 	checkFieldValues(oldJsonArray[i].(map[string]interface{}), newJsonArray[i].(map[string]interface{}))
					// }
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
