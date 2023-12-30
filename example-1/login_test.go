package main

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoLogin(t *testing.T) {
	ctx := context.Background()
	testcases := []struct {
		name          string
		id            string
		password      string
		expectedError error
	}{
		{
			name:          "Test Employee Login",
			id:            "EMP_005",
			password:      "very-s3cret-passwd",
			expectedError: nil,
		},
		{
			name:          "Test Customer Login",
			id:            "CUST_1245",
			password:      "very-s3cret-passwd",
			expectedError: nil,
		},
		{
			name:          "Test Admin Login",
			id:            "ADMIN_652",
			password:      "very-s3cret-passwd",
			expectedError: nil,
		},
		{
			name:          "Test Failed Employee Login",
			id:            "EMP_005",
			password:      "very-passwd",
			expectedError: errors.New("login error response: employee login failed with wrong credentials , Code: 401 "),
		},
		{
			name:          "Test Failed Customer Login",
			id:            "CUST_1245",
			password:      "-s3cret-passwd",
			expectedError: errors.New("login error response: customer login failed with wrong credentials , Code: 401 "),
		},
		{
			name:          "Test Failed Admin Login",
			id:            "ADMIN_652",
			password:      "very-s3cret",
			expectedError: errors.New("login error response: admin login failed with wrong credentials , Code: 401 "),
		},
		{
			name:          "Unknown role",
			id:            "Unknown_123",
			expectedError: errors.New("unsupported request type"),
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			err := doLogin(ctx, testcase.id, testcase.password)
			if err != nil {
				assert.Equal(t, err, testcase.expectedError)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
