package v1

import (
	"fmt"
	"net/url"
	"time"
	"errors"

	"github.com/launchain/exchange-api"
)

// Session ...
type Session struct {
	uri string
}

// SessionResponse ...
type SessionResponse struct {
	ID        string    `json:"_id"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    string    `json:"user_id"`
	DeviceID  string    `json:"device_id"`
	Platform  int       `json:"platform"`
	Token     string    `json:"token"`
}

// NewSession ...
func NewSession(c *api.Config) *Session {
	c.Check()
	uri := c.URI()
	return &Session{uri: uri}
}

// SignIn ...
func (s *Session) SignIn(phone, password, deviceID string, platform int, captcha_id, captcha_number string) (*SessionResponse, error) {
	data := make(url.Values)
	data["phone"] = []string{phone}
	data["password"] = []string{password}
	data["device_id"] = []string{deviceID}
	data["platform"] = []string{fmt.Sprintf("%d", platform)}
	data["captcha_id"] = []string{captcha_id}
	data["captcha_number"] = []string{captcha_number}

	url := s.uri + "/v1/sessions"
	var out = make(map[string]interface{},2)
	err := api.PostForm(url, data, out)
	
	if err != nil {
		return nil, err
	}
	err = errors.New("sign in error")
	if item,ok := out["session"];ok{
		if session,ok := item.(SessionResponse);ok{
			return &session, nil
		}
	}

	return nil, err
}
