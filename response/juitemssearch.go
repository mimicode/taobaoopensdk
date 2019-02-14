package response

import "encoding/json"

//taobao.ju.items.search( 聚划算商品搜索接口 )
type JuItemsSearchResponse struct {
	TopResponse
}

//解析输出结果
func (t *JuItemsSearchResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
