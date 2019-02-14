package response

import "encoding/json"

//taobao.tbk.dg.optimus.material( 淘宝客物料下行-导购 )
type TbkDgOptimusMaterialResponse struct {
	TopResponse
	TbkDgOptimusMaterialResult TbkDgOptimusMaterialResult `json:"tbk_dg_optimus_material_response"`
}

//解析输出结果
func (t *TbkDgOptimusMaterialResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkDgOptimusMaterialResult struct {
	ResultList TbkDgOptimusMaterialResultList `json:"result_list"`
	RequestID  string                         `json:"request_id"`
}

type TbkDgOptimusMaterialResultList struct {
	MapData []TbkDgOptimusMaterialMapDatum `json:"map_data"`
}

type TbkDgOptimusMaterialMapDatum struct {
	CategoryID           int64                           `json:"category_id"`
	ClickURL             string                          `json:"click_url"`
	CommissionRate       string                          `json:"commission_rate"`
	CouponAmount         int64                           `json:"coupon_amount"`
	CouponClickURL       string                          `json:"coupon_click_url"`
	CouponEndTime        string                          `json:"coupon_end_time"`
	CouponRemainCount    int64                           `json:"coupon_remain_count"`
	CouponStartFee       string                          `json:"coupon_start_fee"`
	CouponStartTime      string                          `json:"coupon_start_time"`
	CouponTotalCount     int64                           `json:"coupon_total_count"`
	ItemDescription      string                          `json:"item_description"`
	ItemID               int64                           `json:"item_id"`
	LevelOneCategoryID   int64                           `json:"level_one_category_id"`
	LevelOneCategoryName string                          `json:"level_one_category_name"`
	PictURL              string                          `json:"pict_url"`
	SellerID             int64                           `json:"seller_id"`
	ShopTitle            string                          `json:"shop_title"`
	SmallImages          TbkDgOptimusMaterialSmallImages `json:"small_images"`
	Title                string                          `json:"title"`
	UserType             int64                           `json:"user_type"`
	Volume               int64                           `json:"volume"`
	WhiteImage           string                          `json:"white_image"`
	ZkFinalPrice         string                          `json:"zk_final_price"`
}

type TbkDgOptimusMaterialSmallImages struct {
	String []string `json:"string"`
}
