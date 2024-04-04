package glesys

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNetworkCircuitsDetails(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkcircuit": {
		"id": "ic123456",
		"type": "INTERNET_PIPELINE",
		"billing": {"currency": "SEK",
		  "price": "120.0",
		  "discount": "0.0",
		  "total": "120.0",
		  "details": []
		}
		} } }`}
	s := NetworkCircuitService{client: c}

	ic, _ := s.Details(context.Background(), "ic123456")

	assert.Equal(t, "POST", c.lastMethod, "method used is correct")
	assert.Equal(t, "networkcircuit/details", c.lastPath, "path used is correct")
	assert.Equal(t, "INTERNET_PIPELINE", ic.Type, "networkcircuit Type is correct")
}

func TestNetworkCircuitsList(t *testing.T) {
	c := &mockClient{body: `{ "response": { "networkcircuits": [{
		"id": "ic123456",
		"type": "INTERNET_PIPELINE",
		"billing": {"currency": "SEK",
		  "price": "120.0",
		  "discount": "0.0",
		  "total": "120.0",
		  "details": []
		}
		}] } }`}
	n := NetworkCircuitService{client: c}

	ics, _ := n.List(context.Background())

	assert.Equal(t, "GET", c.lastMethod, "method used is correct")
	assert.Equal(t, "networkcircuit/list", c.lastPath, "path used is correct")
	assert.Equal(t, "INTERNET_PIPELINE", (*ics)[0].Type, "networkcircuit Type is correct")
}
