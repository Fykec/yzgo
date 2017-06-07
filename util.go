package yzgo

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

//PrintResult Print api result
func PrintResult(retBytes []byte, err error) {
	if err != nil {
		panic(err)
	}
	response, err := getRawResponse(retBytes)
	// string to json help to have a better chinese printting https://segmentfault.com/q/1010000006778053
	jsonBytes, _ := json.Marshal(response)

	var out bytes.Buffer
	err = json.Indent(&out, jsonBytes, "", "  ")
	if err != nil {
		fmt.Println(string(retBytes))
	} else {
		fmt.Println(out.String())
	}
}

func PrintObject(ret interface{}) {
	//fmt.Println(err)
	spew.Config.Indent = "  "
	spew.Config.SortKeys = true
	spew.Config.DisableCapacities = true
	spew.Dump(ret)
}
