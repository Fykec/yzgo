package yzgo

import (
	"time"
)

//Trade youzan trade info
type Trade struct {
	AdjustFee           AdjustFee   `json:"adjust_fee"`
	BuyWayStr           string      `json:"buy_way_str"`
	BuyerArea           string      `json:"buyer_area"`
	BuyerMessage        string      `json:"buyer_message"`
	ConsignTime         time.Time   `json:"consign_time"`
	CouponDetails       interface{} `json:"coupon_details"`
	Created             time.Time   `json:"created"`
	DeliveryEndTime     float64     `json:"delivery_end_time"`
	DeliveryStartTime   float64     `json:"delivery_start_time"`
	DeliveryTimeDisplay string      `json:"delivery_time_display"`
	DiscountFee         float64     `json:"discount_fee"`
	FansInfo            FansInfo    `json:"fans_info"`
	Feedback            float64     `json:"feedback"`
	FeedbackNum         float64     `json:"feedback_num"`
	FetchDetail         interface{} `json:"fetch_detail"`
	GoodsKind           float64     `json:"good_kind"`
	GroupNo             string      `json:"group_no"`
	HotelInfo           string      `json:"hotel_info"`
	IsTuanHead          float64     `json:"is_tuan_head"`
	ItemID              string      `json:"item_id"`
	Kind                float64     `json:"kind"`
	Lat                 float64     `json:"lat"`
	Lng                 float64     `json:"lng"`
	Num                 float64     `json:"num"`
	OfflineID           string      `json:"offline_id"`
	OrderType           float64     `json:"order_type"`
	Orders              []Order     `json:"orders"`
	OriginalPostFee     float64     `json:"original_post_fee"`
	OutTradeNo          interface{} `json:"out_trade_no"`
	OuterTID            string      `json:"outer_tid"`
	OuterUserID         string      `json:"outer_user_id"`
	PayTime             time.Time   `json:"pay_time"`
	PayType             string      `json:"pay_type"`
	Payment             float64     `json:"payment"`
	PeriodOrderDetail   interface{} `json:"period_order_detail"`
	PFBuyWayStr         string      `json:"pf_buy_way_str"`
	PicPath             string      `json:"pic_path"`
	PicThumbPath        string      `json:"pic_thumb_path"`
	PointsPrice         float64     `json:"points_price"`
	PostFee             float64     `json:"post_fee"`
	Price               float64     `json:"price"`
	PromotionDetals     interface{} `json:"promotion_details"`
	QRID                float64     `json:"qr_id"`
	ReceiverAddress     string      `json:"receiver_address"`
	ReceiverDistrict    string      `json:"receiver_district"`
	ReceiverMobile      string      `json:"receiver_mobile"`
	ReceiverName        string      `json:"receiver_name"`
	Receiverstate       string      `json:"reciever_state"`
	ReceiverZip         string      `json:"receiver_zip"`
	RefundState         string      `json:"refund_state"`
	RefundFee           string      `json:"refund_fee"`
	RelationType        string      `json:"relation_type"`
	Relations           interface{} `json:"relations"`
	SellerFlag          float64     `json:"seller_flag"`
	SendNum             float64     `json:"send_num"`
	ServicePhone        string      `json:"service_phone"`
	ShippingType        string      `json:"express"`
	ShopID              string      `json:"shop_id"`
	ShopType            float64     `json:"shop_type"`
	SignTime            time.Time   `json:"sign_time"`
	Status              string      `json:"status"`
	StatusStr           string      `json:"status_str"`
	SubTrades           []Trade     `json:"sub_trades"`
	TID                 string      `json:"tid"`
	Title               string      `json:"title"`
	TotalFee            float64     `json:"total_fee"`
	TradeMemo           string      `json:"trade_memo"`
	TransactionTID      string      `json:"transaction_tid"`
	TuanNo              string      `json:"tuan_no"`
	Type                string      `json:"type"`
	UpdateTime          time.Time   `json:"update_time"`
}
