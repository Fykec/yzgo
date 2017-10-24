package yzgo

import (
	"encoding/json"
	"net/http"
)

/*
	User info introduction
	https://bbs.youzan.com/forum.php?mod=viewthread&tid=535089&extra=page%3D1%26filter%3Dtypeid%26typeid%3D160
*/

// User youzan user
type User struct {
	Avatar       string      `json:"avatar"`
	City         string      `json:"city"`
	FollowTime   YZTime      `json:"follow_time"`
	IsFollow     bool        `json:"is_follow"`
	LevelInfo    interface{} `json:"level_info"`
	Nick         string      `json:"nick"`
	Points       float64     `json:"points"`
	Province     string      `json:"province"`
	Sex          string      `json:"sex"`
	Tags         []Tag       `json:"tags"`
	TradedMoney  string      `json:"traded_money"`
	TradedNum    float64     `json:"traded_num"`
	UserID       int         `json:"user_id"`
	WeixinOpenID string      `json:"weixin_openid"`
}

//UsersWeixinFollowersPull "youzan.users.weixin.followers.pull"
func (c *YZClient) UsersWeixinFollowersPull(params Params) ([]User, error) {
	retBytes, err := c.Invoke("youzan.users.weixin.followers.pull", APIVersionDefault, http.MethodGet, params, Params{})
	rawResponse, err := getRawResponse(retBytes)
	if err != nil {
		panic(err)
	}

	respBytes, err := json.Marshal(rawResponse.Response["users"])

	//println(string(respBytes))

	var users []User
	err = json.Unmarshal(respBytes, &users)
	if err != nil {
		panic(err)
	}
	return users, err
}

//UsersWeixinFollowersGet "youzan.users.weixin.followers.get"
func (c *YZClient) UsersWeixinFollowersGet(params Params) ([]User, error) {
	retBytes, err := c.Invoke("youzan.users.weixin.followers.get", APIVersionDefault, http.MethodGet, params, Params{})
	rawResponse, err := getRawResponse(retBytes)

	if err != nil {
		panic(err)
	}
	respBytes, err := json.Marshal(rawResponse.Response["users"])

	var users []User
	err = json.Unmarshal(respBytes, &users)
	return users, err
}

//UsersWeixinFollowerGets "youzan.users.weixin.follower.gets"
func (c *YZClient) UsersWeixinFollowerGets(params Params) ([]User, error) {
	retBytes, err := c.Invoke("youzan.users.weixin.follower.gets", APIVersionDefault, http.MethodGet, params, Params{})
	rawResponse, err := getRawResponse(retBytes)

	if err != nil {
		panic(err)
	}
	respBytes, err := json.Marshal(rawResponse.Response["user"])

	var users []User
	err = json.Unmarshal(respBytes, &users)
	return users, err
}

//UsersWeixinFollowerGet "youzan.users.weixin.follower.get"
func (c *YZClient) UsersWeixinFollowerGet(params Params) (User, error) {
	retBytes, err := c.Invoke("youzan.users.weixin.follower.get", APIVersionDefault, http.MethodGet, params, Params{})
	rawResponse, err := getRawResponse(retBytes)

	if err != nil {
		panic(err)
	}
	respBytes, err := json.Marshal(rawResponse.Response["user"])

	var user User
	err = json.Unmarshal(respBytes, &user)
	return user, err
}

//UsersWeixinFollowerTagsAdd "youzan.users.weixin.follower.tags.add"
func (c *YZClient) UsersWeixinFollowerTagsAdd(params Params) (User, error) {
	retBytes, err := c.Invoke("youzan.users.weixin.follower.tags.add", APIVersionDefault, http.MethodGet, params, Params{})
	rawResponse, err := getRawResponse(retBytes)

	if err != nil {
		panic(err)
	}
	respBytes, err := json.Marshal(rawResponse.Response["user"])

	var user User
	err = json.Unmarshal(respBytes, &user)
	return user, err
}

//UserBasicGet ""
func (c *YZClient) UserBasicGet() ([]byte, error) {
	retBytes, err := c.Invoke("youzan.user.basic.get", APIVersionDefault, http.MethodGet, Params{}, Params{})
	rawResponse, err := getRawResponse(retBytes)

	if err != nil {
		panic(err)
	}
	respBytes, err := json.Marshal(rawResponse.Response)
	return respBytes, err
}
