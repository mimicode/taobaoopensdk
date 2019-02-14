package response

import "encoding/json"

//taobao.tbk.activitylink.get( 淘宝联盟官方活动推广API-媒体 )
type TbkActivitylinkGetResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkActivitylinkGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
