package response

import "encoding/json"

//taobao.tbk.dg.material.optional( 通用物料搜索API（导购） )
type TbkDgMaterialOptionalResponse struct {
	TopResponse
}

//解析输出结果
func (t *TbkDgMaterialOptionalResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}
