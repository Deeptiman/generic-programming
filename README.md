# Generic Programming in Go
Golang released generic programming with a specific goal for developers to write a standard set of types to utilize as type parameters for functions and type structs. The syntax to write a generic function is like `func F[T any] (p T){ … }` that can be used to perform a common use case. Also, types can have type parameters list `type O[T1, T2 any] struct`.

## Publication
**Deep Dive into Go Generic Type Structures and Syntax**

https://codingpirate.com/deep-dive-into-go-generic-type-structures-and-syntax-6f1a68e2c9c5

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

## References
- https://go.googlesource.com/proposal/+/HEAD/design/43651-type-parameters.md
- https://go.dev/blog/type-inference

## License
This project is licensed under the <a href="https://github.com/Deeptiman/generic-programming/blob/main/LICENSE">MIT License</a>
