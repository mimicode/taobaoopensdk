package response

import "encoding/json"

//taobao.tbk.sc.invitecode.get( 淘宝客邀请码生成-社交 )
type TbkScInvitecodeGetResponse struct {
	TopResponse
	TbkScInvitecodeGetResult TbkScInvitecodeGetResult `json:"tbk_sc_invitecode_get_response"`
}

//解析输出结果
func (t *TbkScInvitecodeGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkScInvitecodeGetResult struct {
	Data      TbkScInvitecodeGetData   `json:"data"`
	RequestID string `json:"request_id"`
}

type TbkScInvitecodeGetData struct {
	InviterCode string `json:"inviter_code"`
}

