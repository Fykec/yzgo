package yzgo

import (
	"errors"
	//"fmt"

	"fmt"
	"os"
	"strconv"
	"testing"
)

var accessToken = os.Getenv("YZAccessToken")
var client = YZClient{IsOAuth: true, AccessToken: accessToken}

func init() {
	if accessToken == "" {
		panic(errors.New("Please Fetch the Access token"))
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


func TestUsersWeixinFollowersPull(t *testing.T) {

	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0", "page_size": "10"})
	PrintObject(users)
}

func TestUsersWeixinFollowersGet(t *testing.T) {

	users, _ := client.UsersWeixinFollowersGet(Params{"start_follow": "2017-05-25 12:32:34", "end_follow": "2017-05-26 12:32:34"})
	PrintObject(users)
}

func TestUsersWeixinFollowerGets(t *testing.T) {
	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0"})
	users, _ = client.UsersWeixinFollowerGets(Params{"weixin_openids": users[0].WeixinOpenID})
	PrintObject(users)
}

func TestUsersWeixinFollowerGet(t *testing.T) {
	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0"})

	user, _ := client.UsersWeixinFollowerGet(Params{"weixin_openids": users[0].WeixinOpenID})
	PrintObject(user)
}

func TestUsersWeixinFollowerTagsAdd(t *testing.T) {
	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0"})
	user, _ := client.UsersWeixinFollowerTagsAdd(Params{"weixin_openids": users[0].WeixinOpenID, "tags": "测试"})
	PrintObject(user)
}

func TestUserBasicGet(t *testing.T) {
	ret, err := client.UserBasicGet()
	PrintResult(ret, err)
}
