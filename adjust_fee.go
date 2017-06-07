package yzgo

//AdjustFee the adjust fee in trade
type AdjustFee struct {
	Change     float64 `json:"change"`
	PayChange  float64 `json:"pay_change"`
	PostChange float64 `json:"post_change"`
}
