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