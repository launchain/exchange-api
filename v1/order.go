package v1

import (
	"time"

	"github.com/launchain/exchange-api"
)

type Order struct {
	uri string
}

type ShopCartResponse struct {
	ID        string    `json:"_id"`
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
