package request

import (
	"net/url"

	"github.com/mimicode/taobaoopensdk/utils"
)

//taobao.tbk.item.recommend.get( 淘宝客商品关联推荐查询 )
//http://open.taobao.com/api.htm?docId=24517&docType=2&scopeId=11655
type TbkItemRecommendGetRequest struct {
	Parameters *url.Values //请求参数
}

//参数检测
func (tk *TbkItemRecommendGetRequest) CheckParameters() {
	utils.CheckNotNull(tk.Parameters.Get("fields"), "fields")
	utils.CheckNotNull(tk.Parameters.Get("num_iid"), "num_iid")
}

//添加请求参数
func (tk *TbkItemRecommendGetRequest) AddParameter(key, val string) {
	if tk.Parameters == nil {
		tk.Parameters = &url.Values{}
	}
	tk.Parameters.Add(key, val)
}

//返回接口名称
func (tk *TbkItemRecommendGetRequest) GetApiName() string {
	return "taobao.tbk.item.recommend.get"
}

//返回请求参数
func (tk *TbkItemRecommendGetRequest) GetParameters() url.Values {
	return *tk.Parameters
}
