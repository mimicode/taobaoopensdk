package response

import "encoding/json"

//taobao.wireless.share.tpwd.query( 查询解析淘口令 )
type WirelessShareTpwdQueryResponse struct {
	TopResponse
	WirelessShareTpwdQueryResult WirelessShareTpwdQueryResult `json:"wireless_share_tpwd_query_response"`
}

//解析输出结果
func (t *WirelessShareTpwdQueryResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type WirelessShareTpwdQueryResult struct {
	Content     string `json:"content"`
	NativeURL   string `json:"native_url"`
	PicURL      string `json:"pic_url"`
	Suc         bool   `json:"suc"`
	ThumbPicURL string `json:"thumb_pic_url"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	RequestID   string `json:"request_id"`
}