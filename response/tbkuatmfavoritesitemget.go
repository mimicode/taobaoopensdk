package response

import "encoding/json"

//taobao.tbk.uatm.favorites.item.get( 获取淘宝联盟选品库的宝贝信息 )
type TbkUatmFavoritesItemGetResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkUatmFavoritesItemGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
