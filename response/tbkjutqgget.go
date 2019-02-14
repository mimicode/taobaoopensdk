package response

import "encoding/json"

//taobao.tbk.ju.tqg.get( 淘抢购api )
type TbkJuTqgGetResponse struct {
	TopResponse
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
