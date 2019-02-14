package response

import "encoding/json"

//taobao.tbk.sc.publisher.info.save( 淘宝客渠道信息备案 - 社交 )
type TbkScPublisherInfoSaveResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkScPublisherInfoSaveResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
