package response

import "encoding/json"

//taobao.tbk.coupon.get( 阿里妈妈推广券信息查询 )
type TbkCouponGetResponse struct {
	TopResponse
	TbkCouponGetResult TbkCouponGetResult `json:"tbk_coupon_get_response"`
}

//解析输出结果
func (t *TbkCouponGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkCouponGetResult struct {
	Data      TbkCouponGetData `json:"data"`
	RequestID string           `json:"request_id"`
}

type TbkCouponGetData struct {
	CouponActivityID  string `json:"coupon_activity_id"`
	CouponAmount      string `json:"coupon_amount"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponRemainCount int64  `json:"coupon_remain_count"`
	CouponSrcScene    int64  `json:"coupon_src_scene"`
	CouponStartFee    string `json:"coupon_start_fee"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponTotalCount  int64  `json:"coupon_total_count"`
	CouponType        int64  `json:"coupon_type"`
}