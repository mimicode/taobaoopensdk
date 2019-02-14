package response

import "encoding/json"

//taobao.tbk.shop.get( 淘宝客店铺查询 )
type TbkShopGetResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkShopGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
