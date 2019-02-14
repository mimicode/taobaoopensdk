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

		fmt.Println(result.TbkShopConvertResult.Results)

	}
}

func TestTbkTpwdConvert(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkTpwdConvertRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("password_content", "(7q8mbttcjzE)")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkTpwdConvertResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkTpwdConvertResponse)

		fmt.Println(result.TbkTpwdConvertResult.Data)

	}
}

func TestTbkScOrderGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScOrderGetRequest{}
	getRequest.AddParameter("fields", "trade_parent_id,trade_id,num_iid,item_title,item_num,price,pay_price,seller_nick,seller_shop_title,commission,commission_rate,unid,create_time,earning_time,tk_status,tk3rd_type,tk3rd_pub_id,order_type,income_rate,pub_share_pre_fee,subsidy_rate,subsidy_type,terminal_type,auction_category,site_id,site_name,adzone_id,adzone_name,alipay_total_price,total_commission_rate,total_commission_fee,subsidy_fee,relation_id,click_time")
	getRequest.AddParameter("start_time", "2019-02-13 20:50:00")
	getRequest.AddParameter("span", "600")
	getRequest.AddParameter("order_query_type", "create_time")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScOrderGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScOrderGetResponse)

		fmt.Println(result.TbkScOrderGetResult.Results)

	}
}

func TestWirelessShareTpwdQuery(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.WirelessShareTpwdQueryRequest{}
	getRequest.AddParameter("password_content", "(7q8mbttcjzE)")

	//初始化结果类型
	var getResponse DefaultResponse = &response.WirelessShareTpwdQueryResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.WirelessShareTpwdQueryResponse)

		fmt.Println(result.WirelessShareTpwdQueryResult)

	}
}

func TestShopcatsListGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.ShopcatsListGetRequest{}
	getRequest.AddParameter("fields", "cid,parent_cid,name,is_parent")

	//初始化结果类型
	var getResponse DefaultResponse = &response.ShopcatsListGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.ShopcatsListGetResponse)

		fmt.Println(result.ShopcatsListGetResult.ShopCats)

	}
}

func TestSellercatsListGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.SellercatsListGetRequest{}
	getRequest.AddParameter("nick", "韩都衣舍旗舰店")

	//初始化结果类型
	var getResponse DefaultResponse = &response.SellercatsListGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.SellercatsListGetResponse)

		fmt.Println(result.SellercatsListGetResult.SellerCats)

	}
}

func TestShopGetbytitle(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.ShopGetbytitleRequest{}
	getRequest.AddParameter("title", "韩都衣舍旗舰店")

	//初始化结果类型
	var getResponse DefaultResponse = &response.ShopGetbytitleResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.ShopGetbytitleResponse)

		fmt.Println(result.ShopGetbytitleResult.ResultShop)

	}
}

func TestTbkScInvitecodeGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScInvitecodeGetRequest{}
	getRequest.AddParameter("relation_app", "common")
	getRequest.AddParameter("code_type", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScInvitecodeGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScInvitecodeGetResponse)

		fmt.Println(result.TbkScInvitecodeGetResult.Data)

	}
}

//(未完成)
func TestTbkScPublisherInfoGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScPublisherInfoGetRequest{}
	getRequest.AddParameter("relation_app", "common")
	getRequest.AddParameter("info_type", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScPublisherInfoGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScPublisherInfoGetResponse)

		fmt.Println(result.Body)

	}
}

//(未完成)
func TestTbkScPublisherInfoSave(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScPublisherInfoSaveRequest{}
	getRequest.AddParameter("inviter_code", "common")
	getRequest.AddParameter("info_type", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScPublisherInfoGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScPublisherInfoGetResponse)

		fmt.Println(result.Body)

	}
}
