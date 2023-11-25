package main

import (
	"context"
	"fmt"
	"net/http"
)

func doLogin(ctx context.Context, id, password string) error {
	switch IdSearchDb[id] {
	case Emp:
		req := Employee{AuthReq: AuthReq{Id: id, Password: password}}
		loginFn := LoginFn[Employee, *GeneralClaimResponse](employeeLoginFn)
		resp, err := authMiddleware(ctx, loginFn, req)
		if e := As(err); e != nil {
			return fmt.Errorf("login error response: %s , Code: %d ", e.Message(), e.Code())
		}

		// TODO: do something with authenticated Employee response
		fmt.Println("Employee Authenticated: ID=", resp.Id, ", License=", resp.License, ", Token=", resp.Token)

	case Cust:
		req := Customer{AuthReq: AuthReq{Id: id, Password: password}}
		loginFn := LoginFn[Customer, *GeneralClaimResponse](customerLoginFn)
		resp, err := authMiddleware(ctx, loginFn, req)
		if e := As(err); e != nil {
			return fmt.Errorf("login error response: %s , Code: %d ", e.Message(), e.Code())
		}

		// TODO: do something with authenticated Customer response
		fmt.Println("Customer Authenticated: ID=", resp.Id, ", License=", resp.License, ", Token=", resp.Token)

	case Moderator:
		req := Admin{AuthReq: AuthReq{Id: id, Password: password}}
		loginFn := LoginFn[Admin, *AdminClaimResponse](adminLoginFn)
		resp, err := authMiddleware(ctx, loginFn, req)
		if e := As(err); e != nil {
			return fmt.Errorf("login error response: %s , Code: %d ", e.Message(), e.Code())
		}

		// TODO: do something with authenticated Admin response
		fmt.Println("Admin Authenticated: ID=", resp.Id, ", ACL=", resp.ACL, ", Active=", resp.Active, ", Token=", resp.Token)
	case None:
		panic("unsupported request type")
	}

	return nil
}

var employeeLoginFn = func(ctx context.Context, emp Employee) (*GeneralClaimResponse, error) {
	// try to authenticate with dummy auth secure db
	if authSecureDb[emp.Id].Password != emp.Password {
		return nil, ErrorResponse{code: http.StatusUnauthorized, message: "employee login failed with wrong credentials"}
	}

	token, err := randomTokenGenerator()
	if err != nil {
		return nil, ErrorResponse{code: http.StatusUnauthorized, message: "token generation failed"}
	}

	e := employeeDb[emp.Id]
	fmt.Println("Employee Details: Id=", emp.Id, " Name=", e.Name, " Occupation=", e.Occupation, " Salary=", e.Salary, " Location=", e.Location)

	return &GeneralClaimResponse{
		Id:         emp.Id,
		Token:      fmt.Sprintf("token-%s-%s", emp.Id, token),
		License:    fmt.Sprintf("license-%s", emp.Id),
		MemberType: Emp,
	}, nil
}

var customerLoginFn = func(ctx context.Context, customer Customer) (*GeneralClaimResponse, error) {
	// try to authenticate with dummy auth secure db
	if authSecureDb[customer.Id].Password != customer.Password {
		return nil, ErrorResponse{code: http.StatusUnauthorized, message: "customer login failed with wrong credentials"}
	}

	token, err := randomTokenGenerator()
	if err != nil {
		return nil, ErrorResponse{code: http.StatusUnauthorized, message: "token generation failed"}
	}

	c := customerDb[customer.Id]
	fmt.Println("Customer Details: Id=", customer.Id, " Category=", c.CustomerCategory, " Organization=", c.Organization)

	return &GeneralClaimResponse{
		Id:         customer.Id,
		Token:      fmt.Sprintf("token-%s-%s", customer.Id, token),
		License:    fmt.Sprintf("license-%s", customer.Id),
		MemberType: Cust,
	}, nil
}

var adminLoginFn = func(ctx context.Context, admin Admin) (*AdminClaimResponse, error) {
	// try to authenticate with dummy auth secure db
	if authSecureDb[admin.Id].Password != admin.Password {
		return nil, ErrorResponse{code: http.StatusUnauthorized, message: "admin login failed with wrong credentials"}
	}

	token, err := randomTokenGenerator()
	if err != nil {
		return nil, ErrorResponse{code: http.StatusUnauthorized, message: "token generation failed"}
	}

	a := adminDb[admin.Id]
	fmt.Println("Admin Details: Id=", a.Id, " Acl=", a.Acl, " Active=", a.Active)

	return &AdminClaimResponse{
		Id:         admin.Id,
		Token:      fmt.Sprintf("token-%s-%s", admin.Id, token),
		ACL:        a.Acl,
		MemberType: Moderator,
		Active:     a.Active,
	}, nil
}
