package response

import "encoding/json"

//taobao.tbk.dg.material.optional( 通用物料搜索API（导购） )
type TbkDgMaterialOptionalResponse struct {
	TopResponse
	TbkDgMaterialOptionalResult TbkDgMaterialOptionalResult `json:"tbk_dg_material_optional_response"`
}

//解析输出结果
func (t *TbkDgMaterialOptionalResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkDgMaterialOptionalResult struct {
	ResultList   TbkDgMaterialOptionalResultList `json:"result_list"`
	TotalResults int64                           `json:"total_results"`
	RequestID    string                          `json:"request_id"`
}

type TbkDgMaterialOptionalResultList struct {
	MapData []TbkDgMaterialOptionalMapDatum `json:"map_data"`
}

type TbkDgMaterialOptionalMapDatum struct {
	CategoryID           int64                            `json:"category_id"`
	CategoryName         string                           `json:"category_name"`
	CommissionRate       string                           `json:"commission_rate"`
	CommissionType       string                           `json:"commission_type"`
	CouponID             string                           `json:"coupon_id"`
	CouponInfo           string                           `json:"coupon_info"`
	CouponShareURL       string                           `json:"coupon_share_url"`
	CouponRemainCount    int64                            `json:"coupon_remain_count"`
	CouponTotalCount     int64                            `json:"coupon_total_count"`
	IncludeDxjh          string                           `json:"include_dxjh"`
	IncludeMkt           string                           `json:"include_mkt"`
	InfoDxjh             string                           `json:"info_dxjh"`
	ItemURL              string                           `json:"item_url"`
	LevelOneCategoryID   int64                            `json:"level_one_category_id"`
	LevelOneCategoryName string                           `json:"level_one_category_name"`
	NumIid               int64                            `json:"num_iid"`
	PictURL              string                           `json:"pict_url"`
	Provcity             string                           `json:"provcity"`
	ReservePrice         string                           `json:"reserve_price"`
	SellerID             int64                            `json:"seller_id"`
	ShopDsr              int64                            `json:"shop_dsr"`
	ShopTitle            string                           `json:"shop_title"`
	ShortTitle           string                           `json:"short_title"`
	SmallImages          TbkDgMaterialOptionalSmallImages `json:"small_images"`
	Title                string                           `json:"title"`
	TkTotalCommi         string                           `json:"tk_total_commi"`
	TkTotalSales         string                           `json:"tk_total_sales"`
	URL                  string                           `json:"url"`
	UserType             int64                            `json:"user_type"`
	Volume               int64                            `json:"volume"`
	WhiteImage           string                           `json:"white_image"`
	XID                  string                           `json:"x_id"`
	ZkFinalPrice         string                           `json:"zk_final_price"`
}

type TbkDgMaterialOptionalSmallImages struct {
	String []string `json:"string"`
}
