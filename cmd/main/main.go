package main

import (
	"TESTBOT/internal/tg_api"
	"fmt"
)

func main() {

	//https://api.telegram.org/bot<token>/METHOD_NAME
	umap := tg_api.GetNewMsg()
	for i, v := range umap {
		str := fmt.Sprintf("%v", v)
		fmt.Println(i, str)
	}
	//SendMessage("фывфы", "504416149")

}
