package response

import "encoding/json"

//taobao.tbk.dg.newuser.order.sum( 拉新活动汇总API--导购 )
type TbkDgNewuserOrderSumResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkDgNewuserOrderSumResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
