package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getOldJSON(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(req.Body)
	var reqBodyMap map[string]interface{}
	json.Unmarshal([]byte(reqBody), &reqBodyMap)

}
