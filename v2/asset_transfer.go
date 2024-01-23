package binance

import (
	"context"
	"net/http"
)

type AssetTransferType string

const (
	AssetTransfer_MAIN_UMFUTURE AssetTransferType = "MAIN_UMFUTURE" //  现货钱包转向U本位合约钱包
	// MAIN_CMFUTURE 现货钱包转向币本位合约钱包
	// MAIN_MARGIN 现货钱包转向杠杆全仓钱包
	AssetTransfer_UMFUTURE_MAIN AssetTransferType = "UMFUTURE_MAIN" //  U本位合约钱包转向现货钱包
	// AssetTransfer_MAIN_FUNDING UMFUTURE_MARGIN U本位合约钱包转向杠杆全仓钱包
	// CMFUTURE_MAIN 币本位合约钱包转向现货钱包
	// MARGIN_MAIN 杠杆全仓钱包转向现货钱包
	// MARGIN_UMFUTURE 杠杆全仓钱包转向U本位合约钱包
	// MARGIN_CMFUTURE 杠杆全仓钱包转向币本位合约钱包
	// CMFUTURE_MARGIN 币本位合约钱包转向杠杆全仓钱包
	// ISOLATEDMARGIN_MARGIN 杠杆逐仓钱包转向杠杆全仓钱包
	// MARGIN_ISOLATEDMARGIN 杠杆全仓钱包转向杠杆逐仓钱包
	// ISOLATEDMARGIN_ISOLATEDMARGIN 杠杆逐仓钱包转向杠杆逐仓钱包
	AssetTransfer_MAIN_FUNDING     AssetTransferType = "MAIN_FUNDING"     //  现货钱包转向资金钱包
	AssetTransfer_FUNDING_MAIN     AssetTransferType = "FUNDING_MAIN"     //  资金钱包转向现货钱包
	AssetTransfer_FUNDING_UMFUTURE AssetTransferType = "FUNDING_UMFUTURE" //  资金钱包转向U本位合约钱包
	AssetTransfer_UMFUTURE_FUNDING AssetTransferType = "UMFUTURE_FUNDING" //  U本位合约钱包转向资金钱包

	// MARGIN_FUNDING 杠杆全仓钱包转向资金钱包
	// FUNDING_MARGIN 资金钱包转向杠杆全仓钱包
	// FUNDING_CMFUTURE 资金钱包转向币本位合约钱包
	// CMFUTURE_FUNDING 币本位合约钱包转向资金钱包
	// MAIN_OPTION 现货钱包转向期权钱包
	// OPTION_MAIN 期权钱包转向现货钱包
	// UMFUTURE_OPTION U本位合约钱包转向期权钱包
	// OPTION_UMFUTURE 期权钱包转向U本位合约钱包
	// MARGIN_OPTION 杠杆全仓钱包转向期权钱包
	// OPTION_MARGIN 期权全仓钱包转向杠杆钱包
	// FUNDING_OPTION 资金钱包转向期权钱包
	// OPTION_FUNDING 期权钱包转向资金钱包

	AssetTransfer_MAIN_PORTFOLIO_MARGIN AssetTransferType = "MAIN_PORTFOLIO_MARGIN" //  现货钱包转向统一账户钱包
	AssetTransfer_PORTFOLIO_MARGIN_MAIN AssetTransferType = "PORTFOLIO_MARGIN_MAIN" //  统一账户钱包转向现货钱包
	AssetTransfer_MAIN_ISOLATED_MARGIN  AssetTransferType = "MAIN_ISOLATED_MARGIN"  //  现货钱包转向逐仓账户钱包
	AssetTransfer_ISOLATED_MARGIN_MAIN  AssetTransferType = "ISOLATED_MARGIN_MAIN"  //  逐仓钱包转向现货账户钱包
)

func (c *Client) NewAssetTransfer() *AssetTransfer {
	return &AssetTransfer{
		c: c,
	}
}

type AssetTransfer struct {
	c *Client

	typ        *AssetTransferType
	asset      *string
	amount     *float64
	fromSymbol *string
	toSymbol   *string
	recvWindow *int64
	timestamp  *int64
}

func (s *AssetTransfer) Type(typ AssetTransferType) *AssetTransfer {
	s.typ = &typ
	return s
}
func (s *AssetTransfer) Asset(asset string) *AssetTransfer {
	s.asset = &asset
	return s
}
func (s *AssetTransfer) Amount(amount float64) *AssetTransfer {
	s.amount = &amount
	return s
}
func (s *AssetTransfer) FromSymbol(fromSymbol string) *AssetTransfer {
	s.fromSymbol = &fromSymbol
	return s
}
func (s *AssetTransfer) ToSymbol(toSymbol string) *AssetTransfer {
	s.toSymbol = &toSymbol
	return s
}
func (s *AssetTransfer) RecvWindow(recvWindow int64) *AssetTransfer {
	s.recvWindow = &recvWindow
	return s
}
func (s *AssetTransfer) Timestamp(timestamp int64) *AssetTransfer {
	s.timestamp = &timestamp
	return s
}

func (s *AssetTransfer) Do(ctx context.Context) (*AssertTransferResponse, error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/sapi/v1/asset/transfer",
		secType:  secTypeSigned,
	}
	if s.typ != nil {
		r.setParam("type", *s.typ)
	}
	if s.asset != nil {
		r.setParam("asset", *s.asset)
	}
	if s.amount != nil {
		r.setParam("amount", *s.amount)
	}
	if s.fromSymbol != nil {
		r.setParam("fromSymbol", *s.fromSymbol)
	}
	if s.toSymbol != nil {
		r.setParam("toSymbol", *s.toSymbol)
	}
	if s.recvWindow != nil {
		r.setParam("recvWindow", *s.recvWindow)
	}
	if s.timestamp != nil {
		r.setParam("timestamp", *s.timestamp)
	}
	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}
	res := new(AssertTransferResponse)
	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}

type AssertTransferResponse struct {
	TranId int64 `json:"tranId"`
}
