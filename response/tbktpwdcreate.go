package response

import "encoding/json"

//taobao.tbk.tpwd.create( 淘宝客淘口令 )
type TbkTpwdCreateResponse struct {
	TopResponse
	TbkTpwdCreateResult TbkTpwdCreateResult `json:"tbk_tpwd_create_response"`
}

//解析输出结果
func (t *TbkTpwdCreateResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkTpwdCreateResult struct {
	Data      TbkTpwdCreateData `json:"data"`
	RequestID string            `json:"request_id"`
}

type TbkTpwdCreateData struct {
	Model          string `json:"model"`
	PasswordSimple string `json:"password_simple"`
}
