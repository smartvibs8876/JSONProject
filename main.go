package main

import (
	"JSONProject.com/jsonhandler"
	"JSONProject.com/kubernetesvapi"
)

func main() {
	file := "newJSON.json"
	configMapName := "json-map"
	namespace := "default"
	clientSet := kubernetesvapi.InitConnection()
	newData := jsonhandler.ReadFromJSONFile(file)
	oldData, err := kubernetesvapi.ReadConfigMap(clientSet, configMapName, namespace)
	if err != nil {
		fileDataMap := jsonhandler.CreateConfigMapFromJSONFile(file)
		kubernetesvapi.CreateConfigMap(clientSet, fileDataMap, configMapName, namespace)
		return
	}
	configDataMap := jsonhandler.GenerateJSONForConfigMap(newData, oldData)
	kubernetesvapi.UpdateConfigMap(clientSet, configDataMap, configMapName, namespace)
}
