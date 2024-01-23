package binance

import (
	"context"
	"fmt"
	"net/http"
)

// ListDepositsService fetches deposit history.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-history-user_data
type ListSubDepositsService struct {
	c         *Client
	coin      *string
	status    *int
	startTime *int64
	endTime   *int64
	offset    *int
	limit     *int
	txId      *string
	email     *string
}

// Coin sets the coin parameter.
func (s *ListSubDepositsService) Coin(coin string) *ListSubDepositsService {
	s.coin = &coin
	return s
}

// Status sets the status parameter.
func (s *ListSubDepositsService) Status(status int) *ListSubDepositsService {
	s.status = &status
	return s
}

// StartTime sets the startTime parameter.
// If present, EndTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListSubDepositsService) StartTime(startTime int64) *ListSubDepositsService {
	s.startTime = &startTime
	return s
}

// EndTime sets the endTime parameter.
// If present, StartTime MUST be specified. The difference between EndTime - StartTime MUST be between 0-90 days.
func (s *ListSubDepositsService) EndTime(endTime int64) *ListSubDepositsService {
	s.endTime = &endTime
	return s
}

// Offset set offset
func (s *ListSubDepositsService) Offset(offset int) *ListSubDepositsService {
	s.offset = &offset
	return s
}

// Limit set limit
func (s *ListSubDepositsService) Limit(limit int) *ListSubDepositsService {
	s.limit = &limit
	return s
}

func (s *ListSubDepositsService) TxID(id string) *ListSubDepositsService {
	s.txId = &id
	return s
}

func (s *ListSubDepositsService) Email(email string) *ListSubDepositsService {
	s.email = &email
	return s
}

// Do sends the request.
func (s *ListSubDepositsService) Do(ctx context.Context) (res []*SubDeposit, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/subHisrec",
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.status != nil {
		r.setParam("status", *s.status)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.offset != nil {
		r.setParam("offset", *s.offset)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	if s.txId != nil {
		r.setParam("txId", *s.txId)
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}

	data, err := s.c.callAPI(ctx, r)
	fmt.Println("----------------")
	fmt.Println(data)
	fmt.Println("----------------")
	if err != nil {
		return
	}
	res = make([]*SubDeposit, 0)

	fmt.Println("----------------")
	fmt.Println(res)
	fmt.Println("----------------")

	err = json.Unmarshal(data, &res)
	if err != nil {
		return
	}
	return res, nil
}

// Deposit represents a single deposit entry.
type SubDeposit struct {
	Amount        string `json:"amount"`
	Coin          string `json:"coin"`
	Network       string `json:"network"`
	Status        int    `json:"status"`
	Address       string `json:"address"`
	AddressTag    string `json:"addressTag"`
	TxID          string `json:"txId"`
	InsertTime    int64  `json:"insertTime"`
	TransferType  int64  `json:"transferType"`
	UnlockConfirm int64  `json:"unlockConfirm"`
	ConfirmTimes  string `json:"confirmTimes"`
}

// GetDepositsAddressService retrieves the details of a deposit address.
//
// See https://binance-docs.github.io/apidocs/spot/en/#deposit-address-supporting-network-user_data
type GetSubDepositsAddressService struct {
	c       *Client
	coin    string
	email   *string
	network *string
}

// Coin sets the coin parameter (MANDATORY).
func (s *GetSubDepositsAddressService) Coin(coin string) *GetSubDepositsAddressService {
	s.coin = coin
	return s
}

// Network sets the network parameter.
func (s *GetSubDepositsAddressService) Network(network string) *GetSubDepositsAddressService {
	s.network = &network
	return s
}
func (s *GetSubDepositsAddressService) Email(email string) *GetSubDepositsAddressService {
	s.email = &email
	return s
}

// Do sends the request.
func (s *GetSubDepositsAddressService) Do(ctx context.Context) (*GetSubDepositAddressResponse, error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/capital/deposit/subAddress",
		secType:  secTypeSigned,
	}
	r.setParam("coin", s.coin)
	if s.network != nil {
		r.setParam("network", *s.network)
	}
	if s.email != nil {
		r.setParam("email", *s.email)
	}

	data, err := s.c.callAPI(ctx, r)
	if err != nil {
		return nil, err
	}

	res := &GetSubDepositAddressResponse{}

	fmt.Println("=================")
	fmt.Println(res)
	fmt.Println("=================")

	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetDepositAddressResponse represents a response from GetDepositsAddressService.
type GetSubDepositAddressResponse struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
	Coin    string `json:"coin"`
	URL     string `json:"url"`
}

func (c *Client) NewListSubDepositsService() *ListSubDepositsService {
	return &ListSubDepositsService{c: c}
}

func (c *Client) NewSubDepositsAddressService() *GetSubDepositsAddressService {
	return &GetSubDepositsAddressService{c: c}
}
