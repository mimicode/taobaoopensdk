package response

import "encoding/json"

//taobao.tbk.item.get( 淘宝客商品查询 )
type TbkItemGetResponse struct {
	TopResponse
	TbkItemGetResult TbkItemGetResult `json:"tbk_item_get_response"`
}

//解析输出结果
func (t *TbkItemGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkItemGetResult struct {
	Results      TbkItemGetResults `json:"results"`
	TotalResults int64   `json:"total_results"`
	RequestID    string  `json:"request_id"`
}

type TbkItemGetResults struct {
	NTbkItem []TbkItemGetNTbkItem `json:"n_tbk_item"`
}

type TbkItemGetNTbkItem struct {
	ItemURL      string      `json:"item_url"`
	Nick         string      `json:"nick"`
	NumIid       int64       `json:"num_iid"`
	PictURL      string      `json:"pict_url"`
	Provcity     string      `json:"provcity"`
	ReservePrice string      `json:"reserve_price"`
	SellerID     int64       `json:"seller_id"`
	SmallImages  TbkItemGetSmallImages `json:"small_images"`
	Title        string      `json:"title"`
	UserType     int64       `json:"user_type"`
	Volume       int64       `json:"volume"`
	ZkFinalPrice string      `json:"zk_final_price"`
}

type TbkItemGetSmallImages struct {
	String []string `json:"string"`
}