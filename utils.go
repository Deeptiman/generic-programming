package main

import (
	"crypto/rand"
)

const random = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomTokenGenerator() (string, error) {
	bytes := make([]byte, 32)

	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	for i, b := range bytes {
		bytes[i] = random[b%byte(len(random))]
	}

	return string(bytes), nil
}

// Dummy Database
var IdSearchDb = map[string]MemberType{
	"CUST_001":  Cust,
	"EMP_005":   Emp,
	"EMP_585":   Emp,
	"ADMIN_100": Moderator,
	"CUST_1245": Cust,
	"EMP_678":   Emp,
	"ADMIN_652": Moderator,
	"EMP_9821":  Emp,
	"EMP_470":   Emp,
	"CUST_537":  Cust,
}

var authSecureDb = map[string]AuthReq{
	"EMP_678":   {Id: "EMP_678", Password: "very-s3cret-passwd"},
	"EMP_005":   {Id: "EMP_005", Password: "very-s3cret-passwd"},
	"EMP_9821":  {Id: "EMP_9821", Password: "very-s3cret-passwd"},
	"EMP_585":   {Id: "EMP_585", Password: "very-s3cret-passwd"},
	"EMP_470":   {Id: "EMP_470", Password: "very-s3cret-passwd"},
	"CUST_001":  {Id: "CUST_001", Password: "very-s3cret-passwd"},
	"CUST_1245": {Id: "CUST_1245", Password: "very-s3cret-passwd"},
	"CUST_537":  {Id: "CUST_537", Password: "very-s3cret-passwd"},
	"ADMIN_100": {Id: "ADMIN_100", Password: "very-s3cret-passwd"},
	"ADMIN_652": {Id: "ADMIN_652", Password: "very-s3cret-passwd"},
}

var employeeDb = map[string]Employee{
	"EMP_678":  {Name: "John Doe", Occupation: ENGINEER, Salary: 180000, JoiningDate: "2023-10-15T12:34:37Z", Location: "San Francisco"},
	"EMP_005":  {Name: "Michael Scott", Occupation: ARCHITECT, Salary: 220000, JoiningDate: "2022-05-19T10:23:17Z", Location: "New York"},
	"EMP_9821": {Name: "Jerry Seinfeld", Occupation: MANAGER, Salary: 250000, JoiningDate: "2022-07-04T04:18:45Z", Location: "San Diego"},
	"EMP_585":  {Name: "Nelson Ford", Occupation: TESTER, Salary: 130000, JoiningDate: "2023-11-18T09:27:21Z", Location: "Florida"},
	"EMP_470":  {Name: "Jaylen Frank", Occupation: DEVOPS, Salary: 210000, JoiningDate: "2023-01-27T08:17:11Z", Location: "Los Angeles"},
}

var customerDb = map[string]Customer{
	"CUST_001":  {CustomerCategory: PARTNER, JoiningDate: "2022-05-19T10:23:17Z", Organization: "Google"},
	"CUST_1245": {CustomerCategory: PLATINUM, JoiningDate: "2021-04-27T17:28:27Z", Organization: "Amazon"},
	"CUST_537":  {CustomerCategory: GOLD, JoiningDate: "2022-11-18T12:08:37Z", Organization: "Uber"},
}

var adminDb = map[string]Admin{
	"ADMIN_100": {Acl: READER, Policy: "OR('Org1MSP.member', 'Org1MSP.client')", Active: true},
	"ADMIN_652": {Acl: WRITER, Policy: "OR('Org1MSP.admin')", Active: true},
}
