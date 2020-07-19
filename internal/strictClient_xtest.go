package internal

import (
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	c := NewClient()
	c.Strict = t
	r, _ := http.NewRequest("GET", "nothing", nil)
	c.Do(r)
}
