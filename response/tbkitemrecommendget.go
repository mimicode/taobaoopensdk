package response

import "encoding/json"

//taobao.tbk.item.recommend.get( 淘宝客商品关联推荐查询 )
type TbkItemRecommendGetResponse struct {
	TopResponse
	TbkItemRecommendGetResult TbkItemRecommendGetResult `json:"tbk_item_recommend_get_response"`
}

//解析输出结果
func (t *TbkItemRecommendGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkItemRecommendGetResult struct {
	Results   TbkItemRecommendGetResults `json:"results"`
	RequestID string                     `json:"request_id"`
}

type TbkItemRecommendGetResults struct {
	NTbkItem []TbkItemRecommendGetNTbkItem `json:"n_tbk_item"`
}

type TbkItemRecommendGetNTbkItem struct {
	ItemURL      string                         `json:"item_url"`
	Nick         string                         `json:"nick"`
	NumIid       int64                          `json:"num_iid"`
	PictURL      string                         `json:"pict_url"`
	Provcity     string                         `json:"provcity"`
	ReservePrice string                         `json:"reserve_price"`
	SellerID     int64                          `json:"seller_id"`
	SmallImages  TbkItemRecommendGetSmallImages `json:"small_images"`
	Title        string                         `json:"title"`
	UserType     int64                          `json:"user_type"`
	Volume       int64                          `json:"volume"`
	ZkFinalPrice string                         `json:"zk_final_price"`
}

type TbkItemRecommendGetSmallImages struct {
	String []string `json:"string"`
}
