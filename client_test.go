package taobaoopensdk

import (
	"fmt"
	"os"
	"testing"

	"github.com/mimicode/taobaoopensdk/request"
	"github.com/mimicode/taobaoopensdk/response"
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

func TestTbkScCouponRealtimeRecommend(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScCouponRealtimeRecommendRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScCouponRealtimeRecommendResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScCouponRealtimeRecommendResponse)

		fmt.Println(result.TbkScCouponRealtimeRecommendResult.Results.TbkCoupon)

	}
}

func TestTbkScCouponBrandRecommend(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkScCouponBrandRecommendRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("site_id", "7418269")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkScCouponBrandRecommendResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkScCouponBrandRecommendResponse)

		fmt.Println(result.TbkScCouponBrandRecommendResult.Results.TbkCoupon)

	}
}

func TestTbkSpreadGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkSpreadGetRequest{}
	getRequest.AddParameter("requests", `[{"url":"https://uland.taobao.com/coupon/edetail?e=wog6mTjZv%2FIGQASttHIRqT8%2FTgo%2BDKYNn3MtPDDRCxLuFdGM%2Fy8ADJwJ6dLD4BPh4h5828KZV1siM37WwZ3M460TeAL%2BmcF1hLjdYI0pIJBwx7%2FmoMcv1BemP0hpIIPvjDppvlX%2Bob8NlNJBuapvQ2MDg9t1zp0R8pjV3C9qcwRg5t1bGd7ixKOQJ6Ic4SbV&traceId=0b0b40de15501301853222175e&union_lens=lensId:0b1832c2_0b51_168eaf5847e_8589&xId=vJinjnbsDyrka2dVuwvpFpBvPqgdfSmfEm3bA2Jig86yBLgQear0ZghwNkuOsiTtahUYTf5Kq7I6X8iIlkrpeU&activityId=d1204e96d378436cabe64bfebccc6f74"}]`)

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkSpreadGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkSpreadGetResponse)

		fmt.Println(result.TbkSpreadGetResult.Results.TbkSpread)

	}
}

func TestTbkItemCouponGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkItemCouponGetRequest{}
	getRequest.AddParameter("pid", "mm_27437251_7418269_24546980")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkItemCouponGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkItemCouponGetResponse)

		fmt.Println(result.TbkItemCouponGetResult.Results.TbkCoupon)

	}
}

func TestTbkitemGet(t *testing.T) {

	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkItemGetRequest{}
	getRequest.AddParameter("fields", "num_iid")
	getRequest.AddParameter("q", "连衣裙")
	getRequest.AddParameter("page_size", "40")
	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkItemGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkItemGetResponse)

		fmt.Println(result.TbkItemGetResult.Results.NTbkItem)

	}
}

func TestTbkItemRecommendGet(t *testing.T) {

	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.TbkItemRecommendGetRequest{}
	getRequest.AddParameter("fields", "num_iid,title,pict_url,small_images,reserve_price,zk_final_price,user_type,provcity,item_url,seller_id,volume,nick")
	getRequest.AddParameter("num_iid", "562775487987")
	getRequest.AddParameter("count", "20")
	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkItemRecommendGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkItemRecommendGetResponse)

		fmt.Println(result.TbkItemRecommendGetResult.Results.NTbkItem)

	}

}

func TestJuItemsSearch(t *testing.T) {

	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	sessionKey := os.Getenv("SESSIONKEY")
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, sessionKey)

	//初始化请求接口信息
	getRequest := &request.JuItemsSearchRequest{}
	getRequest.AddParameter("param_top_item_query", `{"current_page":1,"page_size":2,"pid":"mm_27437251_7418269_24546980"}`)

	//初始化结果类型
	var getResponse DefaultResponse = &response.JuItemsSearchResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.JuItemsSearchResponse)

		fmt.Println(result.JuItemsSearchResult.Result)

	}

}

func TestTbkActivitylinkGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkActivitylinkGetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")

	getRequest.AddParameter("promotion_scene_id", "8493178")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkActivitylinkGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkActivitylinkGetResponse)

		fmt.Println(result.TbkActivitylinkGetResult)

	}
}

func TestTbkDgItemCouponGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkDgItemCouponGetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")

	getRequest.AddParameter("cat", "16")
	getRequest.AddParameter("page_size", "2")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkDgItemCouponGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkDgItemCouponGetResponse)

		fmt.Println(result.TbkDgItemCouponGetResult.Results.TbkCoupon)

	}
}

