package internal

import (
	"fmt"
	"net/http"
)

type Client struct {
	Strict
}

type Strict interface {
	Fatal(args ...interface{})
}

type StrictFunc func(...interface{})

func (s StrictFunc) Fatal(args ...interface{}) {
	s.Fatal(args...)
}

func NewClient() *Client {
	return &Client{
		Strict: lax,
	}
}

var lax = StrictFunc(func(...interface{}) {})

func (c *Client) Do(r *http.Request) (*http.Response, error) {
	if err := c.checkContentType(r); err != nil {
		c.Fatal(err)
		return nil, err
	}
	resp, err := http.DefaultClient.Do(r)
	c.Fatal(err)
	return resp, err
}

func (c *Client) checkContentType(r *http.Request) error {
	got := r.Header.Get("Content-Type")
	exp := "application/json"
	if got != exp {
		return fmt.Errorf("checkContentType: %q must be %s", got, exp)
	}
	return nil
}
