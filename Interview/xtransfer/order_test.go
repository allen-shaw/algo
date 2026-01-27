package xtransfer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func handler(id string) int {
	if id == "1234" {
		return 1
	}
	return 2
}

func Test_UpdateOrderHander(t *testing.T) {
	testCases := []struct {
		name            string
		orderID         string
		exprectedStatus int
	}{
		{
			name:            "update to success",
			orderID:         "1234",
			exprectedStatus: 1,
		},
		{
			name:            "update to failed",
			orderID:         "1235",
			exprectedStatus: 1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out := handler(tc.orderID)
			assert.Equal(t, tc.exprectedStatus, out, "wrong expected status")
		})
	}
}
