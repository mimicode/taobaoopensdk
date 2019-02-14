package response

import "encoding/json"

//taobao.tbk.ju.tqg.get( 淘抢购api )
type TbkJuTqgGetResponse struct {
	TopResponse
	TbkJuTqgGetResult TbkJuTqgGetResult `json:"tbk_ju_tqg_get_response"`
}

//解析输出结果
func (t *TbkJuTqgGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkJuTqgGetResult struct {
	Results      TbkJuTqgGetResults `json:"results"`
	TotalResults int64              `json:"total_results"`
	RequestID    string             `json:"request_id"`
}

type TbkJuTqgGetResults struct {
	Results []TbkJuTqgGetResultResults `json:"results"`
}

type TbkJuTqgGetResultResults struct {
	CategoryName string `json:"category_name"`
	ClickURL     string `json:"click_url"`
	EndTime      string `json:"end_time"`
	NumIid       int64  `json:"num_iid"`
	PicURL       string `json:"pic_url"`
	ReservePrice string `json:"reserve_price"`
	SoldNum      int64  `json:"sold_num"`
	StartTime    string `json:"start_time"`
	Title        string `json:"title"`
	TotalAmount  int64  `json:"total_amount"`
	ZkFinalPrice string `json:"zk_final_price"`
}
