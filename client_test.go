package taobaoopensdk

import (
	"fmt"
	"os"
	"taobaoopensdk/request"
	"taobaoopensdk/response"
	"testing"
)

func TestTbkiteminfoget(t *testing.T) {

	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
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
		result := getResponse.(*response.TbkItemInfoGetResponse)

		if result.IsError() {
			fmt.Println(result.Body)
		} else {
			items := result.TbkItemInfoGetResult.Results.NTbkItem
			for _, v := range items {
				fmt.Println(v.Title)
			}
		}

	}
}

func TestTbkPrivilegeGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkPrivilegeGetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")
	getRequest.AddParameter("item_id", "583866215568")

	getRequest.AddParameter("relation_id", "")
	getRequest.AddParameter("me", "")
	getRequest.AddParameter("platform", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkPrivilegeGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkPrivilegeGetResponse)

		fmt.Println(result.TbkPrivilegeGetResult.Result.Data)

	}
}

func TestTbkScMaterialOptional(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScMaterialOptionalRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")

	getRequest.AddParameter("q", "女装")
	getRequest.AddParameter("platform", "1")
	getRequest.AddParameter("page_no", "1")
	getRequest.AddParameter("page_size", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScMaterialOptionalResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScMaterialOptionalResponse)

		fmt.Println(result.TbkScMaterialOptionalResult)

	}
}

func TestTbkScActivitylinkToolget(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScActivitylinkToolgetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")

	getRequest.AddParameter("promotion_scene_id", "8479247")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScActivitylinkToolgetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScActivitylinkToolgetResponse)

		fmt.Println(result.TbkScActivitylinkToolgetResult)

	}
}

func TestTbkItemConvert(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkItemConvertRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("fields", "num_iid,click_url")

	getRequest.AddParameter("num_iids", "583866215568,578307080718")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkItemConvertResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkItemConvertResponse)

		fmt.Println(result.TbkItemConvertResult.Results.NTbkItem)

	}
}

func TestTbkShopConvert(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkShopConvertRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("fields", "user_id,click_url")

	getRequest.AddParameter("user_ids", "188124207,383602586")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkShopConvertResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkShopConvertResponse)

		fmt.Println(result.Body)

	}
}
