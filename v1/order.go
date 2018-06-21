package v1

import (
<<<<<<< HEAD
	"errors"
	"fmt"
=======
>>>>>>> 23e0c7f70f3818c3dfbde1b504303eaa2bd8d37b
	"time"

	"github.com/launchain/exchange-api"
)

<<<<<<< HEAD
// Email ...
=======
>>>>>>> 23e0c7f70f3818c3dfbde1b504303eaa2bd8d37b
type Order struct {
	uri string
}

<<<<<<< HEAD
// Code ...
type ShopCart struct {
	ID        string    `json:"_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
	Type      int       `json:"type"`
	Code      string    `json:"code"`
	ExpiredAt time.Time `json:"expired_at"`
}

// NewEmail ..
func NewEmail(c *api.Config) *Email {
	c.Check()
	uri := c.URI()
	return &Email{uri: uri}
}

// FindCode ...
func (e *Email) FindCode(email, code string) (*Code, error) {
	if email == "" || code == "" {
		return nil, errors.New("参数错误")
	}

	url := fmt.Sprintf("%s/v1/emailcode/%s/email/%s", e.uri, code, email)
	c := &Code{}
	err := api.Get(url, c)
	if err != nil {
		return nil, err
	}
	return c, nil
=======
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
>>>>>>> 23e0c7f70f3818c3dfbde1b504303eaa2bd8d37b
}
