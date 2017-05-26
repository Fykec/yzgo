package yzgo

import (
	"bytes"
	"fmt"
	"encoding/json"
	"testing"
	"os"
)

var appID = os.Getenv("YZAppID")
var appSecret = os.Getenv("YZAppSecret")
var client = YZClient{AppID: appID, AppSecret: appSecret}

func printTestName(testName string) {
	fmt.Println("Test: " + testName + "\n")
}

func printResult(ret string) {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(ret), "", "\t")
	if err != nil {
		fmt.Println(ret)
	} else {
		fmt.Println(out.String())
	}

}

func TestShopBasicGet(t *testing.T) {
	apiName := "kdt.shop.basic.get"
	printTestName(apiName)
	ret := client.Invoke(apiName, "1.0.0", "GET", map[string]string{}, map[string]string{})
	printResult(ret)
}

