package v1

import (
	"fmt"
	"gitlab.com/asset-exchange/exchange-api"
)

// Order ...
type Order struct {
	uri string
}

// NewOrder ...
func NewOrder(config *api.Config) *Order {
	uri := config.URI()
	return &Order{uri: uri}
}

// GetOneOrder ...
func (o *Order) GetOneOrder(orderId string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	if orderId == "" {
		return out, nil
	}

	url := fmt.Sprintf("%s/v1/order/detail/%s", o.uri, orderId)
	err := api.Get(url, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
