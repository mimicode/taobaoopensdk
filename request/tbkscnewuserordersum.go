package request

import (
	"net/url"

	"github.com/mimicode/taobaoopensdk/utils"
)

//taobao.tbk.sc.newuser.order.sum( 拉新活动汇总API--社交 )
//http://open.taobao.com/api.htm?docId=36837&docType=2&scopeId=11655
type TbkScNewuserOrderSumRequest struct {
	Parameters *url.Values //请求参数
}

func (tk *TbkScNewuserOrderSumRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("activity_id"), "activity_id")
	utils.CheckNotNull(tk.Parameters.Get("page_no"), "page_no")
	utils.CheckNumber(tk.Parameters.Get("page_no"), "page_no")
	utils.CheckNotNull(tk.Parameters.Get("page_size"), "page_size")
	utils.CheckNumber(tk.Parameters.Get("page_size"), "page_size")

}

//添加请求参数
func (tk *TbkScNewuserOrderSumRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkScNewuserOrderSumRequest) GetApiName() string {
	return "taobao.tbk.sc.newuser.order.sum"
}

//返回请求参数
func (tk *TbkScNewuserOrderSumRequest) GetParameters() url.Values {
	return *tk.Parameters
}
