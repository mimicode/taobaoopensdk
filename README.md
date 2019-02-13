# TAOBAO OPEN SDK

## Installation

Download and install it:

```sh
$ go get -u github.com/mimicode/taobaoopensdk
```

Import it in your code:

```go
import "github.com/mimicode/taobaoopensdk"
```

Examples
```
appKey := ""
appSecret := ""
sessionKey := ""
//初始化TopClient
client := &TopClient{}
client.Init(appKey, appSecret, sessionKey)

//初始化请求接口信息
getRequest := &request.TbkItemInfoGetRequest{}
getRequest.AddParameter("num_iids", "583866215568,578307080718")
getRequest.AddParameter("platform", "1")
getRequest.AddParameter("ip", "11.22.33.43")
//初始化结果类型
var getResponse DefaultResponse = &response.TbkItemInfoGetResponse{}
//执行请求接口得到结果
err := client.Exec(getRequest, getResponse)
if err != nil {
    t.Log(err)
} else {
    tbkItemInfoGetResponse := getResponse.(*response.TbkItemInfoGetResponse)

    if tbkItemInfoGetResponse.IsError() {
        fmt.Println(tbkItemInfoGetResponse.Body)
    } else {
        items := tbkItemInfoGetResponse.TbkItemInfoGetResult.Results.NTbkItem
        for _, v := range items {
            fmt.Println(v.Title)
        }
    }

}
```

API support list
- [taobao.tbk.item.convert( 淘宝客商品链接转换 )](http://open.taobao.com/api.htm?docId=24516&docType=2&scopeId=11653)
- [taobao.tbk.item.info.get 淘宝客商品详情查询（简版）](http://open.taobao.com/api.htm?docId=24518&docType=2)
- [taobao.tbk.privilege.get( 单品券高效转链API )](http://open.taobao.com/api.htm?docId=28625&docType=2&scopeId=12403)
- [taobao.tbk.sc.activitylink.toolget( 淘宝联盟官方活动推广API-工具 )](http://open.taobao.com/api.htm?docId=41921&docType=2&scopeId=15675)
- [taobao.tbk.sc.material.optional( 通用物料搜索API )](http://open.taobao.com/api.htm?docId=35263&docType=2&scopeId=13991)
- [taobao.tbk.sc.order.get( 淘宝客订单查询 - 社交 )](http://open.taobao.com/api.htm?docId=38078&docType=2&scopeId=14814)
- [taobao.tbk.shop.convert( 淘宝客店铺链接转换 )](http://open.taobao.com/api.htm?docId=24523&docType=2&scopeId=11653)
- [taobao.tbk.tpwd.convert( 淘口令转链 )](http://open.taobao.com/api.htm?docId=32932&docType=2&scopeId=11653)
