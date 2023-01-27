package tg_api

import (
	"TESTBOT/api"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetNewMsg() map[string]interface{} {
	resp, err := http.Get(api.UrlToken + api.GetUpdates)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	var r interface{}
	err = json.Unmarshal(body, &r)
	//r, _ := utf8.DecodeRuneInString(string(body))
	fmt.Println(r)
	//fmt.Println(body)
	//ioutil.WriteFile("dump", body, 0600)
	myMap := r.(map[string]interface{})

	return myMap
}
