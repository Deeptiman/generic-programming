# Generic Programming in Go
Golang released generic programming with a specific goal for developers to write a standard set of types to utilize as type parameters for functions and type structs. The syntax to write generic type comes with a `func F[T any] (p T){ ... }` that can be used to declare common types, and the function arguments
will use the set of types as parameters. Also, types can have type parameters list `type O[T any] struct.`

## Publication


## Repository
This repository will provide an example module written with the Go generic type structure. The following code snippet shows a middleware module to showcase the **type parameter** and **function type arguments** features in the Go generic design.

````````````````go

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
````````````````

### How to run the test?

``````sh
go test -v -count=1 -race -cover --coverprofile=unit.out ./... && go tool cover -html=unit.out
``````

## License
This project is licensed under the <a href="https://github.com/Deeptiman/generic-programming/blob/main/LICENSE">MIT License</a>
