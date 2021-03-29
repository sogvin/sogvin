package spaceflight

import "testing"

func Test_Service_Use(t *testing.T) {
	var srv Service
	var role Pilot
	srv.Use(&role)
}
