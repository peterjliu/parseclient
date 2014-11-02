package client

import (
	"fmt"
	"log"

	"github.com/peterjliu/rest"
)

type Client struct {
	AppId  string // Parse-Application-Id
	ApiKey string // Parse REST API key
}

// Response when creating Parse objects
type CreateResp struct {
	CreatedAt string
	ObjectId  string
}

// Response when updating Parse objects
type UpdateResp struct {
	UpdatedAt string
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
		log.Println("Failed to get object")
		return err
	}
	return nil
}

// Get Parse objects of a certain class
func (c Client) GetList(class, out interface{}) error {
	r := rest.Request{
		Method:  rest.GET,
		Headers: c.Headers(),
		Url:     fmt.Sprintf("https://api.parse.com/1/classes/%s", class),
	}
	err := r.Do(&out)
	if err != nil {
		log.Println("Failed to get objects")
		return err
	}
	return nil
}

// Get Parse object and fill in struct
func (c Client) DeleteObj(class, id string, out interface{}) error {
	r := rest.Request{
		Method:  rest.DELETE,
		Headers: c.Headers(),
		Url:     fmt.Sprintf("https://api.parse.com/1/classes/%s/%s", class, id),
	}
	err := r.Do(&out)
	if err != nil {
		log.Println("Failed to get object")
		return err
	}
	return nil
}

// Creates a new Parse object of given class.
func (c Client) CreateObj(class, in interface{}) (*CreateResp, error) {
	r := new(CreateResp)
	err := rest.Post(
		fmt.Sprintf("https://api.parse.com/1/classes/%s", class),
		c.Headers(),
		in,
		r)
	if err != nil {
		log.Println("Failed to create object")
		return nil, err
	}
	return r, nil
}

func (c Client) UpdateObj(class, id string, in interface{}) (*UpdateResp, error) {
	var r UpdateResp
	err := rest.Put(
		fmt.Sprintf("https://api.parse.com/1/classes/%s/%s", class, id),
		c.Headers(),
		in,
		&r)
	if err != nil {
		log.Println("Failed to create object")
		return nil, err
	}
	return &r, nil
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
		log.Println("Failed to get object")
		return nil, err
	}
	return v.(map[string]interface{}), nil
}
