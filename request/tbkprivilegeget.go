package request

import (
	"net/url"

	"github.com/mimicode/taobaoopensdk/utils"
)

//taobao.tbk.privilege.get( 单品券高效转链API )
//http://open.taobao.com/api.htm?docId=28625&docType=2&scopeId=12403
type TbkPrivilegeGetRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkPrivilegeGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("adzone_id"), "adzone_id")
	utils.CheckNumber(tk.Parameters.Get("adzone_id"), "adzone_id")

	utils.CheckNotNull(tk.Parameters.Get("site_id"), "site_id")
	utils.CheckNumber(tk.Parameters.Get("site_id"), "site_id")

}

//添加请求参数
func (tk *TbkPrivilegeGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkPrivilegeGetRequest) GetApiName() string {
	return "taobao.tbk.privilege.get"
}

//返回请求参数
func (tk *TbkPrivilegeGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
