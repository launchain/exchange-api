package v1

import (
	"time"
	"github.com/launchain/exchange-api"
)

// Email ...
type Order struct {
	uri string
}

// Code ...
type ShopCart struct {
	ID        string    `json:"id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Type      int       `json:"type"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

type ShopCartResponse struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Apikey    string    `bson:"apikey" json:"apikey"`
	AssetId   string    `bson:"assetid" json:"assetid"`
	UserId    string    `bson:"userid" json:"userid"`
}
type ShopCartsResponse struct {
	Count int                `json:"count"`
	Data  []ShopCartResponse `json:"data"`
}

// NewEmail ..
func NewOrder(c *api.Config) *Order {
	c.Check()
	uri := c.URI()
	return &Order{uri: uri}
}

// FindCode ...
func (s *Order) ShopCartList(userid string) (*[]ShopCartResponse, error) {

	out := &ShopCartsResponse{}
	err := api.Get(s.uri+"/v1/shopcart/list/"+userid, out)
	if err != nil {
		return nil, err
	}

	return &(out.Data), nil
}
