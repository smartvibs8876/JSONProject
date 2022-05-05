/*
Main Package to initialise the file path, cofigmap and namespace in which the config map will be created
Makes calls to jsonhandler package to read data from config map , json file and do a comparison between them
Makes calls to kubernetesvapi package to read , create and update config maps
*/
package main

import (
	"JSONProject.com/jsonhandler"
	"JSONProject.com/kubernetesvapi"
)

func main() {
	//Stores relative filepath to the json file
	file := "newJSON.json"
	//Name for the config map
	configMapName := "json-map"
	//Namespace in which the config map will be stored
	namespace := "default"
	//Connect to minikube cluster
	clientSet := kubernetesvapi.InitConnection()
	//Fetch json data in string format
	newData := jsonhandler.ReadFromJSONFile(file)
	//Fetch data from config map in string format if it exists
	oldData, err := kubernetesvapi.ReadConfigMap(clientSet, configMapName, namespace)
	//Create a config map if it doesn't exist
	if err != nil {
		fileDataMap := jsonhandler.CreateConfigMapFromJSONFile(file)
		kubernetesvapi.CreateConfigMap(clientSet, fileDataMap, configMapName, namespace)
		return
	}
	//Compare config map and file data and do necessary changes in config map
	configDataMap := jsonhandler.GenerateJSONForConfigMap(newData, oldData)
	//Update the config map
	kubernetesvapi.UpdateConfigMap(clientSet, configDataMap, configMapName, namespace)
}
