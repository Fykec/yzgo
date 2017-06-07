package yzgo

import (
	"errors"
	//"fmt"

	"fmt"
	"os"
	"strconv"
	"testing"
)

var appID = os.Getenv("YZAppID")
var appSecret = os.Getenv("YZAppSecret")
var client = YZClient{AppID: appID, AppSecret: appSecret}

func init() {
	if appID == "" || appSecret == "" {
		panic(errors.New("Please setup the the YZAppID and YZAppSecret"))
	}
}

func printTestName(testName string) {
	fmt.Println("Test: " + testName + "\n")
}

/*
Shop Test
*/

func TestShopBasicGet(t *testing.T) {
	apiName := "kdt.shop.basic.get"
	printTestName(apiName)
	ret, err := client.Invoke(apiName, "1.0.0", "GET", Params{}, Params{})
	PrintResult(ret, err)
}

/*
Points Test
*/

func TestCrmFansPointsGet(t *testing.T) {
	apiName := "youzan.users.weixin.followers.pull"
	printTestName(apiName)
	ret, err := client.Invoke(apiName, "3.0.0", "GET", Params{"after_fans_id": "0"}, Params{})
	jsonObject, err := getRawResponse(ret)

	response := jsonObject.Response

	lastFansID := ""
	for k, v := range response {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "is ", vv)
			if k == "last_fans_id" {
				lastFansID = strconv.Itoa(int(vv))
			}
		default:
		}
	}

	fmt.Println("user id (last fans id) " + lastFansID)

	apiName = "youzan.crm.fans.points.get"
	printTestName(apiName)
	ret, err = client.Invoke(apiName, "3.0.0", "GET", Params{"fans_ids": lastFansID}, Params{})
	PrintResult(ret, err)
}

/*
Trade test
*/

// youzan.trades.sold.get
func TestTradesSoldGet(t *testing.T) {
	apiName := "youzan.users.weixin.followers.pull"
	printTestName(apiName)
	ret, err := client.Invoke(apiName, "3.0.0", "GET", Params{"after_fans_id": "0"}, Params{})
	jsonObject, err := getRawResponse(ret)

	response := jsonObject.Response

	lastFansID := ""
	for k, v := range response {
		switch vv := v.(type) {
		case float64:
			fmt.Println(k, "is ", vv)
			if k == "last_fans_id" {
				lastFansID = strconv.Itoa(int(vv))
			}
		default:
		}
	}

	fmt.Println("user id " + lastFansID)

	apiName = "youzan.trades.sold.get"
	printTestName(apiName)
	ret, err = client.Invoke(apiName, "3.0.0", "GET", Params{"fans_id": lastFansID}, Params{})
	PrintResult(ret, err)
}

func TestTradeGet(t *testing.T) {
	apiName := "youzan.trade.get"
	printTestName(apiName)
	ret, err := client.Invoke(apiName, "3.0.0", "GET", Params{}, Params{})
	PrintResult(ret, err)
}
