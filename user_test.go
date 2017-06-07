package yzgo

import (
	"os"
	"testing"
)

func TestUsersWeixinFollowersPull(t *testing.T) {
	var appID = os.Getenv("YZAppID")
	var appSecret = os.Getenv("YZAppSecret")
	var client = YZClient{AppID: appID, AppSecret: appSecret}

	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0"})
	PrintObject(users)
}

func TestUsersWeixinFollowersGet(t *testing.T) {
	var appID = os.Getenv("YZAppID")
	var appSecret = os.Getenv("YZAppSecret")
	var client = YZClient{AppID: appID, AppSecret: appSecret}

	users, _ := client.UsersWeixinFollowersGet(Params{"start_follow": "2017-05-25 12:32:34", "end_follow": "2017-05-26 12:32:34"})
	PrintObject(users)
}

func TestUsersWeixinFollowerGets(t *testing.T) {
	var appID = os.Getenv("YZAppID")
	var appSecret = os.Getenv("YZAppSecret")
	var client = YZClient{AppID: appID, AppSecret: appSecret}
	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0"})

	users, _ = client.UsersWeixinFollowerGets(Params{"weixin_openids": users[0].WeixinOpenID})
	PrintObject(users)
}

func TestUsersWeixinFollowerGet(t *testing.T) {
	var appID = os.Getenv("YZAppID")
	var appSecret = os.Getenv("YZAppSecret")
	var client = YZClient{AppID: appID, AppSecret: appSecret}
	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0"})

	user, _ := client.UsersWeixinFollowerGet(Params{"weixin_openids": users[0].WeixinOpenID})
	PrintObject(user)
}

func TestUsersWeixinFollowerTagsAdd(t *testing.T) {
	var appID = os.Getenv("YZAppID")
	var appSecret = os.Getenv("YZAppSecret")
	var client = YZClient{AppID: appID, AppSecret: appSecret}
	users, _ := client.UsersWeixinFollowersPull(Params{"after_fans_id": "0"})

	user, _ := client.UsersWeixinFollowerTagsAdd(Params{"weixin_openids": users[0].WeixinOpenID, "tags": "测试"})
	PrintObject(user)
}

func TestUserBasicGet(t *testing.T) {
	var appID = os.Getenv("YZAppID")
	var appSecret = os.Getenv("YZAppSecret")
	var client = YZClient{AppID: appID, AppSecret: appSecret}
	ret, err := client.UserBasicGet()
	PrintResult(ret, err)
}
