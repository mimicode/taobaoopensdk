package response

import "encoding/json"

//taobao.tbk.activitylink.get( 淘宝联盟官方活动推广API-媒体 )
type TbkActivitylinkGetResponse struct {
	TopResponse
	TbkActivitylinkGetResult TbkActivitylinkGetResult `json:"tbk_activitylink_get_response"`
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

type TbkActivitylinkGetResult struct {
	BizErrorCode int64  `json:"biz_error_code"`
	Data         string `json:"data"`
	ResultCode   int64  `json:"result_code"`
	RequestID    string `json:"request_id"`
}