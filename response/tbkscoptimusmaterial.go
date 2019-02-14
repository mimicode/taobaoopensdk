package response

import "encoding/json"

//taobao.tbk.sc.optimus.material( 淘宝客擎天柱通用物料API - 社交 )
type TbkScOptimusMaterialResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkScOptimusMaterialResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
