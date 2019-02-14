package response

import "encoding/json"

//taobao.tbk.sc.coupon.realtime.recommend( 好券直播API【社交】 )
type TbkScCouponRealtimeRecommendResponse struct {
	TopResponse
	TbkScCouponRealtimeRecommendResult TbkScCouponRealtimeRecommendResult `json:"tbk_sc_coupon_realtime_recommend_response"`
}

//解析输出结果
func (t *TbkScCouponRealtimeRecommendResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkScCouponRealtimeRecommendResult struct {
	Results   TbkScCouponRealtimeRecommendResults `json:"results"`
	RequestID string                              `json:"request_id"`
}

type TbkScCouponRealtimeRecommendResults struct {
	TbkCoupon []TbkScCouponRealtimeRecommendTbkCoupon `json:"tbk_coupon"`
}

type TbkScCouponRealtimeRecommendTbkCoupon struct {
	Category          int64                                   `json:"category"`
	CommissionRate    string                                  `json:"commission_rate"`
	CouponEndTime     string                                  `json:"coupon_end_time"`
	CouponInfo        string                                  `json:"coupon_info"`
	CouponRemainCount int64                                   `json:"coupon_remain_count"`
	CouponStartTime   string                                  `json:"coupon_start_time"`
	CouponTotalCount  int64                                   `json:"coupon_total_count"`
	ItemDescription   string                                  `json:"item_description"`
	ItemURL           string                                  `json:"item_url"`
	Nick              string                                  `json:"nick"`
	NumIid            int64                                   `json:"num_iid"`
	PictURL           string                                  `json:"pict_url"`
	SellerID          int64                                   `json:"seller_id"`
	ShopTitle         string                                  `json:"shop_title"`
	SmallImages       TbkScCouponRealtimeRecommendSmallImages `json:"small_images"`
	Title             string                                  `json:"title"`
	UserType          int64                                   `json:"user_type"`
	Volume            int64                                   `json:"volume"`
	ZkFinalPrice      string                                  `json:"zk_final_price"`
}

type TbkScCouponRealtimeRecommendSmallImages struct {
	String []string `json:"string"`
}
