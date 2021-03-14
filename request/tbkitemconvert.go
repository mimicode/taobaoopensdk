package request

import (
	"net/url"

	"github.com/mimicode/taobaoopensdk/utils"
)

//taobao.tbk.item.convert( 淘宝客商品链接转换 )
//http://open.taobao.com/api.htm?docId=24516&docType=2&scopeId=11653
type TbkItemConvertRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkItemConvertRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils.CheckNotNull(tk.Parameters.Get("num_iids"), "num_iids")
	utils.CheckMaxListSize(tk.Parameters.Get("num_iids"), 40, "num_iids")
}

//添加请求参数
func (tk *TbkItemConvertRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemConvertRequest) GetApiName() string {
	return "taobao.tbk.item.convert"
}

//返回请求参数
func (tk *TbkItemConvertRequest) GetParameters() url.Values {
	return *tk.Parameters
}
