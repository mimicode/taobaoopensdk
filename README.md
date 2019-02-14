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
- [taobao.sellercats.list.get( 获取前台展示的店铺内卖家自定义商品类目 )](http://open.taobao.com/api.htm?docId=65&docType=2&scopeId=386)
- [taobao.shopcats.list.get( 获取前台展示的店铺类目 )](http://open.taobao.com/api.htm?docId=64&docType=2&scopeId=386)
- [taobao.shop.getbytitle( 根据店铺名称获取店铺信息 )](http://open.taobao.com/api.htm?docId=24852&docType=2&scopeId=386)
- [taobao.tbk.item.convert( 淘宝客商品链接转换 )](http://open.taobao.com/api.htm?docId=24516&docType=2&scopeId=11653)
- [taobao.tbk.item.coupon.get( 单品加券检索api )](http://open.taobao.com/api.htm?docId=28110&docType=2&scopeId=12332)
- [taobao.tbk.item.get( 淘宝客商品查询 )](http://open.taobao.com/api.htm?docId=24515&docType=2&scopeId=11655)
- [taobao.tbk.item.info.get( 淘宝客商品详情（简版） )](http://open.taobao.com/api.htm?docId=24518&docType=2)
- [taobao.tbk.item.recommend.get( 淘宝客商品关联推荐查询 )](http://open.taobao.com/api.htm?docId=24517&docType=2&scopeId=11655)
- [taobao.tbk.privilege.get( 单品券高效转链API )](http://open.taobao.com/api.htm?docId=28625&docType=2&scopeId=12403)
- [taobao.tbk.sc.activitylink.toolget( 淘宝联盟官方活动推广API-工具 )](http://open.taobao.com/api.htm?docId=41921&docType=2&scopeId=15675)
- [taobao.tbk.sc.coupon.brand.recommend( 品牌券API【社交】 )](http://open.taobao.com/api.htm?docId=29823&docType=2&scopeId=12331)
- [taobao.tbk.sc.coupon.realtime.recommend( 好券直播API【社交】 )](http://open.taobao.com/api.htm?docId=29820&docType=2&scopeId=12331)
- [taobao.tbk.sc.invitecode.get( 淘宝客邀请码生成-社交 )](http://open.taobao.com/api.htm?docId=38046&docType=2&scopeId=14474)
- [taobao.tbk.sc.material.optional( 通用物料搜索API )](http://open.taobao.com/api.htm?docId=35263&docType=2&scopeId=13991)
- [taobao.tbk.sc.order.get( 淘宝客订单查询 - 社交 )](http://open.taobao.com/api.htm?docId=38078&docType=2&scopeId=14814)
- [taobao.tbk.sc.publisher.info.get( 淘宝客信息查询 - 社交 )](http://open.taobao.com/api.htm?docId=37989&docType=2&scopeId=14474)
- [taobao.tbk.sc.publisher.info.save( 淘宝客渠道信息备案 - 社交 )](http://open.taobao.com/api.htm?docId=37988&docType=2&scopeId=14474)
- [taobao.tbk.shop.convert( 淘宝客店铺链接转换 )](http://open.taobao.com/api.htm?docId=24523&docType=2&scopeId=11653)
- [taobao.tbk.spread.get( 物料传播方式获取 )](http://open.taobao.com/api.htm?docId=27832&docType=2&scopeId=12340)
- [taobao.tbk.tpwd.convert( 淘口令转链 )](http://open.taobao.com/api.htm?docId=32932&docType=2&scopeId=11653)
- [taobao.wireless.share.tpwd.query( 查询解析淘口令 )](http://open.taobao.com/api.htm?docId=32461&docType=2&scopeId=11998)
