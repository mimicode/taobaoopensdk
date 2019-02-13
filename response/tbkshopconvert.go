package response

import "encoding/json"

//taobao.tbk.shop.convert( 淘宝客店铺链接转换 )
type TbkShopConvertResponse struct {
	TopResponse
	TbkShopConvertResult TbkShopConvertResult `json:"tbk_shop_convert_response"`
}

//解析输出结果
func (t *TbkShopConvertResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkShopConvertResult struct {
	Results   TbkShopConvertResults `json:"results"`
	RequestID string                `json:"request_id"`
}

type TbkShopConvertResults struct {
	NTbkShop []TbkShopConvertNTbkShop `json:"n_tbk_shop"`
}

type TbkShopConvertNTbkShop struct {
	ClickURL string `json:"click_url"`
	UserID   int64  `json:"user_id"`
}
