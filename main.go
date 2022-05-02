package main

import (
	"JSONProject.com/jsonhandler"
	"JSONProject.com/kubernetesvapi"
)

func main() {
	clientSet := kubernetesvapi.InitConnection()
	oldFile := "oldJSON.json"
	newFile := "newJSON.json"
	data := jsonhandler.ReadJSONFromFiles(oldFile, newFile)
	configMapName := "json-map"
	kubernetesvapi.CreateConfigMap(clientSet, data, configMapName)
	data = kubernetesvapi.ReadConfigMap(clientSet, configMapName)
	data = jsonhandler.GenerateJSON(data)
	jsonhandler.WriteJSONToFile(oldFile, data)
	kubernetesvapi.CreateConfigMap(clientSet, data, configMapName)
}
