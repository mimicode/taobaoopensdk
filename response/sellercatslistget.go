package response

import "encoding/json"

//taobao.sellercats.list.get( 获取前台展示的店铺内卖家自定义商品类目 )
type SellercatsListGetResponse struct {
	TopResponse
	SellercatsListGetResult SellercatsListGetResult `json:"sellercats_list_get_response"`
}

//解析输出结果
func (t *SellercatsListGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type SellercatsListGetResult struct {
	SellerCats SellercatsListGetSellerCats `json:"seller_cats"`
	RequestID  string                      `json:"request_id"`
}

type SellercatsListGetSellerCats struct {
	SellerCat []SellercatsListGetSellerCat `json:"seller_cat"`
}

type SellercatsListGetSellerCat struct {
	Cid       int64  `json:"cid"`
	Name      string `json:"name"`
	ParentCid int64  `json:"parent_cid"`
	PicURL    string `json:"pic_url"`
	SortOrder int64  `json:"sort_order"`
	Type      string `json:"type"`
}
