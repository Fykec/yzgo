package yzgo

//Order the youzan order
type Order struct {
	Alias                 string              `json:"alias"`
	AllowSend             bool                `json:"allow_send"`
	BuyerMessages         []map[string]string `json:"buyer_messages"`
	DiscountFee           float64             `json:"discount_fee"`
	IsPresent             bool                `json:"is_present"`
	IsSend                bool                `json:"is_send"`
	IsVirtual             bool                `json:"is_virtual"`
	ItemID                string              `json:"item_id"`
	ItemRefundState       string              `json:"item_refund_state"`
	ItemType              float64             `json:"item_type"`
	Num                   float64             `json:"num"`
	OID                   string              `json:"oid"`
	OrderPromotionDetails interface{}         `json:"order_promotion_details"`
	OuterItemID           string              `json:"outer_item_id"`
	OuterSKUID            string              `json:"outer_sku_id"`
	Payment               float64             `json:"payment"`
	PicPath               string              `json:"pic_path"`
	PicThumbPath          string              `json:"pic_thumb_path"`
	Price                 float64             `json:"price"`
	RefundID              string              `json:"refund_id"`
	RefundFee             float64             `json:"refund_fee"`
	SellerNick            string              `json:"seller_nick"`
	SKUID                 string              `json:"sku_id"`
	SKUPropertiesName     string              `json:"sku_properties_name"`
	SKUUniqueCode         string              `json:"sku_unique_code"`
	StateStr              string              `json:"state_str"`
	Title                 string              `json:"title"`
	TotalFee              float64             `json:"total_fee"`
	Unit                  string              `json:"unit"`
}
