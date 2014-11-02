package client

import (
	"fmt"

	"github.com/peterjliu/rest"
)

type Client struct {
	AppId  string // Parse-Application-Id
	ApiKey string // Parse REST API key
}

// Get new Parse REST client
func New(id, key string) *Client {
	c := new(Client)
	c.AppId = id
	c.ApiKey = key
	return c
}

func (c Client) Headers() map[string]string {
	return map[string]string{
		"X-Parse-Application-Id": c.AppId,
		"X-Parse-REST-API-Key":   c.ApiKey,
	}
}

// Get Parse object and fill in struct
func (c Client) GetObj(class, id string, out interface{}) error {
	r := rest.Request{
		Method:  rest.GET,
		Headers: c.Headers(),
		Url:     fmt.Sprintf("https://api.parse.com/1/classes/%s/%s", class, id),
	}
	err := r.Do(&out)
	if err != nil {
		fmt.Println("Failed to get object")
		return err
	}
	return nil
}

// Get Parse object, without specifying type
func (c Client) GetObjMap(class, id string) (map[string]interface{}, error) {
	var v interface{}
	r := rest.Request{
		Method:  rest.GET,
		Headers: c.Headers(),
		Url:     fmt.Sprintf("https://api.parse.com/1/classes/%s/%s", class, id),
	}
	err := r.Do(&v)
	if err != nil {
		fmt.Println("Failed to get object")
		return nil, err
	}
	return v.(map[string]interface{}), nil
}
