package response

import "encoding/json"

//taobao.tbk.tpwd.convert( 淘口令转链 )
type TbkTpwdConvertResponse struct {
	TopResponse
	TbkTpwdConvertResult TbkTpwdConvertResult `json:"tbk_tpwd_convert_response"`
}

//解析输出结果
func (t *TbkTpwdConvertResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkTpwdConvertResult struct {
	Data      TbkTpwdConvertData `json:"data"`
	RequestID string             `json:"request_id"`
}

type TbkTpwdConvertData struct {
	ClickURL string `json:"click_url"`
	NumIid   string `json:"num_iid"`
}
