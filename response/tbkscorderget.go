package response

import "encoding/json"

//taobao.tbk.sc.order.get( 淘宝客订单查询 - 社交 )
type TbkScOrderGetResponse struct {
	TopResponse
	TbkScOrderGetResult TbkScOrderGetResult `json:"tbk_sc_order_get_response"`
}

//解析输出结果
func (t *TbkScOrderGetResponse) WrapResult(result string) {
	unmarshal := json.Unmarshal([]byte(result), t)
	//保存原始信息
	t.Body = result
	//解析错误
	if unmarshal != nil {
		t.ErrorResponse.Code = -1
		t.ErrorResponse.Msg = unmarshal.Error()
	}
}

type TbkScOrderGetResult struct {
	Results   TbkScOrderGetResults `json:"results"`
	RequestID string  `json:"request_id"`
}

type TbkScOrderGetResults struct {
	NTbkOrder []TbkScOrderGetNTbkOrder `json:"n_tbk_order"`
}

type TbkScOrderGetNTbkOrder struct {
	AdzoneID            string `json:"adzone_id"`
	AdzoneName          string `json:"adzone_name"`
	AlipayTotalPrice    string `json:"alipay_total_price"`
	AuctionCategory     string `json:"auction_category"`
	ClickTime           string `json:"click_time"`
	Commission          string `json:"commission"`
	CommissionRate      string `json:"commission_rate"`
	CreateTime          string `json:"create_time"`
	IncomeRate          string `json:"income_rate"`
	ItemNum             int64  `json:"item_num"`
	ItemTitle           string `json:"item_title"`
	NumIid              int64  `json:"num_iid"`
	OrderType           string `json:"order_type"`
	PayPrice            string `json:"pay_price"`
	Price               string `json:"price"`
	PubSharePreFee      string `json:"pub_share_pre_fee"`
	SellerNick          string `json:"seller_nick"`
	SellerShopTitle     string `json:"seller_shop_title"`
	SiteID              string `json:"site_id"`
	SiteName            string `json:"site_name"`
	SubsidyRate         string `json:"subsidy_rate"`
	SubsidyType         string `json:"subsidy_type"`
	TerminalType        string `json:"terminal_type"`
	Tk3RDType           string `json:"tk3rd_type"`
	TkStatus            int64  `json:"tk_status"`
	TotalCommissionRate string `json:"total_commission_rate"`
	TradeID             int64  `json:"trade_id"`
	TradeParentID       int64  `json:"trade_parent_id"`
}

