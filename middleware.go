package main

import (
	"context"
	"fmt"
)

type LoginFn[P AuthParameter, R AuthClaimResponse] func(context.Context, P) (R, error)

type AuthParameter interface {
	Employee | Admin | Customer
}

type AuthClaimResponse interface {
	*GeneralClaimResponse | *AdminClaimResponse
}

func authMiddleware[P AuthParameter, R AuthClaimResponse](
	ctx context.Context,
	loginFn LoginFn[P, R], authReq P) (R, error) {
	var retry = 5
	return loginWithRetry(ctx, retry, loginFn, authReq)
}

func loginWithRetry[P AuthParameter, R AuthClaimResponse](
	ctx context.Context,
	retry int,
	loginFn LoginFn[P, R], authReq P) (R, error) {
	resp, err := loginFn(ctx, authReq)
	if err != nil {
		if retry--; retry >= 0 {
			return loginWithRetry(ctx, retry, loginFn, authReq)
		}

		return nil, err
	}

	fmt.Println("login successful ")

	return resp, nil
}
