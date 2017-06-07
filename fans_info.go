package yzgo

//FansInfo the user(fans is the new user) info summary
type FansInfo struct {
	BuyerID      string `json:"buyer_id"`
	FansID       string `json:"fans_id"`
	FansNickname string `json:"fans_nickname"`
	FansType     string `json:"fans_type"`
}
