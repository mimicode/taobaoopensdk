package response

import "encoding/json"

//taobao.tbk.item.convert( 淘宝客商品链接转换 )
type TbkItemConvertResponse struct {
	TopResponse
	TbkItemConvertResult TbkItemConvertResult `json:"tbk_item_convert_response"`
}

//解析输出结果
func (t *TbkItemConvertResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkItemConvertResult struct {
	Results   TbkItemConvertResults `json:"results"`
	RequestID string                `json:"request_id"`
}

type TbkItemConvertResults struct {
	NTbkItem []TbkItemConvertNTbkItem `json:"n_tbk_item"`
}

type TbkItemConvertNTbkItem struct {
	ClickURL string `json:"click_url"`
	NumIid   int64  `json:"num_iid"`
}
