package response

import "encoding/json"

//taobao.tbk.sc.newuser.order.get( 淘宝客新用户订单API--社交 )
type TbkScNewuserOrderGetResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkScNewuserOrderGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
