package main

import (
	"context"
)

type LoginFn[P AuthParameter, R AuthClaimResponse] func(context.Context, P) (R, error)

type AuthParameter interface {
	Employee | Admin | Customer
}

type AuthClaimResponse interface {
	*GeneralClaimResponse | *AdminClaimResponse
}

type Occupation int

const (
	ENGINEER Occupation = iota
	ARCHITECT
	TESTER
	MANAGER
	DEVOPS
)

type ACL int

const (
	READER ACL = iota
	WRITER
)

type CustomerCategory int

const (
	GOLD CustomerCategory = iota
	PLATINUM
	PARTNER
)

type MemberType int

const (
	Emp MemberType = iota
	Cust
	Moderator
	Annonymous
	None
)

type AuthReq struct {
	Id       string
	Password string
}

type Employee struct {
	AuthReq
	Name        string
	Occupation  Occupation
	Salary      int
	JoiningDate string
	Location    string
}

type Admin struct {
	AuthReq
	Acl    ACL
	Policy string
	Active bool
}

type Customer struct {
	AuthReq
	CustomerCategory CustomerCategory
	JoiningDate      string
	Organization     string
}

type GeneralClaimResponse struct {
	Id         string
	Token      string
	License    string
	MemberType MemberType
	*ErrorResponse
}

type AdminClaimResponse struct {
	Id         string
	Token      string
	ACL        ACL
	MemberType MemberType
	Active     bool
	*ErrorResponse
}

