package request

import (
	"net/url"
)

//淘宝客商品详情查询（简版） 免费 不需要授权
type TbkItemInfoGetRequest struct {
	Parameters *url.Values //请求参数
}

//添加请求参数
func (tk *TbkItemInfoGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemInfoGetRequest) GetApiName() string {
	return "taobao.tbk.item.info.get"
}

//返回请求参数
func (tk *TbkItemInfoGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
