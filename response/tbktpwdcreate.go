package response

import "encoding/json"

//taobao.tbk.tpwd.create( 淘宝客淘口令 )
type TbkTpwdCreateResponse struct {
	TopResponse
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
