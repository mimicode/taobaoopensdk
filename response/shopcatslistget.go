package response

import "encoding/json"

//taobao.shopcats.list.get( 获取前台展示的店铺类目 )
type ShopcatsListGetResponse struct {
	TopResponse
	ShopcatsListGetResult ShopcatsListGetResult `json:"shopcats_list_get_response"`
}

//解析输出结果
func (t *ShopcatsListGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type ShopcatsListGetResult struct {
	ShopCats  ShopcatsListGetShopCats `json:"shop_cats"`
	RequestID string   `json:"request_id"`
}

type ShopcatsListGetShopCats struct {
	ShopCat []ShopcatsListGetShopCat `json:"shop_cat"`
}

type ShopcatsListGetShopCat struct {
	Cid       int64  `json:"cid"`
	IsParent  bool   `json:"is_parent"`
	Name      string `json:"name"`
	ParentCid int64  `json:"parent_cid"`
}