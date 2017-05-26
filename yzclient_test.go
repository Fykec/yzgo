package yzgo

import (
	"bytes"
	"fmt"
	"encoding/json"
	"testing"
	"os"
	"strconv"
)

type ErrorResponse struct {
	Code float64 `json:code`
	Msg string `json:msg`
}

type RawResponse struct {
	Response map[string]interface{}
	ErrorResponse ErrorResponse `json:error_response`
}

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

func getJson(ret string) RawResponse {
	var jsonObject RawResponse
	json.Unmarshal([]byte(ret), &jsonObject)
	return  jsonObject
}

func TestShopBasicGet(t *testing.T) {
	apiName := "kdt.shop.basic.get"
	printTestName(apiName)
	ret := client.Invoke(apiName, "1.0.0", "GET", map[string]string{}, map[string]string{})
	printResult(ret)
}

func TestUsersWeixinFollowersPull(t *testing.T) {
	apiName := "youzan.users.weixin.followers.pull"
	printTestName(apiName)
	ret := client.Invoke(apiName, "3.0.0", "GET", map[string]string{"after_fans_id": "0"}, map[string]string{})
	printResult(ret)
}

func TestUsersWeixinFollowerGet(t *testing.T) {
	apiName := "youzan.users.weixin.followers.pull"
	printTestName(apiName)
	ret := client.Invoke(apiName, "3.0.0", "GET", map[string]string{"after_fans_id": "0"}, map[string]string{})
	jsonObject := getJson(ret)

	response := jsonObject.Response

	lastFansId := ""
	for k, v := range response {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "is ", vv)
			if k == "last_fans_id" {
				lastFansId = strconv.Itoa(int(vv))
			}
		default:
		}
	}

	fmt.Println("user id " + lastFansId)

	apiName = "youzan.users.weixin.follower.get"
	printTestName(apiName)
	ret = client.Invoke(apiName, "3.0.0", "GET", map[string]string{"fans_id": lastFansId}, map[string]string{})
	printResult(ret)
}

// youzan.trades.sold.get
func TestTradesSoldGet(t *testing.T) {
	apiName := "youzan.users.weixin.followers.pull"
	printTestName(apiName)
	ret := client.Invoke(apiName, "3.0.0", "GET", map[string]string{"after_fans_id": "0"}, map[string]string{})
	jsonObject := getJson(ret)

	response := jsonObject.Response

	lastFansId := ""
	for k, v := range response {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "is ", vv)
			if k == "last_fans_id" {
				lastFansId = strconv.Itoa(int(vv))
			}
		default:
		}
	}

	fmt.Println("user id " + lastFansId)

	apiName = "youzan.trades.sold.get"
	printTestName(apiName)
	ret = client.Invoke(apiName, "3.0.0", "GET", map[string]string{"fans_id": lastFansId}, map[string]string{})
	printResult(ret)
}


func TestTradeGet(t *testing.T) {
	apiName := "youzan.trade.get"
	printTestName(apiName)
	ret := client.Invoke(apiName, "3.0.0", "GET", map[string]string{}, map[string]string{})
	printResult(ret)
}
