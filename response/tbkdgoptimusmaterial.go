package response

import "encoding/json"

//taobao.tbk.dg.optimus.material( 淘宝客物料下行-导购 )
type TbkDgOptimusMaterialResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkDgOptimusMaterialResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
