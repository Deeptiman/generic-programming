package main

import (
	"context"
)

// Employee, Admin, Customer
func main() {
	ctx := context.Background()

	id := "EMP_9821"
	password := "very-s3cret-passwd"

	err := doLogin(ctx, id, password)
	if err != nil {
		panic(err)
	}
}
