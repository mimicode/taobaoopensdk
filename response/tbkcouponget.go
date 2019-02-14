package response

import "encoding/json"

//taobao.tbk.coupon.get( 阿里妈妈推广券信息查询 )
type TbkCouponGetResponse struct {
	TopResponse
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
