package response

import "encoding/json"

//taobao.tbk.item.coupon.get( 单品加券检索api )
type TbkItemCouponGetResponse struct {
	TopResponse
	TbkItemCouponGetResult TbkItemCouponGetResult `json:"tbk_item_coupon_get_response"`
}

//解析输出结果
func (t *TbkItemCouponGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkItemCouponGetResult struct {
	Results      TbkItemCouponGetResults `json:"results"`
	TotalResults int64                   `json:"total_results"`
	RequestID    string                  `json:"request_id"`
}

type TbkItemCouponGetResults struct {
	TbkCoupon []TbkItemCouponGetTbkCoupon `json:"tbk_coupon"`
}

type TbkItemCouponGetTbkCoupon struct {
	Category          int64                        `json:"category"`
	CommissionRate    string                       `json:"commission_rate"`
	CouponClickURL    string                       `json:"coupon_click_url"`
	CouponEndTime     string                       `json:"coupon_end_time"`
	CouponInfo        string                       `json:"coupon_info"`
	CouponRemainCount int64                        `json:"coupon_remain_count"`
	CouponStartTime   string                       `json:"coupon_start_time"`
	CouponTotalCount  int64                        `json:"coupon_total_count"`
	ItemDescription   string                       `json:"item_description"`
	ItemURL           string                       `json:"item_url"`
	Nick              string                       `json:"nick"`
	NumIid            int64                        `json:"num_iid"`
	PictURL           string                       `json:"pict_url"`
	SellerID          int64                        `json:"seller_id"`
	ShopTitle         string                       `json:"shop_title"`
	SmallImages       *TbkItemCouponGetSmallImages `json:"small_images,omitempty"`
	Title             string                       `json:"title"`
	UserType          int64                        `json:"user_type"`
	Volume            int64                        `json:"volume"`
	ZkFinalPrice      string                       `json:"zk_final_price"`
}

type TbkItemCouponGetSmallImages struct {
	String []string `json:"string"`
}
