package yzgo

import (
	"bytes"
	//"fmt"
	"encoding/json"
	"testing"
	"os"
	"strconv"
	"fmt"
)

type ErrorResponse struct {
	Code float64 `json:"code"`
	Msg string `json:"msg"`
}

type RawResponse struct {
	Response map[string]interface{} `json:"response"`
	ErrorResponse ErrorResponse `json:"error_response"`
}

var appID = os.Getenv("YZAppID")
var appSecret = os.Getenv("YZAppSecret")
var client = YZClient{AppID: appID, AppSecret: appSecret}

func printTestName(testName string) {
	fmt.Println("Test: " + testName + "\n")
}

func printResult(ret string) {
	response := getJson(ret)
	// string to json help to have a better chinese printting https://segmentfault.com/q/1010000006778053
	jsonBytes, _ := json.Marshal(response)

	var out bytes.Buffer
	err := json.Indent(&out, jsonBytes, "", "  ")
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

/*
User Test
*/

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

func TestUsersWeixinFollowerGets(t *testing.T) {
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

	fmt.Println("user id (last fans id) " + lastFansId)

	apiName = "youzan.users.weixin.follower.gets"
	printTestName(apiName)
	ret = client.Invoke(apiName, "3.0.0", "GET", map[string]string{"fans_ids": lastFansId}, map[string]string{})
	printResult(ret)
}

func TestUsersWeixinFollowerTagsAdd(t *testing.T) {
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

	fmt.Println("user id (last fans id) " + lastFansId)

	apiName = "youzan.users.weixin.follower.tags.add"
	printTestName(apiName)
	ret = client.Invoke(apiName, "3.0.0", "GET", map[string]string{"fans_ids": lastFansId, "tags": "api_test"}, map[string]string{})
	printResult(ret)
}

func TestUserBasicGet(t *testing.T) {
	apiName := "youzan.user.basic.get"
	printTestName(apiName)
	ret := client.Invoke(apiName, "3.0.0", "GET", map[string]string{}, map[string]string{})
	printResult(ret)
}

/*
Shop Test
*/

func TestShopBasicGet(t *testing.T) {
	apiName := "kdt.shop.basic.get"
	printTestName(apiName)
	ret := client.Invoke(apiName, "1.0.0", "GET", map[string]string{}, map[string]string{})
	printResult(ret)
}

/*
Points Test
*/

func TestCrmFansPointsGet(t *testing.T) {
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

	fmt.Println("user id (last fans id) " + lastFansId)

	apiName = "youzan.crm.fans.points.get"
	printTestName(apiName)
	ret = client.Invoke(apiName, "3.0.0", "GET", map[string]string{"fans_ids": lastFansId}, map[string]string{})
	printResult(ret)
}

/*
Trade test
*/

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
