package response

import "encoding/json"

//淘宝客商品详情查询（简版） 免费 不需要授权
type TbkPrivilegeGetResponse struct {
	TopResponse
	TbkPrivilegeGetResult TbkPrivilegeGetResult `json:"tbk_privilege_get_response"`
}

//解析输出结果
func (t *TbkPrivilegeGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkPrivilegeGetResult struct {
	Result    TbkPrivilegeGetResultData `json:"result"`
	RequestID string                    `json:"request_id"`
}

type TbkPrivilegeGetResultData struct {
	Data TbkPrivilegeGetData `json:"data"`
}

type TbkPrivilegeGetData struct {
	CategoryID        int64  `json:"category_id"`
	CouponClickURL    string `json:"coupon_click_url"`
	CouponEndTime     string `json:"coupon_end_time"`
	CouponInfo        string `json:"coupon_info"`
	CouponRemainCount int64  `json:"coupon_remain_count"`
	CouponStartTime   string `json:"coupon_start_time"`
	CouponTotalCount  int64  `json:"coupon_total_count"`
	CouponType        int64  `json:"coupon_type"`
	ItemID            int64  `json:"item_id"`
	ItemURL           string `json:"item_url"`
	MaxCommissionRate string `json:"max_commission_rate"`
}
