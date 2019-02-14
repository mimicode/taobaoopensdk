package response

import "encoding/json"

//taobao.tbk.spread.get( 物料传播方式获取 )
type TbkSpreadGetResponse struct {
	TopResponse
	TbkSpreadGetResult TbkSpreadGetResult `json:"tbk_spread_get_response"`
}

//解析输出结果
func (t *TbkSpreadGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkSpreadGetResult struct {
	Results      TbkSpreadGetResults `json:"results"`
	TotalResults int64               `json:"total_results"`
	RequestID    string              `json:"request_id"`
}

type TbkSpreadGetResults struct {
	TbkSpread []TbkSpreadGetTbkSpread `json:"tbk_spread"`
}

type TbkSpreadGetTbkSpread struct {
	Content string `json:"content"`
	ErrMsg  string `json:"err_msg"`
}
