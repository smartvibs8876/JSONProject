package main

import (
	"JSONProject.com/jsonhandler"
	"JSONProject.com/kubernetesvapi"
)

func main() {
	clientSet := kubernetesvapi.InitConnection()
	data := jsonhandler.ReadJSONFromFiles()
	configMapName := "json-map"
	kubernetesvapi.CreateConfigMap(clientSet, data, configMapName)
	data = kubernetesvapi.ReadConfigMap(clientSet)
	data = jsonhandler.GenerateJSON(data)
	kubernetesvapi.CreateConfigMap(clientSet, data, configMapName)
}
