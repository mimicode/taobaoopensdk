package response

import "encoding/json"

//taobao.tbk.dg.newuser.order.get( 淘宝客新用户订单API--导购 )
type TbkDgNewuserOrderGetResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkDgNewuserOrderGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