func TestTbkDgMaterialOptional(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkDgMaterialOptionalRequest{}
	getRequest.AddParameter("adzone_id", "24546980")

	getRequest.AddParameter("cat", "16")
	getRequest.AddParameter("page_size", "2")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkDgMaterialOptionalResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkDgMaterialOptionalResponse)

		fmt.Println(result.TbkDgMaterialOptionalResult.ResultList.MapData)

	}
}

func TestTbkDgNewuserOrderGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkDgNewuserOrderGetRequest{}
	getRequest.AddParameter("activity_id", "119013_6")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkDgNewuserOrderGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkDgNewuserOrderGetResponse)

		fmt.Println(result.TbkDgNewuserOrderGetResult.Results.Data.Results.MapData)

	}
}

func TestTbkDgNewuserOrderSum(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkDgNewuserOrderSumRequest{}
	getRequest.AddParameter("activity_id", "119013_6")
	getRequest.AddParameter("page_size", "2")
	getRequest.AddParameter("page_no", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkDgNewuserOrderSumResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkDgNewuserOrderSumResponse)

		fmt.Println(result.TbkDgNewuserOrderSumResult.Results.Data.Results)

	}
}

func TestTbkDgOptimusMaterial(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkDgOptimusMaterialRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("material_id", "4093")

	getRequest.AddParameter("page_size", "2")
	getRequest.AddParameter("page_no", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkDgOptimusMaterialResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkDgOptimusMaterialResponse)

		fmt.Println(result.TbkDgOptimusMaterialResult.ResultList.MapData)

	}
}

func TestTbkJuTqgGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkJuTqgGetRequest{}
	getRequest.AddParameter("adzone_id", "24546980")
	getRequest.AddParameter("fields", "click_url,pic_url,reserve_price,zk_final_price,total_amount,sold_num,title,category_name,start_time,end_time,num_iid")

	getRequest.AddParameter("start_time", "2019-02-15 07:00:00")
	getRequest.AddParameter("end_time", "2019-02-15 15:00:00")

	getRequest.AddParameter("page_size", "2")
	getRequest.AddParameter("page_no", "1")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkJuTqgGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkJuTqgGetResponse)

		fmt.Println(result.TbkJuTqgGetResult.Results.Results)

	}
}

func TestTbkCouponGet(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkCouponGetRequest{}
	getRequest.AddParameter("item_id", "573509331754")

	getRequest.AddParameter("activity_id", "335fa4c30e8a4f5da28596d122b48c28")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkCouponGetResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkCouponGetResponse)

		fmt.Println(result.TbkCouponGetResult.Data)

	}
}

func TestTbkTpwdCreate(t *testing.T) {
	appKey := os.Getenv("APPKEY")
	appSecret := os.Getenv("APPSECRET")
	//sessionKey := os.Getenv("SESSIONKEY")
	//fmt.Println(sessionKey)
	//
	//初始化TopClient
	client := &TopClient{}
	client.Init(appKey, appSecret, "")

	//初始化请求接口信息
	getRequest := &request.TbkTpwdCreateRequest{}
	getRequest.AddParameter("text", "573509331754")
	getRequest.AddParameter("url", "https://uland.taobao.com/coupon/edetail?e=mO4IEAPR5NAGQASttHIRqSS40sV7V7nidNcM7bsgMaPuFdGM%2Fy8ADJwJ6dLD4BPh4h5828KZV1siM37WwZ3M460TeAL%2BmcF1hLjdYI0pIJBwx7%2FmoMcv1BemP0hpIIPvjDppvlX%2Bob8NlNJBuapvQ2MDg9t1zp0R8pjV3C9qcwTuCTbQanBq6DXyPsc63NLC&traceId=0b15285b15502249874743900e&union_lens=lensId:0b0175f0_0bb6_168f09c1564_38e3&xId=9qB7wC7cWoDxvRAEWhBT5GywBdnahUokfIoYEnBxjUdmgTwHqmSiABnh5LI6gVcBmjPO7VeGfMM0Cqk2zkVD84&activityId=cd802bbce5554df0a9d0a2b93230ae65")

	//初始化结果类型
	var getResponse DefaultResponse = &response.TbkTpwdCreateResponse{}
	//执行请求接口得到结果
	err := client.Exec(getRequest, getResponse)
	if err != nil {
		t.Log(err)
	} else {
		result := getResponse.(*response.TbkTpwdCreateResponse)

		fmt.Println(result.Body)

	}
}
