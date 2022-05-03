package main

import (
	"JSONProject.com/jsonhandler"
	"JSONProject.com/kubernetesvapi"
)

func main() {
	clientSet := kubernetesvapi.InitConnection()
	file := "newJSON.json"
	newData := jsonhandler.ReadFromJSONFile(file)
	configMapName := "json-map"
	namespace := "default"
	oldData, err := kubernetesvapi.ReadConfigMap(clientSet, configMapName, namespace)
	if err != nil {
		fileDataMap := jsonhandler.CreateConfigMapFromJSONFile(file)
		kubernetesvapi.CreateConfigMap(clientSet, fileDataMap, configMapName, namespace)
		return
	}
	dataMap := jsonhandler.GenerateJSONForConfigMap(newData, oldData)
	kubernetesvapi.CreateConfigMap(clientSet, dataMap, configMapName, namespace)
}
